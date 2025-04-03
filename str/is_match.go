package str

import "regexp"

// IsMatch returns true if the string matches any of the given patterns.
func IsMatch(str string, patterns ...string) bool {
	if len(patterns) == 0 {
		return false
	}

	for _, pattern := range patterns {
		if pattern == "" {
			continue
		}

		reg := regexp.MustCompile(pattern)
		if reg.MatchString(str) {
			return true
		}
	}

	return false
}
