package str

import (
	"regexp"
	"strings"
)

// Is returns true if the string matches any of the given patterns.
func Is(str string, patterns ...string) bool {
	if len(patterns) == 0 {
		return false
	}
	for _, pattern := range patterns {
		if pattern == "" {
			if str == "" {
				return false
			}
			continue
		}

		// Direct match optimization
		if pattern == str {
			return true
		}

		// Convert glob pattern to regex
		var regexPattern strings.Builder
		regexPattern.WriteString("^")

		escaped := false
		for _, r := range pattern {
			if escaped {
				// Previous char was '\', escape this char literally
				regexPattern.WriteString(regexp.QuoteMeta(string(r)))
				escaped = false
			} else if r == '\\' {
				escaped = true
				// Don't append anything yet, wait for the next char
			} else if r == '*' {
				regexPattern.WriteString(".*")
			} else {
				// Escape other regex metacharacters
				regexPattern.WriteString(regexp.QuoteMeta(string(r)))
			}
		}
		// Handle trailing escape character if present (treat as literal backslash)
		if escaped {
			regexPattern.WriteString(regexp.QuoteMeta("\\"))
		}

		regexPattern.WriteString("$")

		// Compile and match
		// Use regexp.Compile, not MustCompile, to handle potential pattern errors gracefully
		regex, err := regexp.Compile(regexPattern.String())
		if err != nil {
			// Skip invalid patterns - consider logging or error handling if needed
			continue
		}

		if regex.MatchString(str) {
			return true
		}
	}

	return false
}
