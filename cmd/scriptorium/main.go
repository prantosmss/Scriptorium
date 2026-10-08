package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/prantosmss/Scriptorium/internal/bootstrap"
	"github.com/prantosmss/Scriptorium/internal/eval"
	"github.com/prantosmss/Scriptorium/internal/i18n"
	"github.com/prantosmss/Scriptorium/internal/rules"
	buildversion "github.com/prantosmss/Scriptorium/internal/version"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

// headlessMode 记录本次是否 headless 启动，供 die 决定错误退出时是否暂停。
var headlessMode bool

func main() {
	// 界面语言必须先于一切子命令分发生效：doctor/service/skills/rag 这些在
	// LoadConfig 之前就被拦截，等配置加载完再切语言就来不及了。
	i18n.Set(bootstrap.LoadLanguage(scanConfigFlag(os.Args[1:])))

	// 子命令在常规 flag 解析之前拦截：eval 是离线评测 harness，参数体系独立。
	if len(os.Args) > 1 && os.Args[1] == "eval" {
		os.Exit(eval.Command(os.Args[2:]))
	}
	if len(os.Args) > 1 && os.Args[1] == "lang" {
		os.Exit(runLangCommand(os.Args[2:], scanConfigFlag(os.Args[1:])))
	}
	if len(os.Args) > 1 && os.Args[1] == "service" {
		os.Exit(runServiceCommand(os.Args[2:]))
	}
	if len(os.Args) > 1 && os.Args[1] == "doctor" {
		os.Exit(runDoctorCommand(os.Args[2:], versionInfo()))
	}
	if len(os.Args) > 1 && os.Args[1] == "skills" {
		os.Exit(runSkillsCommand(os.Args[2:]))
	}
	if len(os.Args) > 1 && os.Args[1] == "rag" {
		os.Exit(runRAGCommand(os.Args[2:]))
	}
	if len(os.Args) > 1 && (os.Args[1] == "list" || os.Args[1] == "novels") {
		os.Exit(runNovelsListCommand(os.Args[2:]))
	}
	if len(os.Args) > 1 && os.Args[1] == "reader-metrics" {
		opts, _, err := parseCLIOptions(nil)
		if err != nil {
			die("flags: %v", err)
		}
		if err := readerMetricsPipeline(opts, os.Args[2:]); err != nil {
			die("reader-metrics: %v", err)
		}
		return
	}

	opts, args, err := parseCLIOptions(os.Args[1:])
	if err != nil {
		die("flags: %v", err)
	}
	// 顶层 --help：无子命令时打印顶层 usage；有子命令 token 时把 --help 留给子命令自己处理。
	if opts.Help && !hasAnySubcommand(args) {
		printTopUsage(os.Stdout)
		return
	}
	if opts.Version {
		buildversion.Print(os.Stdout, versionInfo())
		return
	}
	if opts.Update {
		if err := runSelfUpdate(opts.UpdateVersion); err != nil {
			fmt.Fprintf(os.Stderr, "update: %v\n", err)
			os.Exit(1)
		}
		return
	}
	headlessMode = opts.Headless

	// 注意：纯路由 token（不期望值的）需要从 argv 中剥离，否则 Go flag 包会把
	// 后面的 --from/--to/--budget 误当作该 string flag 的值（如 --review-existing
	// 是 fs.StringVar 注册的，下一个 token --from 会被吞成它的字符串值）。
	if hasPipelineFlag(args) {
		headlessMode = true
		if err := pipelinePipeline(opts, stripRoutingTokens(args, "--pipeline")); err != nil {
			die("pipeline: %v", err)
		}
		return
	}
	if hasArchitectCheckFlag(args) {
		headlessMode = true
		if err := architectCheckPipeline(opts, stripRoutingTokens(args, "--architect-check")); err != nil {
			die("architect-check: %v", err)
		}
		return
	}
	if hasDraftAIJudgeFlag(args) {
		headlessMode = true
		if err := draftAIJudgePipeline(opts, stripRoutingTokens(args, "--draft-ai-judge")); err != nil {
			die("draft-ai-judge: %v", err)
		}
		return
	}
	if hasReviewExistingFlag(args) {
		headlessMode = true
		reviewArgs := stripRoutingTokens(args, "--review-existing")
		if hasHelpToken(reviewArgs) {
			if err := reviewExistingPipeline(opts, reviewArgs); err != nil {
				die("review-existing: %v", err)
			}
			return
		}
		if err := runPipelineAlias(opts, []string{"review"}, "", map[string][]string{"review": reviewArgs}); err != nil {
			die("review-existing: %v", err)
		}
		return
	}
	if hasRewriteExistingFlag(args) {
		headlessMode = true
		rewriteArgs := stripRoutingTokens(args, "--rewrite-existing")
		if hasHelpToken(rewriteArgs) {
			if err := rewriteExistingPipeline(opts, rewriteArgs); err != nil {
				die("rewrite-existing: %v", err)
			}
			return
		}
		if err := runPipelineAlias(opts, []string{"rewrite"}, "", map[string][]string{"rewrite": rewriteArgs}); err != nil {
			die("rewrite-existing: %v", err)
		}
		return
	}
	if hasZeroInitFlag(args) {
		headlessMode = true
		if err := zeroInitPipeline(opts, stripRoutingTokens(args, "--zero-init")); err != nil {
			die("zero-init: %v", err)
		}
		return
	}
	if hasCheckFlag(args) {
		headlessMode = true
		if err := checkPipeline(opts, stripRoutingTokens(args, "--check")); err != nil {
			die("check: %v", err)
		}
		return
	}
	if hasDiagFlag(args) {
		headlessMode = true
		if err := diagPipeline(opts, stripRoutingTokens(args, "--diag")); err != nil {
			die("diag: %v", err)
		}
		return
	}
	if hasWritingAssetsFlag(args) {
		headlessMode = true
		if err := writingAssetsPipeline(opts, stripRoutingTokens(args, "--writing-assets")); err != nil {
			die("writing-assets: %v", err)
		}
		return
	}
	if hasRefreshProgressFlag(args) {
		headlessMode = true
		if err := refreshProgressPipeline(opts, stripRoutingTokens(args, "--refresh-progress")); err != nil {
			die("refresh-progress: %v", err)
		}
		return
	}
	if hasBuildRAGFlag(args) {
		headlessMode = true
		if err := buildRAGPipeline(opts, stripRoutingTokens(args, "--build-rag")); err != nil {
			die("build-rag: %v", err)
		}
		return
	}
	if hasRAGReadyFlag(args) {
		headlessMode = true
		if err := ragReadyPipeline(opts, stripRoutingTokens(args, "--rag-ready")); err != nil {
			die("rag-ready: %v", err)
		}
		return
	}
	if hasSimulateFlag(args) {
		headlessMode = true
		if err := simulatePipeline(opts, stripRoutingTokens(args, "--simulate")); err != nil {
			die("simulate: %v", err)
		}
		return
	}
	if hasImportSimFlag(args) {
		headlessMode = true
		if err := importSimPipeline(opts, stripRoutingTokens(args, "--import-sim")); err != nil {
			die("import-sim: %v", err)
		}
		return
	}
	if hasSteerFlag(args) {
		headlessMode = true
		if err := steerPipeline(opts, stripRoutingTokens(args, "--steer")); err != nil {
			die("steer: %v", err)
		}
		return
	}
	if hasCocreateFlag(args) {
		headlessMode = true
		if err := cocreatePipeline(opts, stripRoutingTokens(args, "--cocreate")); err != nil {
			die("cocreate: %v", err)
		}
		return
	}

	// 首次引导
	if bootstrap.NeedsSetup(opts.ConfigPath) {
		if opts.Headless {
			die("%s", i18n.T("die.headlessNoSetup"))
		}
		setupCfg, err := bootstrap.RunSetupAt(opts.ConfigPath)
		if err != nil {
			die("setup: %v", err)
		}
		// 引导完成后使用生成的配置继续
		i18n.Set(setupCfg.Language)
		runWithConfig(setupCfg, opts, args)
		return
	}

	// 加载配置
	cfg, err := bootstrap.LoadConfig(opts.ConfigPath)
	if err != nil {
		die("config: %v", err)
	}
	i18n.Set(cfg.Language)

	runWithConfig(cfg, opts, args)
}

