package isnow

type bareRoute int

const (
	routeWeekday bareRoute = iota
	routeTime
	routeNumber
)

// assignBare routes a bare group's field to a role by its content.
func assignBare(s *slots, f rawField, threeForm bool) error {
	route, err := classifyBare(f)
	if err != nil {
		return err
	}
	switch route {
	case routeWeekday:
		return claim(s, roleWeekday, f)
	case routeTime:
		return assignTimeSymbol(s, f)
	default:
		return assignBareNumber(s, f, threeForm)
	}
}

// classifyBare inspects the first atom: a weekday symbol or '*' routes to
// weekday, a time symbol to the time group, a number to hour/weekday.
func classifyBare(f rawField) (bareRoute, error) {
	a := firstAtom(f)
	switch {
	case a == nil || a.star:
		return routeWeekday, nil
	case a.name != "":
		return classifyName(a.name)
	default:
		return routeNumber, nil
	}
}

func classifyName(name string) (bareRoute, error) {
	if _, res := resolveWeekday(name); res == resOne {
		return routeWeekday, nil
	} else if res == resAmbiguous {
		return routeWeekday, ErrSymbol
	}
	switch _, res := resolveTime(name); res {
	case resOne:
		return routeTime, nil
	default:
		return routeTime, ErrSymbol
	}
}

// firstAtom returns the first term's lo atom (nil for an incr-only term); the
// grammar guarantees a present field has at least one term.
func firstAtom(f rawField) *rawAtom {
	return f.terms[0].lo
}

// assignTimeSymbol expands a bare time symbol into exact H:M:S fields. The
// caller (classifyBare) has already confirmed the symbol resolves uniquely.
func assignTimeSymbol(s *slots, f rawField) error {
	hms, _ := resolveTime(f.terms[0].lo.name)
	roles := []role{roleHour, roleMinute, roleSecond}
	for i, r := range roles {
		if err := claim(s, r, *exactRaw(hms[i])); err != nil {
			return err
		}
	}
	return nil
}

// assignBareNumber fills an unconstrained hour slot (nil or present-but-empty).
// If the hour is explicitly constrained, the number is a weekday only in the
// full three-group form; otherwise it is a context error.
func assignBareNumber(s *slots, f rawField, threeForm bool) error {
	if hourFree(s) {
		slot := f
		s[roleHour] = &slot
		return nil
	}
	if threeForm {
		return claim(s, roleWeekday, f)
	}
	return ErrContext
}

func hourFree(s *slots) bool {
	return s[roleHour] == nil || !s[roleHour].present
}

// fillDefaults assigns wildcards and time defaults to unassigned roles. secWild
// makes an unprovided second default to wildcard (a second-grained interval owns
// the second field) rather than to 0.
func fillDefaults(s *slots, secWild, timeWild bool) {
	for _, r := range []role{roleYear, roleMonth, roleDay, roleWeekday} {
		if s[r] == nil {
			s[r] = &wildRaw
		}
	}
	fillTimeDefaults(s, secWild, timeWild)
}

// fillTimeDefaults applies the time-default rule: when no time field is provided
// at all, the time is *:*:00; otherwise unassigned time roles (always finer than
// the coarsest provided one) default to 0.
func fillTimeDefaults(s *slots, secWild, timeWild bool) {
	roles := []role{roleHour, roleMinute, roleSecond}
	if !anyProvided(s, roles) {
		s[roleHour], s[roleMinute], s[roleSecond] = &wildRaw, &wildRaw, defaultSecond(secWild || timeWild)
		return
	}
	for _, r := range roles {
		if s[r] == nil {
			s[r] = exactRaw(0)
		}
	}
}

func defaultSecond(wild bool) *rawField {
	if wild {
		return &wildRaw
	}
	return exactRaw(0)
}

func anyProvided(s *slots, roles []role) bool {
	for _, r := range roles {
		if s[r] != nil {
			return true
		}
	}
	return false
}
