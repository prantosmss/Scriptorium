package bootstrap

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// readConfigJSONC parses a config the way LoadConfig does (BOM + // comments,
// then strict JSON) so a bad edit fails the test instead of the app.
func readConfigJSONC(t *testing.T, path string) map[string]any {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read config: %v", err)
	}
	var out map[string]any
	if err := json.Unmarshal(stripJSONComments(bytes.TrimPrefix(data, utf8BOM)), &out); err != nil {
		t.Fatalf("config no longer parses after the edit: %v\n---\n%s", err, data)
	}
	return out
}

func writeConfig(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write fixture: %v", err)
	}
	return path
}

// The regression this suite exists for: the first implementation spliced at the
// START of the regex group instead of the end, replacing `  "language": "zh",`
// with `en",` and destroying the config.
func TestSetLanguageInFilePatchesExistingValue(t *testing.T) {
	path := writeConfig(t, "{\n  // keep this comment\n  \"language\": \"zh\",\n  \"provider\": \"llama\"\n}\n")

	if err := SetLanguageInFile(path, "en"); err != nil {
		t.Fatalf("SetLanguageInFile: %v", err)
	}
	cfg := readConfigJSONC(t, path)
	if got := cfg["language"]; got != "en" {
		t.Errorf("language = %v, want en", got)
	}
	if got := cfg["provider"]; got != "llama" {
		t.Errorf("provider = %v, want llama (unrelated keys must survive)", got)
	}
	raw, _ := os.ReadFile(path)
	if n := strings.Count(string(raw), `"language"`); n != 1 {
		t.Errorf(`"language" appears %d times, want exactly 1`, n)
	}
	if !strings.Contains(string(raw), "keep this comment") {
		t.Error("// comment was lost")
	}
}

func TestSetLanguageInFileInsertsWhenMissing(t *testing.T) {
	path := writeConfig(t, "{\n  // 带注释的配置\n  \"provider\": \"llama\",\n  \"model\": \"local-model\"\n}\n")

	if err := SetLanguageInFile(path, "zh"); err != nil {
		t.Fatalf("SetLanguageInFile: %v", err)
	}
	cfg := readConfigJSONC(t, path)
	if got := cfg["language"]; got != "zh" {
		t.Errorf("language = %v, want zh", got)
	}
	if got := cfg["model"]; got != "local-model" {
		t.Errorf("model = %v, want local-model", got)
	}
	raw, _ := os.ReadFile(path)
	if !strings.Contains(string(raw), "带注释的配置") {
		t.Error("// comment was lost")
	}
}

// A comment that merely mentions the key must not be mistaken for the entry:
// the pattern is anchored to the start of a line with only whitespace before it.
func TestSetLanguageInFileIgnoresCommentedKey(t *testing.T) {
	path := writeConfig(t, "{\n  // \"language\": \"zh\" documented here\n  \"provider\": \"llama\"\n}\n")

	if err := SetLanguageInFile(path, "en"); err != nil {
		t.Fatalf("SetLanguageInFile: %v", err)
	}
	readConfigJSONC(t, path)
	raw, _ := os.ReadFile(path)
	if strings.Contains(string(raw), "// \"language\": \"en\"") {
		t.Error("patch rewrote a comment instead of adding a real entry")
	}
	if cfg := readConfigJSONC(t, path); cfg["language"] != "en" {
		t.Errorf("language = %v, want en", cfg["language"])
	}
}

// An empty root object must not end up with a trailing comma (JSON parsers take
// // comments but not trailing commas).
func TestSetLanguageInFileOnEmptyObject(t *testing.T) {
	path := writeConfig(t, "{}")

	if err := SetLanguageInFile(path, "en"); err != nil {
		t.Fatalf("SetLanguageInFile: %v", err)
	}
	if cfg := readConfigJSONC(t, path); cfg["language"] != "en" {
		t.Errorf("language = %v, want en", cfg["language"])
	}
}

func TestSetLanguageInFilePreservesLineEndingsAndBOM(t *testing.T) {
	path := writeConfig(t, "\uFEFF{\r\n  \"language\": \"zh\",\r\n  \"provider\": \"llama\"\r\n}\r\n")

	if err := SetLanguageInFile(path, "en"); err != nil {
		t.Fatalf("SetLanguageInFile: %v", err)
	}
	raw, _ := os.ReadFile(path)
	if !strings.HasPrefix(string(raw), "\uFEFF") {
		t.Error("BOM was dropped")
	}
	if strings.Contains(strings.ReplaceAll(string(raw), "\r\n", ""), "\n") {
		t.Error("bare LF introduced into a CRLF file")
	}
	cfg := readConfigJSONC(t, path)
	if cfg["language"] != "en" {
		t.Errorf("language = %v, want en", cfg["language"])
	}
}

func TestLoadLanguageNeverFails(t *testing.T) {
	dir := t.TempDir()

	if got := LoadLanguage(filepath.Join(dir, "missing.json")); got != "" {
		t.Errorf("missing config language = %q, want empty (defaults to English)", got)
	}

	broken := writeConfig(t, "{ this is not json")
	if got := LoadLanguage(broken); got != "" {
		t.Errorf("broken config language = %q, want empty", got)
	}

	good := writeConfig(t, "{\n  \"language\": \"zh\"\n}\n")
	if got := LoadLanguage(good); got != "zh" {
		t.Errorf("language = %q, want zh", got)
	}
}
