package unicat

import "unicode"

// From returns the General Unicode Category of the given rune.
func From(r rune) (category Category) {
	// TODO: order switch by frequency
	switch {
	case r >= 'a' && r <= 'z':
		category = UnicodeLl
	case r == ' ':
		category = UnicodeZs
	case r >= 'A' && r <= 'Z':
		category = UnicodeLu
	// L
	case unicode.Is(unicode.Ll, r):
		category = UnicodeLl
	case unicode.Is(unicode.Lu, r):
		category = UnicodeLu
	case unicode.Is(unicode.Lt, r):
		category = UnicodeLt
	case unicode.Is(unicode.Lm, r):
		category = UnicodeLm
	case unicode.Is(unicode.Lo, r):
		category = UnicodeLo
	// N
	case unicode.Is(unicode.Nd, r):
		category = UnicodeNd
	case unicode.Is(unicode.Nl, r):
		category = UnicodeNl
	case unicode.Is(unicode.No, r):
		category = UnicodeNo
	// P
	case unicode.Is(unicode.Pc, r):
		category = UnicodePc
	case unicode.Is(unicode.Pd, r):
		category = UnicodePd
	case unicode.Is(unicode.Ps, r):
		category = UnicodePs
	case unicode.Is(unicode.Pe, r):
		category = UnicodePe
	case unicode.Is(unicode.Pi, r):
		category = UnicodePi
	case unicode.Is(unicode.Pf, r):
		category = UnicodePf
	case unicode.Is(unicode.Po, r):
		category = UnicodePo
	// S
	case unicode.Is(unicode.Sm, r):
		category = UnicodeSm
	case unicode.Is(unicode.Sc, r):
		category = UnicodeSc
	case unicode.Is(unicode.Sk, r):
		category = UnicodeSk
	case unicode.Is(unicode.So, r):
		category = UnicodeSo
	// Z
	case unicode.Is(unicode.Zs, r):
		category = UnicodeZs
	case unicode.Is(unicode.Zl, r):
		category = UnicodeZl
	case unicode.Is(unicode.Zp, r):
		category = UnicodeZp
	// M
	case unicode.Is(unicode.Mn, r):
		category = UnicodeMn
	case unicode.Is(unicode.Mc, r):
		category = UnicodeMc
	case unicode.Is(unicode.Me, r):
		category = UnicodeMe
	// C
	case unicode.Is(unicode.Cc, r):
		category = UnicodeCc
	case unicode.Is(unicode.Cf, r):
		category = UnicodeCf
	case unicode.Is(unicode.Cs, r):
		category = UnicodeCs
	case unicode.Is(unicode.Co, r):
		category = UnicodeCo
	default:
		category = UnicodeCn
	}
	return
}

// IsLetter checks if the category is Letter (L).
func (c Category) IsLetter() bool {
	switch c {
	case UnicodeLl, UnicodeLm, UnicodeLo, UnicodeLt, UnicodeLu:
		return true
	default:
		return false
	}
}

// IsMark checks if the category is Mark (M).
func (c Category) IsMark() bool {
	switch c {
	case UnicodeMc, UnicodeMe, UnicodeMn:
		return true
	default:
		return false
	}
}

// IsNumber checks if the category is Number (N).
func (c Category) IsNumber() bool {
	switch c {
	case UnicodeNd, UnicodeNl, UnicodeNo:
		return true
	default:
		return false
	}
}

// IsPunctuation checks if the category is Punctuation (P).
func (c Category) IsPunctuation() bool {
	switch c {
	case UnicodePc, UnicodePd, UnicodePe, UnicodePf, UnicodePi, UnicodePo, UnicodePs:
		return true
	default:
		return false
	}
}

// IsSymbol checks if the category is Symbol (S).
func (c Category) IsSymbol() bool {
	switch c {
	case UnicodeSc, UnicodeSk, UnicodeSm, UnicodeSo:
		return true
	default:
		return false
	}
}

// IsSeparator checks if the category is Separator (Z).
func (c Category) IsSeparator() bool {
	switch c {
	case UnicodeZl, UnicodeZp, UnicodeZs:
		return true
	default:
		return false
	}
}

// IsOther checks if the category is Other (C).
func (c Category) IsOther() bool {
	switch c {
	case UnicodeCc, UnicodeCf, UnicodeCn, UnicodeCo, UnicodeCs:
		return true
	default:
		return false
	}
}
