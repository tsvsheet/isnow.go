package command

import (
	"errors"
	"strings"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/tsvsheet/isnow.go/internal/constants"
)

func TestCompletionEmitsScripts(t *testing.T) {
	for _, shell := range []string{"bash", "zsh", "fish"} {
		h := newHarness()
		if code := h.run("completion", shell); code != 0 {
			t.Fatalf("completion %s: exit %d, stderr %q", shell, code, h.err.String())
		}
		if !strings.Contains(h.out.String(), "isnow") {
			t.Fatalf("completion %s script does not name the program: %q", shell, h.out.String())
		}
	}
}

func TestCompletionMissingShell(t *testing.T) {
	h := newHarness()
	if code := h.run("completion"); code != 2 {
		t.Fatalf("missing shell: exit %d, want 2 (stderr %q)", code, h.err.String())
	}
	if !strings.Contains(h.err.String(), string(constants.ErrMissingShell)) {
		t.Fatalf("missing shell diagnostic: %q", h.err.String())
	}
}

func TestCompletionUnsupportedShell(t *testing.T) {
	h := newHarness()
	if code := h.run("completion", "powershell"); code != 2 {
		t.Fatalf("unsupported shell: exit %d, want 2 (stderr %q)", code, h.err.String())
	}
	if !strings.Contains(h.err.String(), string(constants.ErrUnsupportedShell)) {
		t.Fatalf("unsupported shell diagnostic: %q", h.err.String())
	}
}

// TestManEmitsRoff proves the man command emits a well-formed roff manual
// page — title header, NAME, the root pattern synopsis, GLOBAL OPTIONS, and
// one .SS subsection per subcommand — dispatched through the command wiring.
func TestManEmitsRoff(t *testing.T) {
	h := newHarness()
	if code := h.run("man"); code != 0 {
		t.Fatalf("man: exit %d, stderr %q", code, h.err.String())
	}
	out := h.out.String()
	for _, want := range []string{
		`.TH ISNOW 1`,
		".SH NAME\nisnow \\- match instants against isnow date/time patterns",
		".SH SYNOPSIS",
		".SH GLOBAL OPTIONS",
		".SH COMMANDS",
		".SS \"completion\"\n.B isnow completion",
		".SS \"man\"\n.B isnow man",
		".SS \"next\"\n.B isnow next",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("man output missing %q: %q", want, out[:min(len(out), 200)])
		}
	}
}

// TestManFormatting pins the formatting contract details: flags are .TP
// tagged paragraphs carrying their defaults, and the injected help
// command/flag noise is absent.
func TestManFormatting(t *testing.T) {
	h := newHarness()
	if code := h.run("man"); code != 0 {
		t.Fatalf("man: exit %d, stderr %q", code, h.err.String())
	}
	out := h.out.String()
	for _, want := range []string{
		".TP\n\\fB\\-\\-tz\\fR=\\fIvalue\\fR\n",
		"(default: \":8601\")",
		".EX\n  isnow man | man -l -",
	} {
		if !strings.Contains(out, want) {
			t.Fatalf("man output missing %q", want)
		}
	}
	for _, reject := range []string{"Shows a list of commands", `.SS "help`, `\-\-help`} {
		if strings.Contains(out, reject) {
			t.Fatalf("man output contains help noise %q", reject)
		}
	}
}

// TestManPageBare proves a bare command — no args form, no flags, no
// subcommands — renders only the header, synopsis, and description skeleton.
func TestManPageBare(t *testing.T) {
	page := manPage(&cli.Command{Name: "app", Version: "v9.9.9"})
	if !strings.Contains(page, `.TH APP 1 "" "app v9.9.9" "User Commands"`) {
		t.Fatalf("versioned title header missing: %q", page)
	}
	for _, reject := range []string{".br", ".SH GLOBAL OPTIONS", ".SH COMMANDS"} {
		if strings.Contains(page, reject) {
			t.Fatalf("bare man page contains %q: %q", reject, page)
		}
	}
}

// TestRoffLine pins the roff escaping contract: backslashes become printable
// \e, and lines that would read as roff requests are neutralized with \&.
func TestRoffLine(t *testing.T) {
	for in, want := range map[manLine]string{
		`a \ b`:       `a \e b`,
		".request":    `\&.request`,
		"  .indented": `\&  .indented`,
		"'quote":      `\&'quote`,
		"plain":       "plain",
	} {
		if got := roffLine(in); got != want {
			t.Fatalf("roffLine(%q) = %q, want %q", in, got, want)
		}
	}
}

// TestFlagDefault covers the default-text resolution: hidden defaults render
// nothing, an explicit DefaultText wins, a value-taking flag falls back to
// its initial value, and a bool flag shows no default.
func TestFlagDefault(t *testing.T) {
	if got := flagDefault(&cli.StringFlag{Name: "a", Value: "x", HideDefault: true}); got != "" {
		t.Fatalf("hidden default rendered %q", got)
	}
	if got := flagDefault(&cli.StringFlag{Name: "a", Value: "x", DefaultText: "why"}); got != "why" {
		t.Fatalf("explicit DefaultText rendered %q", got)
	}
	if got := flagDefault(&cli.StringFlag{Name: "a", Value: "x"}); got != `"x"` {
		t.Fatalf("value fallback rendered %q", got)
	}
	if got := flagDefault(&cli.BoolFlag{Name: "a"}); got != "" {
		t.Fatalf("bool default rendered %q", got)
	}
}

// bareFlag implements cli.Flag but not cli.DocGenerationFlag, proving the
// documentation filter drops flags that cannot describe themselves.
type bareFlag struct{}

func (bareFlag) String() string           { return "" }
func (bareFlag) Get() any                 { return nil }
func (bareFlag) PreParse() error          { return nil }
func (bareFlag) PostParse() error         { return nil }
func (bareFlag) Set(string, string) error { return nil }
func (bareFlag) Names() []string          { return []string{"bare"} }
func (bareFlag) IsSet() bool              { return false }

// TestDocumentedFlags proves the filter keeps documentable flags and drops
// the injected help flag and non-documentable implementations.
func TestDocumentedFlags(t *testing.T) {
	flags := []cli.Flag{
		&cli.BoolFlag{Name: "help", Aliases: []string{"h"}},
		bareFlag{},
		&cli.StringFlag{Name: "keep"},
	}
	docs := documentedFlags(flags)
	if len(docs) != 1 || docs[0].names[0] != "keep" {
		t.Fatalf("documentedFlags kept %+v, want only keep", docs)
	}
}

// TestManProse covers the block transitions: paragraph, example, blank
// separator, and the example-to-paragraph return edge.
func TestManProse(t *testing.T) {
	want := ".PP\npara one\ncontinued\n.EX\n  example\n.EE\n.PP\nback to prose\n.EX\n  trailing example\n.EE\n"
	if got := manProse("para one\ncontinued\n\n  example\nback to prose\n\n  trailing example"); got != want {
		t.Fatalf("manProse = %q, want %q", got, want)
	}
}

// failWriter always errors, exercising the man write-failure path.
type failWriter struct{}

func (failWriter) Write([]byte) (int, error) { return 0, errors.New("write failed") }

func TestManWriteError(t *testing.T) {
	if err := runMan(failWriter{}, Root(newHarness().env)); !errors.Is(err, constants.ErrManPage) {
		t.Fatalf("write failure: %v, want ErrManPage", err)
	}
}
