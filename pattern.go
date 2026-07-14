package isnow

import "time"

// Pattern is a parsed, canonicalized isnow. The zero value is not usable; obtain
// one from Parse. Patterns are immutable and safe to copy and share.
type Pattern struct {
	fields      [numRoles]fieldSpec
	bounds      []boundSpec
	canon       string
	explanation string
}

// Parse recognizes src, resolves symbols and the shorthand ladder, and validates
// field domains. It returns one of ErrSyntax, ErrSymbol, ErrRange, or ErrContext.
func Parse(src string) (Pattern, error) {
	raw, err := parseRaw(src)
	if err != nil {
		return Pattern{}, err
	}
	if err := validateUnits(raw); err != nil {
		return Pattern{}, err
	}
	sl, err := mapGroups(raw.groups)
	if err != nil {
		return Pattern{}, err
	}
	return assemble(sl, raw.bounds)
}

func assemble(sl slots, rawBounds []rawBound) (Pattern, error) {
	fields, err := compileAll(sl)
	if err != nil {
		return Pattern{}, err
	}
	bounds, err := compileBounds(rawBounds)
	if err != nil {
		return Pattern{}, err
	}
	return Pattern{
		fields:      fields,
		bounds:      bounds,
		canon:       renderCanonical(sl, bounds),
		explanation: renderExplain(sl, bounds),
	}, nil
}

func compileAll(sl slots) ([numRoles]fieldSpec, error) {
	var out [numRoles]fieldSpec
	for r := role(0); r < numRoles; r++ {
		spec, err := compileRole(r, sl[r])
		if err != nil {
			return out, err
		}
		out[r] = spec
	}
	return out, nil
}

func compileRole(r role, f *rawField) (fieldSpec, error) {
	if f == nil || !f.present {
		return wildcardField(), nil
	}
	return compileField(r, *f)
}

// Canonical returns the fully-qualified seven-field form of the isnow.
func (p Pattern) Canonical() string { return p.canon }

// String returns the canonical form (fmt.Stringer).
func (p Pattern) String() string { return p.canon }

// Is is the one-shot membership test: Parse then Holds.
func Is(src string, at time.Time) (bool, error) {
	p, err := Parse(src)
	if err != nil {
		return false, err
	}
	return p.Holds(at), nil
}
