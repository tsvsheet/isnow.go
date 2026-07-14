package isnow

import (
	"errors"
	"testing"
	"time"
)

func at(t *testing.T, s string) time.Time {
	t.Helper()
	ts, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t.Fatalf("bad time %q: %v", s, err)
	}
	return ts
}

func mustParse(t *testing.T, src string) Pattern {
	t.Helper()
	p, err := Parse(src)
	if err != nil {
		t.Fatalf("Parse(%q): %v", src, err)
	}
	return p
}

// --- public API ---

func TestIs(t *testing.T) {
	ok, err := Is("M,W,F noon", at(t, "2026-07-15T12:00:00Z"))
	if err != nil || !ok {
		t.Fatalf("Is = (%v, %v), want (true, nil)", ok, err)
	}
	if _, err := Is("25", time.Now()); !errors.Is(err, ErrRange) {
		t.Fatalf("Is(bad) err = %v, want ErrRange", err)
	}
}

func TestStringMatchesCanonical(t *testing.T) {
	p := mustParse(t, "6")
	if p.String() != p.Canonical() || p.String() != "*/*/* * 06:00:00" {
		t.Fatalf("String=%q Canonical=%q", p.String(), p.Canonical())
	}
}

func TestExplainDeterministic(t *testing.T) {
	p := mustParse(t, "M,W,F noon")
	first := p.Explain()
	if first != mustParse(t, "M,W,F noon").Explain() {
		t.Fatal("Explain not deterministic")
	}
	want := "holds at 12:00:00 on Monday,Wednesday,Friday"
	if first != want {
		t.Fatalf("Explain = %q, want %q", first, want)
	}
}

func TestExplainBounds(t *testing.T) {
	got := mustParse(t, "12 >=2011 <2016").Explain()
	want := "holds at 12:00:00 from 2011 until 2016"
	if got != want {
		t.Fatalf("Explain = %q, want %q", got, want)
	}
}

// --- algebra edges not in the corpus ---

func TestYearArithmeticStep(t *testing.T) {
	yearStep := mustParse(t, "2000+[4]// noon")
	if !yearStep.Holds(at(t, "2004-06-01T12:00:00Z")) {
		t.Fatal("2000+[4] should hold in 2004")
	}
	if yearStep.Holds(at(t, "2005-06-01T12:00:00Z")) {
		t.Fatal("2000+[4] should miss 2005")
	}
}

func TestSpanStep(t *testing.T) {
	p := mustParse(t, "8-12+[2]")
	for _, h := range []struct {
		hour int
		want bool
	}{{8, true}, {9, false}, {10, true}, {12, true}, {13, false}} {
		got := p.Holds(at(t, isoHour(h.hour)))
		if got != h.want {
			t.Fatalf("8-12+[2] at hour %d = %v, want %v", h.hour, got, h.want)
		}
	}
}

func TestWeekdaySpanWraps(t *testing.T) {
	p := mustParse(t, "F-M noon")                // Fri, Sat, Sun, Mon
	if !p.Holds(at(t, "2026-07-19T12:00:00Z")) { // Sunday
		t.Fatal("F-M should hold Sunday")
	}
	if p.Holds(at(t, "2026-07-15T12:00:00Z")) { // Wednesday
		t.Fatal("F-M should miss Wednesday")
	}
}

func TestDescendingStepWithAnchor(t *testing.T) {
	p := mustParse(t, ":30-[10]") // 30,20,10,00
	if !p.Holds(at(t, "2026-07-14T09:10:00Z")) {
		t.Fatal(":30-[10] should hold minute 10")
	}
	if p.Holds(at(t, "2026-07-14T09:40:00Z")) {
		t.Fatal(":30-[10] should miss minute 40")
	}
}

func TestHourFromEnd(t *testing.T) {
	p := mustParse(t, "*/*/* * -3:00:00") // last 3 hours: 21,22,23
	if !p.Holds(at(t, "2026-07-14T23:00:00Z")) {
		t.Fatal("-3 hours should hold 23")
	}
	if p.Holds(at(t, "2026-07-14T20:00:00Z")) {
		t.Fatal("-3 hours should miss 20")
	}
}

func TestWeekdayFromEnd(t *testing.T) {
	// Bare -2 routes to the hour; the weekday from-end needs the explicit slot.
	p := mustParse(t, "*/*/* -2 12:00")          // last 2 weekdays: Fri(6), Sat(7)
	if !p.Holds(at(t, "2026-07-18T12:00:00Z")) { // Saturday
		t.Fatal("-2 weekday should hold Saturday")
	}
	if p.Holds(at(t, "2026-07-15T12:00:00Z")) { // Wednesday
		t.Fatal("-2 weekday should miss Wednesday")
	}
}

func TestMonthFromEnd(t *testing.T) {
	p := mustParse(t, "/-2/ noon") // last 2 months: Nov, Dec
	if !p.Holds(at(t, "2026-12-15T12:00:00Z")) {
		t.Fatal("/-2 months should hold December")
	}
	if p.Holds(at(t, "2026-06-15T12:00:00Z")) {
		t.Fatal("/-2 months should miss June")
	}
}

// --- bound comparators ---

func TestBoundGreaterThanExclusive(t *testing.T) {
	p := mustParse(t, "noon >2011")
	if p.Holds(at(t, "2011-06-01T12:00:00Z")) {
		t.Fatal(">2011 should exclude all of 2011")
	}
	if !p.Holds(at(t, "2012-06-01T12:00:00Z")) {
		t.Fatal(">2011 should include 2012")
	}
}

func TestBoundLessOrEqualInclusive(t *testing.T) {
	p := mustParse(t, "noon <=18")
	if !p.Holds(at(t, "2026-07-14T12:00:00Z")) {
		t.Fatal("<=18 should include noon")
	}
}

// --- derivation edges ---

func TestPrevDaily(t *testing.T) {
	p := mustParse(t, "6")
	got, ok := p.Prev(at(t, "2026-07-14T07:00:00Z"))
	if !ok || !got.Equal(at(t, "2026-07-14T06:00:00Z")) {
		t.Fatalf("Prev = (%s, %v)", got.Format(time.RFC3339), ok)
	}
}

func TestNextNoneWhenPast(t *testing.T) {
	p := mustParse(t, "12 <2016")
	if got, ok := p.Next(at(t, "2020-01-01T00:00:00Z")); ok {
		t.Fatalf("Next past until-bound = %s, want none", got.Format(time.RFC3339))
	}
}

func TestPrevNoneWhenBeforeSince(t *testing.T) {
	p := mustParse(t, "noon >=2027")
	if _, ok := p.Prev(at(t, "2026-01-01T00:00:00Z")); ok {
		t.Fatal("Prev before since-bound should be none")
	}
}

func isoHour(h int) string {
	return "2026-07-14T" + twoDigit(h) + ":00:00Z"
}

func twoDigit(n int) string {
	if n < 10 {
		return "0" + string(rune('0'+n))
	}
	return string(rune('0'+n/10)) + string(rune('0'+n%10))
}
