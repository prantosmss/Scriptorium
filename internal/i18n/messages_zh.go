package i18n

// catalogZH is the legacy Chinese catalog. It is allowed to be incomplete:
// i18n.T falls back to catalogEN for any key that is missing here, which is
// what makes the final "remove Chinese altogether" pass a one-line deletion
// (drop this file and its lookup branch).
var catalogZH = map[string]string{
	// ---- 顶层用法（printTopUsage）----
	"usage.title": "novel-studio — AI 长篇小说创作引擎",

	"usage.firstRun": `首次运行:
  1. novel-studio doctor       # 检查本机环境并给出修复建议
  2. novel-studio              # 创建配置（仅首次需要）
  3. novel-studio --check      # 发起最小真实模型请求`,

	"usage.usage": `用法:
  novel-studio --pipeline --prompt <text>     # 可恢复流水线：设计→按弧推演→逐章渲染审核
  novel-studio --pipeline --prompt-file p.md  # 从文件读 prompt 后进入流水线
  novel-studio --cocreate                     # 多轮对话澄清需求，定稿创作指令
  novel-studio --headless --prompt <text>     # 兼容别名：内部转为 --pipeline`,

	"usage.features": `功能子命令（无 TTY、CI / 远程可用）:
  novel-studio --check                        # LLM 连通性自检（先确认能用再创作）
  novel-studio --pipeline --stages review     # 逐章 Editor 评审（不改原文）
  novel-studio --draft-ai-judge --chapter N   # 对当前草稿做独立 DeepSeek 裸正文预审
  novel-studio --pipeline --stages rewrite    # 按评审反馈逐章 Writer 重写
  novel-studio --diag                         # 诊断当前项目产物
  novel-studio --writing-assets list          # 查看/启停/组合/绑定/试写写法资产
  novel-studio --writing-assets seed-defaults # 初始化本书基础写法资产
  novel-studio --refresh-progress [--dir d]   # 回填章节推进/人物变化/下一章计划台账
  novel-studio --build-rag [--dir d]          # 构建本书 RAG 索引并可探测召回
  novel-studio --rag-ready [--dir d]          # 只修复/验证 RAG，不启动写作
  novel-studio rag audit [--root data/runs]   # 审计全部主索引和历史 RAG 快照
  novel-studio rag maintain --apply           # 备份、整理主索引并压缩重复快照
  novel-studio --architect-check [--dir d]    # 检查 Architect foundation，通过后才允许 zero-init
  novel-studio --zero-init [--dir d]          # 新书第一章前的角色/关系/资源推演资产
  novel-studio eval inspect --cases evals/cases/harness # Harness 检查既有项目产物
  novel-studio --simulate [--no-diag]         # 分析 simulate/ 语料合成仿写画像
  novel-studio --import-sim <profile.json>    # 导入此前生成的仿写画像（默认写 diag）
  novel-studio --steer "<指令>"               # 排队一条干预，下次启动生效`,

	"usage.other": `其它:
  novel-studio doctor                         # 本地环境/配置/看板前置检查（不调用模型）
  novel-studio service start                  # 启动浏览器进度看板（长篇 output/novel + 短篇服务）
  novel-studio service open                   # 手动打开小说项目进度看板
  novel-studio service status                 # 检查看板服务 /api/health
  novel-studio skills list                    # 列出内置 skills
  novel-studio skills export --to <dir>       # 导出 skills 到项目目录
  novel-studio lang [en|zh]                   # 查看/切换界面语言
  novel-studio --version                      # 打印版本信息
  novel-studio update [version]               # 自我更新
  novel-studio --config <path>                # 用指定配置文件启动
  novel-studio --dir <project>                # 指定项目根目录（OutputDir 基准），免 cd`,

	"usage.subUsage": `每个子命令的专属选项：
  novel-studio service --help
  novel-studio --pipeline --help
  novel-studio --review-existing --help      # 兼容别名
  novel-studio --rewrite-existing --help     # 兼容别名
  novel-studio skills --help`,

	"usage.tips": `提示:
  · 配置默认读 ~/.novel-studio/config.json（项目内可用 ./.novel-studio/config.json 覆盖）
  · 首次启动会跑 setup 引导：选 Provider / 填 Key / 填 Base URL / 填模型
  · 章节输出在 output/novel/chapters/*.md（可在配置里改 OutputDir）`,

	// ---- CLI 参数错误（parseCLIOptions）----
	"flag.versionArgs":      "version 不接受参数",
	"flag.updateOnce":       "update 只能指定一次",
	"flag.updateVersionArg": "update 只接受一个可选版本参数",
	"flag.configValue":      "--config 缺少值",
	"flag.dirValue":         "--dir 缺少值",
	"flag.promptValue":      "--prompt 缺少值",
	"flag.promptFileValue":  "--prompt-file 缺少值",
	"flag.promptExclusive":  "--prompt 和 --prompt-file 不能同时使用",
	"flag.versionCombo":     "version 不能与其他启动参数混用",
	"flag.updateCombo":      "update 不能与其他启动参数混用",

	// ---- 致命错误（die）----
	"die.headlessNoSetup":   "error: headless 模式不支持首次引导，请先在交互终端运行一次 novel-studio 完成配置，或手写配置文件",
	"die.directPrompt":      "error: 不支持命令行直接传入小说需求，请用 --pipeline --prompt <文本> 或对应子命令",
	"die.promptNeedsPipeline": "error: --prompt/--prompt-file 需要配合 --pipeline 使用",
	"die.loggedAt":          "（详细错误已记录到 %s）",
	"die.pressEnter":        "\n按回车键退出...",
	"load.promptRead":       "读取 prompt 失败: %w",

	// ---- 自我更新 ----
	"update.upToDate":    "novel-studio 已是最新版本 %s",
	"update.updated":     "novel-studio 已更新到 %s",
	"update.installPath": "安装位置：%s",

	// ---- doctor ----
	"doctor.fix":          "  修复：%s",
	"doctor.resultReady":  "\n结果：本地运行前置条件已就绪。下一步运行 novel-studio --check 验证真实模型连接。",
	"doctor.resultIssues": "\n结果：存在必须修复的项目；按上方建议处理后重新运行 doctor。",

	// ---- --check ----
	"check.usage": "用法: novel-studio --check [--timeout 30s] [--provider <name> --model <model>]\n\n对默认模型与各角色模型做一次最小真实调用，逐一报告是否可用。\n指定 --provider/--model 时只测该目标（不改配置），用于验证某个备用 provider。\n\n选项：\n",
	"check.timeout":         "单次连通性调用的超时",
	"check.provider":        "只测指定 provider（配置里 providers 的 key 名），需配 --model",
	"check.model":           "配合 --provider 指定要测的模型名",
	"check.unknownArgs":     "--check 不接受额外参数：%v",
	"check.pairRequired":    "--provider 和 --model 必须同时指定",
	"check.noConfig":        "尚未配置，请先在交互终端运行一次 novel-studio 完成配置引导，或手写配置文件",
	"check.loadConfig":      "加载配置失败: %w",
	"check.unknownProvider": "配置里没有名为 %q 的 provider",
	"check.buildModel":      "构建模型失败（配置层就不通，无需联网即失败）: %w",
	"check.targets":         "[check] 共 %d 个待测模型目标（含兜底，超时 %s/个）\n\n",
	"check.byRole":          "\n按角色汇总：",
	"check.primaryOK":       "  ✓ %s：主模型可用",
	"check.fallbackOK":      "  ⚠ %s：主模型不可用，已可走兜底",
	"check.defaultDown":     "  ⚠ %s：默认模型不可用（仅影响未配置兜底的辅助路径，如共创）",
	"check.allDown":         "  ✗ %s：主与兜底均不可用",
	"check.noUsableModel":   "以下角色无任何可用模型：%s（常见原因：代理未启动 / api_key 失效 / base_url 写错）",
	"check.degraded":        "[check] 可创作（降级）：%s 将走兜底；如需主模型请启动其 provider。\n",
	"check.allOK":           "[check] 全部角色主模型可用 ✓",
	"check.emptyResponse":   "模型返回空响应（连接通但无内容，疑似代理/模型名问题）",
	"check.label.primary":   "主",
	"check.label.fallback":  "兜底",

	// ---- service ----
	"service.staleReusable": "[dashboard] 旧服务不可复用（创作继续）：%v\n",
	"service.startFailed":   "[dashboard] 启动失败（创作继续）：%v\n",
	"service.pythonNeeded":  "需要 Python 3.9+ 才能启动看板；请安装 Python 3 后重试，或运行 novel-studio doctor 查看修复建议",
	"service.pythonProbe":   "无法读取 Python 版本（%s）：%w",
	"service.pythonOld":     "看板需要 Python 3.9+，当前为 %d.%d（%s）",

	// ---- lang ----
	"lang.header":      "界面语言",
	"lang.current":     "当前：%s",
	"lang.available":   "可选：%s",
	"lang.hint":        "切换：novel-studio lang en|zh",
	"lang.changed":     "界面语言已设为 %s（%s），已写入 %s",
	"lang.unchanged":   "界面语言已经是 %s",
	"lang.unknownLang": "未知语言 %q，可选值：%s",
	"lang.usage":       "用法:\n  novel-studio lang          # 查看当前语言\n  novel-studio lang en|zh    # 切换语言（写入配置）\n",
	"lang.noConfig":    "未找到配置文件；请先在交互终端运行一次 novel-studio 完成配置，再切换语言",
	"lang.writeFailed": "写入配置失败：%v",
}
