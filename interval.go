package isnow

import "fmt"

// An interval is a true periodic recurrence — "every N units" — as distinct from
// a field-local step, which resets within a single field's cycle. An interval
// crosses field boundaries (`+[90mn]` spans hours, `+[25h]` spans days, `+[10d]`
// spans months) and is anchored at the civil epoch (1970-01-01T00:00:00 in the
// evaluation zone), so it needs no explicit start. It fires at the beginning of
// each grid unit: `+[90mn]` at 00:00, 01:30, 03:00, … (second 0).

// intervalUnit is the grain of an interval step.
type intervalUnit int

const (
	unitSecond intervalUnit = iota
	unitMinute
	unitHour
	unitDay
)

// intervalUnits maps a quantity's unit name to an interval grain; only these
// four fill the cross-cycle gap (weeks/months/years are already fields).
var intervalUnits = map[string]intervalUnit{"s": unitSecond, "mn": unitMinute, "h": unitHour, "d": unitDay}

// intervalSpec is a compiled interval: a grain, a stride, and whether it descends.
type intervalSpec struct {
	unit intervalUnit
	n    int
	text string
}

// holds reports whether the instant lands on the interval grid: it must be an
// integer number of periods from the epoch, at the start of its grid unit.
func (iv intervalSpec) holds(c instantCtx) bool {
	switch iv.unit {
	case unitSecond:
		return mod(totalSeconds(c), iv.n) == 0
	case unitMinute:
		return c.second == 0 && mod(totalMinutes(c), iv.n) == 0
	case unitHour:
		return c.minute == 0 && c.second == 0 && mod(totalHours(c), iv.n) == 0
	default:
		return c.hour == 0 && c.minute == 0 && c.second == 0 && mod(totalDays(c), iv.n) == 0
	}
}

// compileInterval validates and compiles an interval increment.
func compileInterval(in *rawIncr) (intervalSpec, error) {
	if in.fromEnd {
		return intervalSpec{}, ErrContext // a descending interval is meaningless
	}
	q := in.qtys[0]
	n, err := positive(q.num)
	if err != nil {
		return intervalSpec{}, err
	}
	return intervalSpec{unit: intervalUnits[q.unit], n: n, text: fmt.Sprintf("+[%d%s]", n, q.unit)}, nil
}

// hasSecondInterval reports whether any interval is second-grained (so the
// second field must default to wildcard rather than 0).
func hasSecondInterval(ins []*rawIncr) bool {
	for _, in := range ins {
		if intervalUnits[in.qtys[0].unit] == unitSecond {
			return true
		}
	}
	return false
}

func totalDays(c instantCtx) int    { return daysFromCivil(c.year, c.month, c.day) }
func totalHours(c instantCtx) int   { return totalDays(c)*24 + c.hour }
func totalMinutes(c instantCtx) int { return totalHours(c)*60 + c.minute }
func totalSeconds(c instantCtx) int { return totalMinutes(c)*60 + c.second }

// daysFromCivil returns the number of days from 1970-01-01 to the civil date
// (Howard Hinnant's algorithm), valid across the proleptic Gregorian calendar.
func daysFromCivil(y, m, d int) int {
	if m <= 2 {
		y--
	}
	era := eraOf(y)
	yoe := y - era*400
	doy := (153*monthShift(m)+2)/5 + d - 1
	doe := yoe*365 + yoe/4 - yoe/100 + doy
	return era*146097 + doe - 719468
}

func eraOf(y int) int {
	if y >= 0 {
		return y / 400
	}
	return (y - 399) / 400
}

func monthShift(m int) int {
	if m > 2 {
		return m - 3
	}
	return m + 9
}
