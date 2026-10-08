package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/prantosmss/Scriptorium/internal/bootstrap"
	"github.com/prantosmss/Scriptorium/internal/i18n"
)

// runLangCommand implements `scriptorium lang [en|zh]`:
//
//	scriptorium lang          # show the current language
//	scriptorium lang en|zh    # switch it, writing "language" into the config
//
// Like doctor/service it is intercepted before the regular flag parse, so it
// picks --config out of argv itself. The write is a surgical edit
// (bootstrap.SetLanguageInFile) because config.json is JSONC: re-marshalling it
// would throw away the user's // comments.
//
// Exit codes: 0 success, 1 could not write / no config, 2 usage error.
func runLangCommand(rawArgs []string, configPath string) int {
	argv := stripConfigTokens(rawArgs)

	for _, a := range argv {
		if a == "--help" || a == "-h" || a == "help" {
			fmt.Fprint(os.Stdout, i18n.T("lang.usage"))
			return 0
		}
	}

	path := bootstrap.ResolveConfigPath(configPath)

	if len(argv) == 0 {
		fmt.Println(i18n.T("lang.header"))
		fmt.Println(i18n.T("lang.current", i18n.Current()))
		fmt.Println(i18n.T("lang.available", supportedLangs()))
		fmt.Println(i18n.T("lang.hint"))
		return 0
	}
	if len(argv) > 1 {
		fmt.Fprint(os.Stderr, i18n.T("lang.usage"))
		return 2
	}

	lang, ok := i18n.Parse(argv[0])
	if !ok {
		fmt.Fprintln(os.Stderr, i18n.T("lang.unknownLang", argv[0], supportedLangs()))
		fmt.Fprint(os.Stderr, i18n.T("lang.usage"))
		return 2
	}
	if info, err := os.Stat(path); err != nil || info.IsDir() {
		fmt.Fprintln(os.Stderr, i18n.T("lang.noConfig"))
		return 1
	}
	if err := bootstrap.SetLanguageInFile(path, string(lang)); err != nil {
		fmt.Fprintln(os.Stderr, i18n.T("lang.writeFailed", err))
		return 1
	}

	// Confirm in the language the user just switched to, so the reply is a live
	// demonstration that the switch took effect.
	if i18n.Current() == lang {
		fmt.Println(i18n.T("lang.unchanged", lang))
		return 0
	}
	i18n.Set(string(lang))
	fmt.Println(i18n.T("lang.changed", lang, langDisplay(lang), path))
	return 0
}

// supportedLangs formats the switcher's choices for display.
func supportedLangs() string {
	parts := make([]string, 0, len(i18n.Supported()))
	for _, l := range i18n.Supported() {
		parts = append(parts, string(l)+" ("+langDisplay(l)+")")
	}
	return strings.Join(parts, ", ")
}

// langDisplay is the human name of a language, used in confirmations.
func langDisplay(l i18n.Lang) string {
	if l == i18n.ZH {
		return "中文"
	}
	return "English"
}

// stripConfigTokens removes a `--config <path>` pair so it is not mistaken for
// the requested language.
func stripConfigTokens(argv []string) []string {
	out := make([]string, 0, len(argv))
	for i := 0; i < len(argv); i++ {
		if argv[i] == "--config" {
			if i+1 < len(argv) {
				i++
			}
			continue
		}
		out = append(out, argv[i])
	}
	return out
}
