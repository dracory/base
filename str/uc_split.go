package str

import (
	"unicode"
)

// UcSplit splits the string into words using uppercase characters as the delimiter.
func UcSplit(s string) []string {
	words := fieldsFunc(s, func(r rune) bool {
		return false
	}, func(r rune) bool {
		return unicode.IsUpper(r)
	})
	return words
}
