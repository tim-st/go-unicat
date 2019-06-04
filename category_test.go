package unicat

import (
	"testing"
	"unicode"
)

func TestCategory(t *testing.T) {
	c := From('Ä')
	if c != UnicodeLu {
		t.Errorf("Expected Category Lu; got %s.\n", c)
	}
	if c.String() != "Lu" {
		t.Errorf("Expected String `Lu`; got %s", c)
	}

	badRune := From(rune(-1))
	if badRune.String() != "Cn" {
		t.Errorf("Expected Category `Cn`; got %s", c)
	}

	badRune2 := From(rune(0x10FFFF + 1))
	if badRune2.String() != "Cn" {
		t.Errorf("Expected Category `Cn`; got %s", c)
	}

}

func TestAllRunes(t *testing.T) {
	for i := -1; i <= 0x10FFFF; i++ {
		r := rune(i)

		if r, expected, got := rune(i), slowestButSafeFrom(r), From(r); expected != got {
			t.Errorf("Rune %d (%s): expected Category %s; got %s.", i, string(r), expected, got)
		}

	}
}

func slowestButSafeFrom(r rune) (category Category) {
	switch {
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
