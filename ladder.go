package isnow

// slots holds the raw field assigned to each role; nil means unassigned.
type slots [numRoles]*rawField

var wildRaw = rawField{present: false}

// exactRaw builds a present field of a single exact value (for defaults and time
// symbols).
func exactRaw(v int) *rawField {
	return &rawField{present: true, terms: []rawTerm{{lo: &rawAtom{qtys: []rawQty{{num: v, digits: 1}}}}}}
}

// mapGroups applies the shorthand ladder: it assigns each group's fields to the
// seven roles, then fills defaults (specs/contracts/semantics.md §Ladder).
func mapGroups(groups []rawGroup, secWild bool) (slots, error) {
	var s slots
	hasDate, hasTime := kinds(groups)
	if err := assignGroups(&s, groups, hasDate && hasTime); err != nil {
		return s, err
	}
	fillDefaults(&s, secWild)
	return s, nil
}

func kinds(groups []rawGroup) (bool, bool) {
	date, time := false, false
	for _, gr := range groups {
		switch gr.kind {
		case dateKind:
			date = true
		case timeKind:
			time = true
		}
	}
	return date, time
}

// assignGroups assigns date and time groups before bare groups, so a bare
// number sees whether the hour is already explicitly constrained.
func assignGroups(s *slots, groups []rawGroup, threeForm bool) error {
	if err := assignPass(s, groups, false, threeForm); err != nil {
		return err
	}
	return assignPass(s, groups, true, threeForm)
}

func assignPass(s *slots, groups []rawGroup, bareOnly, threeForm bool) error {
	for _, gr := range groups {
		if (gr.kind == bareKind) != bareOnly {
			continue
		}
		if err := assignGroup(s, gr, threeForm); err != nil {
			return err
		}
	}
	return nil
}

func assignGroup(s *slots, gr rawGroup, threeForm bool) error {
	switch gr.kind {
	case dateKind:
		return assignDate(s, gr)
	case timeKind:
		return assignTime(s, gr)
	default:
		return assignBare(s, gr.slots[0], threeForm)
	}
}

func assignDate(s *slots, gr rawGroup) error {
	roles, err := dateRoles(len(gr.slots))
	if err != nil {
		return err
	}
	return assignPositional(s, roles, gr.slots)
}

func assignTime(s *slots, gr rawGroup) error {
	roles, err := timeRoles(len(gr.slots))
	if err != nil {
		return err
	}
	return assignPositional(s, roles, gr.slots)
}

func dateRoles(n int) ([]role, error) {
	switch n {
	case 2:
		return []role{roleMonth, roleDay}, nil
	case 3:
		return []role{roleYear, roleMonth, roleDay}, nil
	default:
		return nil, ErrContext
	}
}

func timeRoles(n int) ([]role, error) {
	switch n {
	case 2:
		return []role{roleHour, roleMinute}, nil
	case 3:
		return []role{roleHour, roleMinute, roleSecond}, nil
	default:
		return nil, ErrContext
	}
}

func assignPositional(s *slots, roles []role, fields []rawField) error {
	for i, r := range roles {
		if err := claim(s, r, fields[i]); err != nil {
			return err
		}
	}
	return nil
}

// claim assigns a role, treating a present-but-empty field as a wildcard and
// rejecting a double assignment.
func claim(s *slots, r role, f rawField) error {
	if s[r] != nil {
		return ErrContext
	}
	slot := f
	s[r] = &slot
	return nil
}
