package str

import "strings"

// fieldsFunc splits the input string into words with preservation, following the rules defined by
// the provided functions f and preserveFunc.
func fieldsFunc(s string, f func(rune) bool, preserveFunc ...func(rune) bool) []string {
	var fields []string
	var currentField strings.Builder

	shouldPreserve := func(r rune) bool {
		for _, preserveFn := range preserveFunc {
			if preserveFn(r) {
				return true
			}
		}
		return false
	}

	for _, r := range s {
		if f(r) {
			if currentField.Len() > 0 {
				fields = append(fields, currentField.String())
				currentField.Reset()
			}
		} else if shouldPreserve(r) {
			if currentField.Len() > 0 {
				fields = append(fields, currentField.String())
				currentField.Reset()
			}
			currentField.WriteRune(r)
		} else {
			currentField.WriteRune(r)
		}
	}

	if currentField.Len() > 0 {
		fields = append(fields, currentField.String())
	}

	return fields
}
