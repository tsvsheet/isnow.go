// Package constants holds the CLI-layer sentinel errors.
package constants

import errs "github.com/gomatic/go-error"

// CLI-layer sentinels, kept sorted. Library errors (ErrSyntax, ErrSymbol,
// ErrRange, ErrContext) come from the root isnow package.
const (
	// ErrBadTime is an unparseable instant argument.
	ErrBadTime errs.Const = "invalid time"
	// ErrBadZone is an unknown IANA time zone.
	ErrBadZone errs.Const = "invalid time zone"
	// ErrMissingCommand is a run entry with no command to execute.
	ErrMissingCommand errs.Const = "missing command"
	// ErrNoOccurrence is a wait/derive with no occurrence in range.
	ErrNoOccurrence errs.Const = "no occurrence"
	// ErrReadTab is a nowtab file that cannot be read.
	ErrReadTab errs.Const = "cannot read nowtab"
	// ErrTimeout is a wait that expired before an occurrence.
	ErrTimeout errs.Const = "timed out"
)
