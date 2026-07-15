package isnow

import (
	"errors"
	"testing"
	"time"
)

func TestIntervalEveryNMinutes(t *testing.T) {
	p := mustParse(t, "+[90mn]")
	if p.Canonical() != "*/*/* * *:*:00 +[90mn]" {
		t.Fatalf("Canonical = %q", p.Canonical())
	}
	holds := map[string]bool{
		"2026-07-14T00:00:00Z": true,
		"2026-07-14T01:30:00Z": true,
		"2026-07-14T22:30:00Z": true,
		"2026-07-14T01:00:00Z": false,
		"2026-07-14T00:00:30Z": false, // not on a minute boundary
	}
	for ts, want := range holds {
		if got := p.Holds(at(t, ts)); got != want {
			t.Fatalf("Holds(%s) = %v, want %v", ts, got, want)
		}
	}
}

func TestIntervalUnits(t *testing.T) {
	// Every unit fires at the start of its grid.
	if !mustParse(t, "+[25h]").Holds(at(t, "2026-07-16T00:00:00Z")) {
		t.Fatal("25h interval should hold")
	}
	if mustParse(t, "+[25h]").Holds(at(t, "2026-07-16T00:30:00Z")) {
		t.Fatal("25h interval must have minute 0")
	}
	if !mustParse(t, "+[30s]").Holds(at(t, "2026-07-14T09:00:30Z")) {
		t.Fatal("30s interval should hold at :30")
	}
	if !mustParse(t, "+[10d]").Holds(at(t, "2026-07-16T00:00:00Z")) {
		t.Fatal("10d interval should hold")
	}
	if mustParse(t, "+[10d]").Holds(at(t, "2026-07-16T12:00:00Z")) {
		t.Fatal("10d interval must be at midnight")
	}
}

func TestIntervalComposesWithFieldsAndBounds(t *testing.T) {
	p := mustParse(t, "M-F +[90mn] >=6 <=18")
	if !p.Holds(at(t, "2026-07-14T07:30:00Z")) { // Tue, on grid, in window
		t.Fatal("should hold")
	}
	if p.Holds(at(t, "2026-07-18T07:30:00Z")) { // Saturday
		t.Fatal("should miss on the weekend")
	}
	if p.Holds(at(t, "2026-07-14T04:30:00Z")) { // on grid but before the window
		t.Fatal("should miss before the since bound")
	}
}

func TestIntervalDerivation(t *testing.T) {
	got, ok := mustParse(t, "+[90mn]").Next(at(t, "2026-07-14T00:00:00Z"))
	if !ok || !got.Equal(at(t, "2026-07-14T01:30:00Z")) {
		t.Fatalf("Next = %s, %v", got.Format(time.RFC3339), ok)
	}
}

func TestIntervalErrors(t *testing.T) {
	if _, err := Parse("-[90mn]"); !errors.Is(err, ErrContext) {
		t.Fatalf("descending interval: %v", err)
	}
	if _, err := Parse("+[0mn]"); !errors.Is(err, ErrRange) {
		t.Fatalf("zero interval: %v", err)
	}
	if _, err := Parse("+[99999999999999999999h]"); !errors.Is(err, ErrRange) {
		t.Fatalf("overflow interval: %v", err)
	}
}

func TestBareWeekUnitIsNotAnInterval(t *testing.T) {
	// A bare `+[3w]` uses the week unit (not an interval grain), so it is not
	// extracted as an interval; it falls through to the (invalid) weekday step.
	if _, err := Parse("+[3w] noon"); !errors.Is(err, ErrContext) {
		t.Fatalf("bare week step: %v", err)
	}
}

func TestDaysFromCivilNegativeEra(t *testing.T) {
	// A day interval at year 0 (January) exercises the pre-epoch era branch.
	p := mustParse(t, "+[1d]")
	if !p.Holds(time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)) {
		t.Fatal("+[1d] holds every midnight, including year 0")
	}
}
