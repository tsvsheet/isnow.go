// Package isnow implements the DTimpalr date/time pattern language: an isnow
// holds at an instant when every field constraint is satisfied. The defining
// operation is the membership test, Pattern.Holds(at); Next/Prev derive
// occurrences from it. See github.com/uplang/isnow (SPECIFICATION.md) for the
// language and specs/contracts/ for the pinned semantics.
package isnow

import (
	"errors"

	errs "github.com/gomatic/go-error"
)

// The four stable error codes every implementation shares (specs/contracts/
// semantics.md). Callers match with errors.Is.
const (
	// ErrSyntax is a malformed pattern the grammar rejects.
	ErrSyntax errs.Const = "syntax"
	// ErrSymbol is an unknown or ambiguous weekday/time/unit name.
	ErrSymbol errs.Const = "symbol"
	// ErrRange is a value outside its field's domain.
	ErrRange errs.Const = "range"
	// ErrContext is a grammatical but semantically invalid construct.
	ErrContext errs.Const = "context"
)

// Code returns the stable string code of a library error (syntax, symbol,
// range, context), or "" if err is nil or not a library error.
func Code(err error) string {
	for _, c := range []errs.Const{ErrSyntax, ErrSymbol, ErrRange, ErrContext} {
		if errors.Is(err, c) {
			return string(c)
		}
	}
	return ""
}