// die 统一处理致命错误退出：打印到 stderr、落盘到 ~/.scriptorium/last-error.log，
// 并在交互式终端（非 headless）下暂停等待回车——双击启动时控制台会随进程退出
// 立即关闭，不暂停的话错误一闪而过，正是 issue #37 里用户无从排查的根因。
func die(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintln(os.Stderr, msg)
	if path := bootstrap.WriteStartupError(msg); path != "" {
		fmt.Fprintln(os.Stderr, i18n.T("die.loggedAt", path))
	}
	if !headlessMode && stdinIsTerminal() {
		fmt.Fprint(os.Stderr, i18n.T("die.pressEnter"))
		_, _ = fmt.Fscanln(os.Stdin)
	}
	os.Exit(1)
}

// stdinIsTerminal 判断标准输入是否连接到终端（字符设备）。双击启动 / 交互式终端
// 为 true；管道、重定向、CI 为 false。零依赖近似，足够区分要不要暂停。
func stdinIsTerminal() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

func runWithConfig(_ bootstrap.Config, opts cliOptions, args []string) {
	rules.EnsureHomeRulesDir()

	if len(args) > 0 {
		die("%s", i18n.T("die.directPrompt"))
	}

	if opts.Headless {
		if err := pipelinePipeline(opts, nil); err != nil {
			die("error: %v", err)
		}
		return
	}
	if opts.Prompt != "" || opts.PromptFile != "" {
		die("%s", i18n.T("die.promptNeedsPipeline"))
	}
	// 交互式 TUI 已移除：无子命令、非 headless 时打印用法供用户选择具体功能。
	printTopUsage(os.Stdout)
}

