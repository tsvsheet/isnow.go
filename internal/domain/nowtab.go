package domain

import (
	"strings"

	"github.com/tsvsheet/isnow.go/internal/constants"
)

// ParseNowtab reads nowtab text into entries: each non-blank, non-comment line
// is `<isnow>  <command> [args…]` (the isnow is one whitespace-free token).
func ParseNowtab(text string) ([]Entry, error) {
	var entries []Entry
	for _, line := range strings.Split(text, "\n") {
		entry, ok, err := parseNowtabLine(line)
		if err != nil {
			return nil, err
		}
		if ok {
			entries = append(entries, entry)
		}
	}
	return entries, nil
}

func parseNowtabLine(line string) (Entry, bool, error) {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" || strings.HasPrefix(trimmed, "#") {
		return Entry{}, false, nil
	}
	fields := strings.Fields(trimmed)
	if len(fields) < 2 {
		return Entry{}, false, constants.ErrMissingCommand
	}
	entry, err := CompileEntry(fields[0], fields[1], fields[2:])
	if err != nil {
		return Entry{}, false, err
	}
	return entry, true, nil
}
