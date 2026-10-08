package i18n

import "testing"

// The Chinese catalog is optional by design (English is canonical), but every
// key it does define must also exist in English — otherwise switching to zh
// would show a raw key.
func TestChineseCatalogIsSubsetOfEnglish(t *testing.T) {
	for key := range catalogZH {
		if _, ok := catalogEN[key]; !ok {
			t.Errorf("key %q exists in the Chinese catalog but not in the English one", key)
		}
	}
}

// A key missing from the active catalog must fall back to English, never to an
// empty string. This is the property that lets us delete catalogZH later.
func TestFallsBackToEnglishThenKey(t *testing.T) {
	defer Set("en")

	Set("zh")
	if got := T("usage.title"); got == "" || got == "usage.title" {
		t.Fatalf("zh lookup of usage.title = %q, want the Chinese title", got)
	}

	catalogEN["test.only.english"] = "english only"
	defer delete(catalogEN, "test.only.english")

	Set("zh")
	if got := T("test.only.english"); got != "english only" {
		t.Errorf("missing zh key = %q, want English fallback", got)
	}
	Set("en")
	if got := T("test.missing.everywhere"); got != "test.missing.everywhere" {
		t.Errorf("unknown key = %q, want the key itself", got)
	}
}

// Formatting arguments must reach the message, and a message without verbs must
// never be run through Sprintf (it would swallow percent signs).
func TestFormatting(t *testing.T) {
	Set("en")
	defer Set("en")

	if got := T("lang.current", EN); got != "current: en" {
		t.Errorf("T(lang.current, en) = %q", got)
	}
	if got := T("usage.title"); got != catalogEN["usage.title"] {
		t.Errorf("argument-less T() altered the message: %q", got)
	}
}

func TestParseAcceptsAliasesAndRejectsJunk(t *testing.T) {
	cases := []struct {
		in   string
		want Lang
		ok   bool
	}{
		{"", "", false},
		{"en", EN, true},
		{" English ", EN, true},
		{"zh", ZH, true},
		{"zh-CN", ZH, true},
		{"zh_CN", ZH, true},
		{"CN", ZH, true},
		{"klingon", "", false},
	}
	for _, c := range cases {
		got, ok := Parse(c.in)
		if got != c.want || ok != c.ok {
			t.Errorf("Parse(%q) = (%q, %v), want (%q, %v)", c.in, got, ok, c.want, c.ok)
		}
	}
}

// An unusable value must not blank out the UI: Set always lands on a real
// language and reports which one.
func TestSetFallsBackToEnglish(t *testing.T) {
	defer Set("en")

	if got := Set("nope"); got != EN {
		t.Errorf("Set(nope) = %q, want en", got)
	}
	if got := Current(); got != EN {
		t.Errorf("Current() = %q, want en", got)
	}
	if got := Set("zh"); got != ZH || Current() != ZH {
		t.Errorf("Set(zh) = %q / Current = %q, want zh", got, Current())
	}
}
