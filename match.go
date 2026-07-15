package isnow

import "time"

// Holds reports whether the isnow holds at the instant: every field constraint
// and every bound is satisfied. The instant is evaluated in its own location
// (at.Location()) and truncated to whole seconds.
func (p Pattern) Holds(at time.Time) bool {
	c := newCtx(at)
	return p.fieldsHold(c) && p.boundsHold(c) && p.intervalsHold(c)
}

func (p Pattern) intervalsHold(c instantCtx) bool {
	for _, iv := range p.intervals {
		if !iv.holds(c) {
			return false
		}
	}
	return true
}

func (p Pattern) fieldsHold(c instantCtx) bool {
	for r := role(0); r < numRoles; r++ {
		if !p.fields[r].holds(c) {
			return false
		}
	}
	return true
}

func (p Pattern) boundsHold(c instantCtx) bool {
	for _, b := range p.bounds {
		if !b.satisfied(c) {
			return false
		}
	}
	return true
}

// newCtx breaks an instant into the fields and derived context matchers read.
func newCtx(at time.Time) instantCtx {
	y, m, d := at.Date()
	dim := daysInMonth(y, int(m), at.Location())
	return instantCtx{
		year: y, month: int(m), day: d,
		hour: at.Hour(), minute: at.Minute(), second: at.Second(),
		weekday:     int(at.Weekday()) + 1,
		daysInMonth: dim, dayOfYear: at.YearDay(),
		occ:        (d-1)/7 + 1,
		occFromEnd: (dim-d)/7 + 1,
	}
}

func daysInMonth(year, month int, loc *time.Location) int {
	return time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, loc).Day()
}