type cliOptions struct {
	providerCallGuard *pipelineProviderCallGuard
	ConfigPath        string
	Dir               string
	Headless          bool
	Prompt            string
	PromptFile        string
	Version           bool
	Update            bool
	UpdateVersion     string
	Help              bool
}

// parseCLIOptions 提取 CLI flag，返回选项和剩余参数。
func parseCLIOptions(argv []string) (cliOptions, []string, error) {
	var opts cliOptions
	var args []string
	for i := 0; i < len(argv); i++ {
		switch argv[i] {
		case "--version", "-v":
			opts.Version = true
		case "version":
			if i+1 < len(argv) {
				return opts, nil, errors.New(i18n.T("flag.versionArgs"))
			}
			opts.Version = true
		case "--help", "-h", "help":
			// 不消费、留在 args 里给 flag 包识别（子命令 --help 触发 fs.Usage）。
			// opts.Help 只用于「无子命令时打顶层 usage」的路由判定。
			opts.Help = true
			args = append(args, argv[i])
		case "update":
			if opts.Update {
				return opts, nil, errors.New(i18n.T("flag.updateOnce"))
			}
			opts.Update = true
			if i+1 < len(argv) {
				if strings.HasPrefix(argv[i+1], "-") {
					return opts, nil, errors.New(i18n.T("flag.updateVersionArg"))
				}
				opts.UpdateVersion = argv[i+1]
				i++
			}
			if i+1 < len(argv) {
				return opts, nil, errors.New(i18n.T("flag.updateVersionArg"))
			}
		case "--config":
			if i+1 >= len(argv) {
				return opts, nil, errors.New(i18n.T("flag.configValue"))
			}
			opts.ConfigPath = argv[i+1]
			i++
		case "--dir":
			// 项目根目录：OutputDir（相对路径时）以它为基准解析，等价于 cd 过去再跑。
			// 子命令（--build-rag/--zero-init/--pipeline 等）由 loadCfgBundle 统一消费。
			if i+1 >= len(argv) {
				return opts, nil, errors.New(i18n.T("flag.dirValue"))
			}
			opts.Dir = argv[i+1]
			i++
		case "--headless":
			opts.Headless = true
		case "--prompt":
			if i+1 >= len(argv) {
				return opts, nil, errors.New(i18n.T("flag.promptValue"))
			}
			opts.Prompt = argv[i+1]
			i++
		case "--prompt-file":
			if i+1 >= len(argv) {
				return opts, nil, errors.New(i18n.T("flag.promptFileValue"))
			}
			opts.PromptFile = argv[i+1]
			i++
		default:
			args = append(args, argv[i])
		}
	}
	if opts.Prompt != "" && opts.PromptFile != "" {
		return opts, nil, errors.New(i18n.T("flag.promptExclusive"))
	}
	if opts.Version && (opts.Update || opts.ConfigPath != "" || opts.Dir != "" || opts.Headless || opts.Prompt != "" || opts.PromptFile != "" || len(args) > 0) {
		return opts, nil, errors.New(i18n.T("flag.versionCombo"))
	}
	if opts.Update && (opts.ConfigPath != "" || opts.Dir != "" || opts.Headless || opts.Prompt != "" || opts.PromptFile != "" || len(args) > 0) {
		return opts, nil, errors.New(i18n.T("flag.updateCombo"))
	}
	return opts, args, nil
}

