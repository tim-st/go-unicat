package unicat

func (c Category) String() (s string) {
	switch c {
	case UnicodeLu:
		s = "Lu"
	case UnicodeLl:
		s = "Ll"
	case UnicodeLt:
		s = "Lt"
	case UnicodeLm:
		s = "Lm"
	case UnicodeLo:
		s = "Lo"
	case UnicodeMn:
		s = "Mn"
	case UnicodeMc:
		s = "Mc"
	case UnicodeMe:
		s = "Me"
	case UnicodeNd:
		s = "Nd"
	case UnicodeNl:
		s = "Nl"
	case UnicodeNo:
		s = "No"
	case UnicodePc:
		s = "Pc"
	case UnicodePd:
		s = "Pd"
	case UnicodePs:
		s = "Ps"
	case UnicodePe:
		s = "Pe"
	case UnicodePi:
		s = "Pi"
	case UnicodePf:
		s = "Pf"
	case UnicodePo:
		s = "Po"
	case UnicodeSm:
		s = "Sm"
	case UnicodeSc:
		s = "Sc"
	case UnicodeSk:
		s = "Sk"
	case UnicodeSo:
		s = "So"
	case UnicodeZs:
		s = "Zs"
	case UnicodeZl:
		s = "Zl"
	case UnicodeZp:
		s = "Zp"
	case UnicodeCc:
		s = "Cc"
	case UnicodeCf:
		s = "Cf"
	case UnicodeCs:
		s = "Cs"
	case UnicodeCo:
		s = "Co"
	case UnicodeCn:
		s = "Cn"
	}
	return
}
