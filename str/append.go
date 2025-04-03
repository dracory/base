package str

import "bytes"

// Append appends one or more strings.
func Append(values ...string) string {
	if len(values) == 0 {
		return ""
	}
	if len(values) == 1 {
		return values[0]
	}
	
	var buf bytes.Buffer
	for _, v := range values {
		buf.WriteString(v)
	}
	return buf.String()
}