func versionInfo() buildversion.Info {
	return buildversion.Resolve(buildversion.Info{
		Version: version,
		Commit:  commit,
		Date:    date,
	})
}

func runSelfUpdate(target string) error {
	info := versionInfo()
	result, err := buildversion.Update(context.Background(), buildversion.UpdateOptions{
		Repo:           "prantosmss/Scriptorium",
		BinaryName:     "scriptorium",
		TargetVersion:  target,
		CurrentVersion: info.Version,
	})
	if err != nil {
		return err
	}
	if !result.Updated {
		fmt.Println(i18n.T("update.upToDate", result.Version))
		return nil
	}
	fmt.Println(i18n.T("update.updated", result.Version))
	fmt.Println(i18n.T("update.installPath", result.Path))
	return nil
}

func loadPrompt(opts cliOptions) (string, error) {
	if opts.PromptFile == "" {
		return strings.TrimSpace(opts.Prompt), nil
	}

	var data []byte
	var err error
	if opts.PromptFile == "-" {
		data, err = os.ReadFile("/dev/stdin")
	} else {
		data, err = os.ReadFile(opts.PromptFile)
	}
	if err != nil {
		return "", fmt.Errorf(i18n.T("load.promptRead"), err)
	}
	return strings.TrimSpace(string(data)), nil
}

// hasAnySubcommand 判断 argv 里是否含任一子命令入口 token。用于区分
// 「单纯 --help 想要顶层 usage」与「--diag --help 要子命令 usage」。
func hasAnySubcommand(argv []string) bool {
	for _, a := range argv {
		switch a {
		case "service", "doctor", "skills", "rag", "--review-existing", "--rewrite-existing", "--draft-ai-judge",
			"--check", "--diag", "--simulate", "--import-sim", "--steer",
			"--cocreate", "--pipeline", "--architect-check", "--writing-assets", "--refresh-progress", "--build-rag", "--rag-ready", "--zero-init":
			return true
		}
	}
	return false
}

// hasHelpToken 判断 argv 里是否含 --help/-h/help。各子命令 pipeline 在
// 解析前调用，看到就触发 fs.Usage() 并 return nil，跳过 flag 包的
// 「help requested」错误退出（那种退出还会写 last-error.log + 暂停等回车）。
func hasHelpToken(argv []string) bool {
	for _, a := range argv {
		if a == "--help" || a == "-h" || a == "help" {
			return true
		}
	}
	return false
}

// stripRoutingTokens 从 argv 中移除指定的路由 token。用于「纯路由标记」
// （不期望值的 flag）：Go flag 包会把它注册为 StringVar，下一个 flag token
// 会被吞成它的字符串值，导致后续 --from/--to 等全部被当 positional args。
// 在 main 路由到子命令之前先剥离路由 token，再交给 flag 包解析，flag 包
// 就只看到真正想要值的 --from 等。
func stripRoutingTokens(argv []string, tokens ...string) []string {
	skip := make(map[string]bool, len(tokens))
	for _, t := range tokens {
		skip[t] = true
	}
	out := make([]string, 0, len(argv))
	for _, a := range argv {
		if skip[a] {
			continue
		}
		out = append(out, a)
	}
	return out
}

// printTopUsage 打印顶层 usage。文案在 i18n 目录里，`scriptorium lang` 切语言即整体换掉。
func printTopUsage(w *os.File) {
	fmt.Fprintln(w, i18n.T("usage.title"))
	fmt.Fprintln(w)
	fmt.Fprintln(w, i18n.T("usage.firstRun"))
	fmt.Fprintln(w)
	fmt.Fprintln(w, i18n.T("usage.usage"))
	fmt.Fprintln(w)
	fmt.Fprintln(w, i18n.T("usage.features"))
	fmt.Fprintln(w)
	fmt.Fprintln(w, i18n.T("usage.other"))
	fmt.Fprintln(w)
	fmt.Fprintln(w, i18n.T("usage.subUsage"))
	fmt.Fprintln(w)
	fmt.Fprintln(w, i18n.T("usage.tips"))
	fmt.Fprintln(w)
}

// scanConfigFlag 从 argv 里挑出 --config 的值。放在最前面是为了在任何子命令
// 分发之前拿到配置路径——语言要先于 doctor/service 这类"早拦截"子命令生效。
func scanConfigFlag(argv []string) string {
	for i := 0; i < len(argv); i++ {
		if argv[i] == "--config" && i+1 < len(argv) {
			return argv[i+1]
		}
	}
	return ""
}
