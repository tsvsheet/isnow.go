package isnow

// validateUnits rejects any quantity carrying a unit other than 'w' or 'd'
// (specs/contracts/semantics.md: an unknown unit is a symbol error).
func validateUnits(raw rawPattern) error {
	if err := groupsUnits(raw.groups); err != nil {
		return err
	}
	for _, b := range raw.bounds {
		if err := groupsUnits(b.groups); err != nil {
			return err
		}
	}
	return nil
}

func groupsUnits(groups []rawGroup) error {
	for _, gr := range groups {
		for _, f := range gr.slots {
			if err := fieldUnits(f); err != nil {
				return err
			}
		}
	}
	return nil
}

func fieldUnits(f rawField) error {
	for _, t := range f.terms {
		if err := termUnits(t); err != nil {
			return err
		}
	}
	return nil
}

func termUnits(t rawTerm) error {
	for _, a := range []*rawAtom{t.lo, t.hi} {
		if err := atomUnits(a); err != nil {
			return err
		}
	}
	return incrUnits(t.incr)
}

func atomUnits(a *rawAtom) error {
	if a == nil {
		return nil
	}
	return qtysUnits(a.qtys)
}

func incrUnits(in *rawIncr) error {
	if in == nil {
		return nil
	}
	return qtysUnits(in.qtys)
}

func qtysUnits(qs []rawQty) error {
	for _, q := range qs {
		if q.unit != "" && q.unit != "w" && q.unit != "d" {
			return ErrSymbol
		}
	}
	return nil
}
