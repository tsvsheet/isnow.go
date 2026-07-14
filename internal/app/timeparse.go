// Package app holds shared CLI runtime concerns: the injected environment and
// instant/zone parsing used across commands.
package app

import (
	"time"

	"github.com/uplang/isnow.go/internal/constants"
)

// timeLayouts are accepted for instant arguments, interpreted in the evaluation
// zone when no offset is present (specs/contracts/cli.md).
var timeLayouts = []string{
	time.RFC3339,
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05",
	"2006-01-02",
}

// ParseZone resolves an IANA zone name, defaulting to the local zone.
func ParseZone(name string) (*time.Location, error) {
	if name == "" {
		return time.Local, nil
	}
	loc, err := time.LoadLocation(name)
	if err != nil {
		return nil, constants.ErrBadZone.With(err, name)
	}
	return loc, nil
}

// ParseInstant parses an instant in the given zone, trying each accepted layout.
func ParseInstant(value string, loc *time.Location) (time.Time, error) {
	for _, layout := range timeLayouts {
		if t, err := time.ParseInLocation(layout, value, loc); err == nil {
			return t, nil
		}
	}
	return time.Time{}, constants.ErrBadTime.With(nil, value)
}
