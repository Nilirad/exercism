package acronym

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func Abbreviate(s string) string {
	wordBoundary := func(r rune) bool {
		return unicode.IsSpace(r) || r == '-' || r == '_'
	}
	words := strings.FieldsFunc(s, wordBoundary)

	var builder strings.Builder
	for _, word := range words {
		firstLetter, size := utf8.DecodeRuneInString(word)
		if size > 0 {
			builder.WriteRune(unicode.ToUpper(firstLetter))
		}
	}

	return builder.String()
}
