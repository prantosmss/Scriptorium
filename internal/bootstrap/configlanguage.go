package bootstrap

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Language helpers for the interface switch (`scriptorium lang`, dashboard
// toggle). The config file is JSONC: it carries // comments and hand-written
// notes the user cares about, and SaveConfig re-marshals the whole document
// (comments and unknown keys would be destroyed). A language change is a
// single scalar, so it is patched in place instead.

// languageEntryRE matches an existing top-level "language" entry. Anchored to
// the start of a line (with only spaces/tabs before it) so a comment line that
// merely mentions "language": ... is never mistaken for the real entry.
var languageEntryRE = regexp.MustCompile(`(?m)^([ \t]*"language"[ \t]*:[ \t]*")[^"]*(")`)

// ResolveConfigPath returns the file a language change should be written to,
// following the same precedence as LoadConfig: --config, else the project
// config when it exists, else the global config.
func ResolveConfigPath(flagPath string) string {
	if strings.TrimSpace(flagPath) != "" {
		return flagPath
	}
	if p := projectConfigPath(); p != "" {
		if info, err := os.Stat(p); err == nil && !info.IsDir() {
			return p
		}
	}
	return DefaultConfigPath()
}

// LoadLanguage reads just the UI language out of the effective config. It never
// fails: a missing, unreadable or broken config means "use the default", because
// language selection must not block a command that would otherwise work.
func LoadLanguage(flagPath string) string {
	if NeedsSetup(flagPath) {
		return ""
	}
	cfg, err := LoadConfig(flagPath)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(cfg.Language)
}

// SetLanguageInFile writes lang into the "language" key of an existing JSONC
// config, preserving // comments, unknown keys and line endings. An entry that
// is already present is updated in place; otherwise one is inserted at the top
// of the root object.
func SetLanguageInFile(path, lang string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	text := string(data)
	bom := strings.HasPrefix(text, "\uFEFF")
	if bom {
		text = strings.TrimPrefix(text, "\uFEFF")
	}

	updated, changed := patchLanguageValue(text, lang)
	if !changed {
		updated, changed = insertLanguageEntry(text, lang)
	}
	if !changed {
		return errors.New("config file has no JSON object to edit")
	}

	out := updated
	if bom {
		out = "\uFEFF" + out
	}
	mode := os.FileMode(0o600)
	if info, statErr := os.Stat(path); statErr == nil {
		mode = info.Mode().Perm()
	}
	if err := os.WriteFile(path, []byte(out), mode); err != nil {
		return fmt.Errorf("write config: %w", err)
	}
	return nil
}

// patchLanguageValue replaces the value of an existing entry.
func patchLanguageValue(text, lang string) (string, bool) {
	loc := languageEntryRE.FindStringSubmatchIndex(text)
	if loc == nil {
		return text, false
	}
	// loc[0:2] = whole match, loc[2:4] = group 1 (indent + key + colon + the
	// opening quote), loc[4:6] = group 2 (the closing quote). Splice after the
	// END of group 1 and before the START of group 2, i.e. replace only the
	// value between the quotes: text[:loc[3]] + lang + text[loc[4]:].
	updated := text[:loc[3]] + lang + text[loc[4]:]
	return updated, true
}

// insertLanguageEntry adds the entry right after the opening brace of the root
// object, skipping any leading comment lines.
func insertLanguageEntry(text, lang string) (string, bool) {
	start := strings.IndexByte(text, '{')
	if start < 0 {
		return text, false
	}
	eol := "\n"
	if strings.Contains(text, "\r\n") {
		eol = "\r\n"
	}
	rest := text[start+1:]
	// A root object that is empty must not end up with a trailing comma — the
	// JSON parser accepts // comments but not trailing commas.
	trailing := ","
	if onlyCloses(rest) {
		trailing = ""
	}
	insert := eol + `  // UI language: "en" | "zh"` + eol +
		`  "language": "` + lang + `"` + trailing
	return text[:start+1] + insert + rest, true
}

// onlyCloses reports whether the remainder of the object is just (whitespace and
// then) the closing brace.
func onlyCloses(rest string) bool {
	trimmed := strings.TrimSpace(rest)
	return strings.HasPrefix(trimmed, "}")
}
