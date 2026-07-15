package isnow

import (
	"errors"
	"testing"
	"time"
)

func TestExclusionExcludesHoliday(t *testing.T) {
	p := mustParse(t, "M-F ! 12/25")
	if p.Canonical() != "*/*/* Monday-Friday *:*:00 ! */12/25 * *:*:*" {
		t.Fatalf("Canonical = %q", p.Canonical())
	}
	cases := map[string]bool{
		"2026-12-25T12:00:00Z": false, // Fri, but Christmas — excluded
		"2026-12-25T23:59:00Z": false, // whole day
		"2026-12-24T12:00:00Z": true,  // Thu, not excluded
	}
	for ts, want := range cases {
		if got := p.Holds(at(t, ts)); got != want {
			t.Fatalf("Holds(%s) = %v, want %v", ts, got, want)
		}
	}
}

func TestExclusionListAndDerivation(t *testing.T) {
	p := mustParse(t, "noon ! 12/25 ! 1/1")
	if p.Holds(at(t, "2027-01-01T12:00:00Z")) {
		t.Fatal("New Year should be excluded")
	}
	got, ok := p.Next(at(t, "2026-12-24T13:00:00Z"))
	if !ok || !got.Equal(at(t, "2026-12-26T12:00:00Z")) { // skips the 25th
		t.Fatalf("Next = %s, %v", got.Format(time.RFC3339), ok)
	}
}

func TestFieldExclusionStillWorks(t *testing.T) {
	// `!12/25` without a separator is a field exclusion, not a pattern exclusion.
	if got := mustParse(t, "M-F !12/25").Canonical(); got != "*/!12/25 Monday-Friday *:*:00" {
		t.Fatalf("Canonical = %q", got)
	}
}

func TestExclusionInvalidSubspec(t *testing.T) {
	if _, err := Parse("noon ! 25"); !errors.Is(err, ErrRange) {
		t.Fatalf("bad exclusion sub-spec: %v", err)
	}
	if _, err := Parse("noon ! 12/5x"); !errors.Is(err, ErrSymbol) {
		t.Fatalf("bad unit in exclusion: %v", err)
	}
	// A sub-spec that fails group mapping (four date slots).
	if _, err := Parse("noon ! ///"); !errors.Is(err, ErrContext) {
		t.Fatalf("bad exclusion group: %v", err)
	}
}
