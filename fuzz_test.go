package isnow

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

// FuzzParse seeds from every corpus isnow and asserts two invariants: Parse
// never panics, and a successful parse's canonical form is a fixed point
// (parsing the canonical form yields the same canonical form).
func FuzzParse(f *testing.F) {
	for _, src := range seedCorpus(f) {
		f.Add(src)
	}
	f.Add("6")
	f.Add("M,W,F noon")
	f.Add("::+[9] >=6 <=18")
	f.Fuzz(func(t *testing.T, src string) {
		p, err := Parse(src)
		if err != nil {
			return
		}
		canon := p.Canonical()
		reparsed, err := Parse(canon)
		if err != nil {
			t.Fatalf("canonical %q of %q did not reparse: %v", canon, src, err)
		}
		if reparsed.Canonical() != canon {
			t.Fatalf("canonical not a fixed point: %q -> %q", canon, reparsed.Canonical())
		}
	})
}

func seedCorpus(f *testing.F) []string {
	f.Helper()
	files, err := filepath.Glob(filepath.Join(corpusDir, "*.yaml"))
	if err != nil {
		return nil
	}
	var out []string
	for _, path := range files {
		out = append(out, corpusIsnows(f, path)...)
	}
	return out
}

func corpusIsnows(f *testing.F, path string) []string {
	f.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var doc struct {
		Cases []struct {
			Isnow string `yaml:"isnow"`
		} `yaml:"cases"`
	}
	if err := yaml.Unmarshal(data, &doc); err != nil {
		return nil
	}
	out := make([]string, 0, len(doc.Cases))
	for _, c := range doc.Cases {
		if c.Isnow != "" {
			out = append(out, c.Isnow)
		}
	}
	return out
}
