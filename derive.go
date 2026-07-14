package isnow

import "time"

// horizonYears bounds an unbounded search so a pattern with no occurrence
// terminates (specs/contracts/library-api.md).
const horizonYears = 100

// Next returns the earliest occurrence strictly after from, or false when none
// exists within the pattern's window or the 100-year horizon.
func (p Pattern) Next(from time.Time) (time.Time, bool) {
	return p.derive(from, true)
}

// Prev returns the latest occurrence strictly before from, or false when none.
func (p Pattern) Prev(from time.Time) (time.Time, bool) {
	return p.derive(from, false)
}

func (p Pattern) derive(from time.Time, forward bool) (time.Time, bool) {
	for day := truncateDay(from); !beyondHorizon(from, day, forward); day = day.AddDate(0, 0, direction(forward)) {
		if inst, ok := p.dayMatch(day, from, forward); ok {
			return inst, true
		}
	}
	return time.Time{}, false
}

func direction(forward bool) int {
	if forward {
		return 1
	}
	return -1
}

func beyondHorizon(from, day time.Time, forward bool) bool {
	if forward {
		return day.After(from.AddDate(horizonYears, 0, 0))
	}
	return day.Before(from.AddDate(-horizonYears, 0, 0))
}

func truncateDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// dayMatch finds the first occurrence in day on the correct side of from.
func (p Pattern) dayMatch(day, from time.Time, forward bool) (time.Time, bool) {
	if !p.dateHolds(newCtx(day)) {
		return time.Time{}, false
	}
	insts := p.dayInstants(day)
	if forward {
		return firstAfter(p, insts, from)
	}
	return lastBefore(p, insts, from)
}

func firstAfter(p Pattern, insts []time.Time, from time.Time) (time.Time, bool) {
	for _, inst := range insts {
		if inst.After(from) && p.Holds(inst) {
			return inst, true
		}
	}
	return time.Time{}, false
}

func lastBefore(p Pattern, insts []time.Time, from time.Time) (time.Time, bool) {
	for i := len(insts) - 1; i >= 0; i-- {
		if insts[i].Before(from) && p.Holds(insts[i]) {
			return insts[i], true
		}
	}
	return time.Time{}, false
}

// dateHolds reports whether the date and weekday fields hold for the day.
func (p Pattern) dateHolds(c instantCtx) bool {
	for _, r := range []role{roleYear, roleMonth, roleDay, roleWeekday} {
		if !p.fields[r].holds(c) {
			return false
		}
	}
	return true
}

// dayInstants enumerates the day's matching instants in ascending order, formed
// from the independent hour/minute/second value sets (field-local stepping).
func (p Pattern) dayInstants(day time.Time) []time.Time {
	c := newCtx(day)
	y, m, d := day.Date()
	loc := day.Location()
	var out []time.Time
	for _, h := range p.matching(roleHour, c, 0, 23) {
		for _, mi := range p.matching(roleMinute, c, 0, 59) {
			for _, s := range p.matching(roleSecond, c, 0, 59) {
				out = append(out, time.Date(y, m, d, h, mi, s, 0, loc))
			}
		}
	}
	return out
}

func (p Pattern) matching(r role, c instantCtx, lo, hi int) []int {
	var out []int
	for v := lo; v <= hi; v++ {
		if p.fields[r].holds(c.with(r, v)) {
			out = append(out, v)
		}
	}
	return out
}

// with returns a copy of the context with one field replaced.
func (c instantCtx) with(r role, v int) instantCtx {
	switch r {
	case roleHour:
		c.hour = v
	case roleMinute:
		c.minute = v
	default:
		c.second = v
	}
	return c
}
