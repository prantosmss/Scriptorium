// Package i18n holds the interface language switch for novel-studio.
//
// Design rules:
//
//  1. English is the canonical catalog. Every key must exist in messages_en.go.
//  2. The Chinese catalog in messages_zh.go is legacy support: a key that is
//     missing there falls back to English. Dropping Chinese from the app
//     entirely therefore means deleting catalogZH — no call site changes.
//  3. T() never returns an empty string. An unknown key falls back to the key
//     itself, which looks wrong on screen but is always traceable.
//  4. The package is intentionally a global (not passed through the whole call
//     graph): every CLI command prints through fmt/slog deep inside helpers, and
//     threading a language parameter through them would touch thousands of
//     signatures for no runtime benefit.
package i18n

import (
	"fmt"
	"strings"
	"sync"
)

// Lang is a supported interface language.
type Lang string

const (
	// EN is English — the default and the canonical catalog.
	EN Lang = "en"
	// ZH is Chinese (legacy catalog kept until the final removal pass).
	ZH Lang = "zh"
)

var (
	mu      sync.RWMutex
	current = EN
)

// Supported lists the languages the switcher offers, in display order.
func Supported() []Lang { return []Lang{EN, ZH} }

// Valid reports whether value names a supported language.
func Valid(value string) bool {
	_, ok := Parse(value)
	return ok
}

// Parse normalises a user-supplied language name. It accepts the bare codes
// ("en", "zh"), a few common aliases ("english", "cn", "zh-CN", "zh_CN") and
// treats "" as "not specified" (ok = false) so callers can keep their default.
func Parse(value string) (Lang, bool) {
	switch strings.ToLower(strings.TrimSpace(strings.ReplaceAll(value, "_", "-"))) {
	case "":
		return "", false
	case "en", "eng", "english":
		return EN, true
	case "zh", "cn", "chi", "zho", "chinese", "zh-cn", "zh-hans", "zh-hans-cn", "zh-tw", "zh-hant":
		return ZH, true
	default:
		return "", false
	}
}

// Set installs the active language from a raw config/flag value. Unknown or
// empty values fall back to English, so a typo can never blank out the UI.
// It returns the language that is now active.
func Set(value string) Lang {
	lang, ok := Parse(value)
	if !ok {
		lang = EN
	}
	mu.Lock()
	current = lang
	mu.Unlock()
	return lang
}

// Current returns the active language.
func Current() Lang {
	mu.RLock()
	defer mu.RUnlock()
	return current
}

// T looks key up in the active catalog, falling back to English and then to
// the key itself. With arguments the message is run through fmt.Sprintf.
func T(key string, args ...any) string {
	msg, ok := lookup(Current(), key)
	if !ok {
		msg, ok = lookup(EN, key)
	}
	if !ok {
		msg = key
	}
	if len(args) == 0 {
		return msg
	}
	return fmt.Sprintf(msg, args...)
}

func lookup(lang Lang, key string) (string, bool) {
	catalog := catalogEN
	if lang == ZH {
		catalog = catalogZH
	}
	msg, ok := catalog[key]
	return msg, ok
}
