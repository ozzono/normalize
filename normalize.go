package normalize

import (
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func Norm(str string) string {
	normStr, _, _ := transform.String(transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC), str)
	return normStr
}
