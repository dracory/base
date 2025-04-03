package str

import "strings"

// WordCount returns the number of words in the string.
func WordCount(str string) int {
	return len(strings.Fields(str))
}
