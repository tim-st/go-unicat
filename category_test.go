package unicat

import "testing"

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
