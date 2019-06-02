// Package unicat implements identification of the General Unicode Category of a single rune.
package unicat

// Category is the General Unicode Category of a single rune.
// see also: https://en.wikipedia.org/w/index.php?title=Unicode_character_property&oldid=896020037#General_Category
type Category uint8

const (
	// UnicodeLu = Letter, uppercase
	UnicodeLu Category = iota
	// UnicodeLl = Letter, lowercase
	UnicodeLl
	// UnicodeLt = Letter, titlecase
	UnicodeLt
	// UnicodeLm = Letter, modifier
	UnicodeLm
	// UnicodeLo = Letter, other
	UnicodeLo
	// UnicodeMn = Mark, nonspacing
	UnicodeMn
	// UnicodeMc = Mark, spacing combining
	UnicodeMc
	// UnicodeMe = Mark, enclosing
	UnicodeMe
	// UnicodeNd = Number, decimal digit
	UnicodeNd
	// UnicodeNl = Number, letter
	UnicodeNl
	// UnicodeNo = Number, other
	UnicodeNo
	// UnicodePc = Punctuation, connector
	UnicodePc
	// UnicodePd = Punctuation, dash
	UnicodePd
	// UnicodePs = Punctuation, open
	UnicodePs
	// UnicodePe = Punctuation, close
	UnicodePe
	// UnicodePi = Punctuation, initial quote
	UnicodePi
	// UnicodePf = Punctuation, final quote
	UnicodePf
	// UnicodePo = Punctuation, other
	UnicodePo
	// UnicodeSm = Symbol, math
	UnicodeSm
	// UnicodeSc = Symbol, currency
	UnicodeSc
	// UnicodeSk = Symbol, modifier
	UnicodeSk
	// UnicodeSo = Symbol, other
	UnicodeSo
	// UnicodeZs = Separator, space
	UnicodeZs
	// UnicodeZl = Separator, line
	UnicodeZl
	// UnicodeZp = Separator, paragraph
	UnicodeZp
	// UnicodeCc = Other, control
	UnicodeCc
	// UnicodeCf = Other, format
	UnicodeCf
	// UnicodeCs = Other, surrogate
	UnicodeCs
	// UnicodeCo = Other, private use
	UnicodeCo
	// UnicodeCn = Other, not assigned
	UnicodeCn
)
