package i18n

// catalogEN is the canonical interface catalog. Every key used anywhere in the
// app must be present here; messages_zh.go may omit keys (English fallback).
//
// Keys are grouped by the command that prints them: usage.* (top-level help),
// flag.* (CLI argument errors), die.* (fatal errors), update.*, doctor.*,
// check.*, service.*, lang.*.
var catalogEN = map[string]string{
	// ---- top-level usage (printTopUsage) ----
	"usage.title": "novel-studio — AI long-form fiction engine",

	"usage.firstRun": `First run:
  1. novel-studio doctor       # check the local environment and get fix suggestions
  2. novel-studio              # create the config (first time only)
  3. novel-studio --check      # make one tiny real model request`,

	"usage.usage": `Usage:
  novel-studio --pipeline --prompt <text>     # resumable pipeline: design -> arc rehearsal -> per-chapter render + review
  novel-studio --pipeline --prompt-file p.md  # read the prompt from a file, then enter the pipeline
  novel-studio --cocreate                     # multi-turn clarification, finalise the writing brief
  novel-studio --headless --prompt <text>     # legacy alias: converted to --pipeline internally`,

	"usage.features": `Feature subcommands (no TTY; CI / remote friendly):
  novel-studio --check                        # LLM connectivity check (verify it works before writing)
  novel-studio --pipeline --stages review     # per-chapter Editor review (does not edit the text)
  novel-studio --draft-ai-judge --chapter N   # independent DeepSeek raw-prose pre-review of the current draft
  novel-studio --pipeline --stages rewrite    # per-chapter Writer rewrite driven by review notes
  novel-studio --diag                         # diagnose the current project artefacts
  novel-studio --writing-assets list          # list / toggle / combine / bind / trial writing assets
  novel-studio --writing-assets seed-defaults # seed this book's baseline writing assets
  novel-studio --refresh-progress [--dir d]   # backfill chapter progress / character changes / next-chapter plan ledger
  novel-studio --build-rag [--dir d]          # build this book's RAG index and probe recall
  novel-studio --rag-ready [--dir d]          # repair / verify RAG only, do not start writing
  novel-studio rag audit [--root data/runs]   # audit every main index and historical RAG snapshot
  novel-studio rag maintain --apply           # back up, compact the main index, merge duplicate snapshots
  novel-studio --architect-check [--dir d]    # check the Architect foundation (required before zero-init)
  novel-studio --zero-init [--dir d]          # character / relationship / resource rehearsal assets before chapter 1
  novel-studio eval inspect --cases evals/cases/harness # inspect existing project artefacts with the harness
  novel-studio --simulate [--no-diag]         # analyse the simulate/ corpus into imitation profiles
  novel-studio --import-sim <profile.json>    # import a previously generated profile (writes diag by default)
  novel-studio --steer "<instruction>"        # queue an intervention, applied on the next start`,

	"usage.other": `Other:
  novel-studio doctor                         # local environment / config / dashboard pre-flight (no model calls)
  novel-studio service start                  # start the browser progress dashboard (novel output/novel + short-story service)
  novel-studio service open                   # open the project progress dashboard manually
  novel-studio service status                 # check the dashboard service /api/health
  novel-studio skills list                    # list built-in skills
  novel-studio skills export --to <dir>       # export skills into a project directory
  novel-studio lang [en|zh]                   # show or change the interface language
  novel-studio --version                      # print version information
  novel-studio update [version]               # self-update
  novel-studio --config <path>                # start with a specific config file
  novel-studio --dir <project>                # project root (OutputDir base) — no need to cd`,

	"usage.subUsage": `Per-command options:
  novel-studio service --help
  novel-studio --pipeline --help
  novel-studio --review-existing --help      # legacy alias
  novel-studio --rewrite-existing --help     # legacy alias
  novel-studio skills --help`,

	"usage.tips": `Tips:
  · config is read from ~/.novel-studio/config.json (a project-level ./.novel-studio/config.json overrides it)
  · the first run starts the setup wizard: pick a provider / paste the key / set the base URL / set the model
  · chapters are written to output/novel/chapters/*.md (OutputDir can be changed in the config)`,

	// ---- CLI argument errors (parseCLIOptions) ----
	"flag.versionArgs":     "version does not take arguments",
	"flag.updateOnce":      "update can be given at most once",
	"flag.updateVersionArg": "update accepts at most one optional version argument",
	"flag.configValue":     "--config is missing its value",
	"flag.dirValue":        "--dir is missing its value",
	"flag.promptValue":     "--prompt is missing its value",
	"flag.promptFileValue": "--prompt-file is missing its value",
	"flag.promptExclusive": "--prompt and --prompt-file cannot be used together",
	"flag.versionCombo":    "version cannot be combined with other startup flags",
	"flag.updateCombo":     "update cannot be combined with other startup flags",

	// ---- fatal errors (die) ----
	"die.headlessNoSetup":  "error: first-run setup is not supported in headless mode; run `novel-studio` once in an interactive terminal, or write the config file by hand",
	"die.directPrompt":     "error: passing a story request directly on the command line is not supported; use --pipeline --prompt <text> or the matching subcommand",
	"die.promptNeedsPipeline": "error: --prompt/--prompt-file requires --pipeline",
	"die.loggedAt":         "(full error written to %s)",
	"die.pressEnter":       "\nPress Enter to exit...",
	"load.promptRead":      "reading the prompt failed: %w",

	// ---- self update ----
	"update.upToDate":   "novel-studio is already at the latest version %s",
	"update.updated":    "novel-studio updated to %s",
	"update.installPath": "installed at: %s",

	// ---- doctor ----
	"doctor.fix":         "  fix: %s",
	"doctor.resultReady": "\nResult: local pre-flight requirements are ready. Next run `novel-studio --check` to verify the real model connection.",
	"doctor.resultIssues": "\nResult: items must be fixed; follow the suggestions above and run doctor again.",

	// ---- --check ----
	"check.usage": "Usage: novel-studio --check [--timeout 30s] [--provider <name> --model <model>]\n\nSends the smallest possible real request to every configured model to prove the\npath works before writing starts.\n\n  --provider/--model are optional (default: the config's default model), and\n  require each other.\n\nOptions:\n",
	"check.timeout":       "per-model timeout for the test request",
	"check.provider":      "provider to test (a key under providers), requires --model",
	"check.model":         "model to test alongside --provider",
	"check.unknownArgs":   "--check does not take arguments: %v",
	"check.pairRequired":  "--provider and --model must be given together",
	"check.noConfig":      "no config yet; run `novel-studio` once to create it, or write the config by hand",
	"check.loadConfig":    "loading config: %w",
	"check.unknownProvider": "unknown provider %q in the config",
	"check.buildModel":    "building the model failed (fails before any network call): %w",
	"check.targets":       "[check] %d model target(s) to test (including fallbacks, timeout %s each)\n\n",
	"check.byRole":        "\nAvailability by role:",
	"check.primaryOK":     "  ✓ %s: primary model available",
	"check.fallbackOK":    "  ⚠ %s: primary unavailable, fallback available",
	"check.defaultDown":   "  ⚠ %s: default model unavailable (only affects auxiliary paths with no configured fallback, e.g. co-create)",
	"check.allDown":       "  ✗ %s: neither primary nor fallback is available",
	"check.noUsableModel": "no usable model for these roles: %s (common causes: proxy down / api_key invalid / base_url wrong)",
	"check.degraded":      "[check] writable (degraded): %s will use the fallback; start its provider to use the primary model.\n",
	"check.allOK":         "[check] primary model available for every role ✓",
	"check.emptyResponse": "the model returned an empty response (the connection works but there is no content — likely a proxy or model-name problem)",
	"check.label.primary": "primary",
	"check.label.fallback": "fallback",

	// ---- service ----
	"service.staleReusable": "[dashboard] old service could not be reused (writing continues): %v\n",
	"service.startFailed":   "[dashboard] failed to start (writing continues): %v\n",
	"service.pythonNeeded":  "Python 3.9+ is required for the dashboard; install Python 3 or run `novel-studio doctor` for fix suggestions",
	"service.pythonProbe":   "probing the Python interpreter failed (%s): %w",
	"service.pythonOld":     "Python 3.9+ is required, found %d.%d (%s)",

	// ---- lang ----
	"lang.header":     "Interface language",
	"lang.current":    "current: %s",
	"lang.available":  "available: %s",
	"lang.hint":       "change with: novel-studio lang en|zh",
	"lang.changed":    "interface language set to %s (%s) — written to %s",
	"lang.unchanged":  "interface language is already %s",
	"lang.unknownLang": "unknown language %q — use one of: %s",
	"lang.usage":      "Usage:\n  novel-studio lang          # show the current language\n  novel-studio lang en|zh    # switch the language (saved in the config)\n",
	"lang.noConfig":   "no config file found; run `novel-studio` once to create it, then switch the language",
	"lang.writeFailed": "writing the config failed: %v",
}
