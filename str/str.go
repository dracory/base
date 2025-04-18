package str

// import (
// 	"bytes"
// 	"crypto/rand"
// 	"encoding/json"
// 	"path/filepath"
// 	"regexp"
// 	"strconv"
// 	"strings"
// 	"unicode"
// 	"unicode/utf8"

// 	"golang.org/x/exp/constraints"
// 	"golang.org/x/text/cases"
// 	"golang.org/x/text/language"
// )

// type String struct {
// 	value string
// }

// // ExcerptOption is the option for Excerpt method
// type ExcerptOption struct {
// 	Radius   int
// 	Omission string
// }

// // Of creates a new String instance with the given value.
// func Of(value string) *String {
// 	return &String{value: value}
// }

// // Basename returns the String instance with the basename of the current file path string,
// // and trims the suffix based on the parameter(optional).
// func (s *String) Basename(suffix ...string) *String {
// 	s.value = filepath.Base(s.value)
// 	if len(suffix) > 0 && suffix[0] != "" {
// 		s.value = strings.TrimSuffix(s.value, suffix[0])
// 	}
// 	return s
// }

// // Between returns the String instance with the substring between the given start and end strings.
// func (s *String) Between(start, end string) *String {
// 	if start == "" || end == "" {
// 		return s
// 	}
// 	return s.After(start).BeforeLast(end)
// }

// // BetweenFirst returns the String instance with the substring between the first occurrence of the given start string and the given end string.
// func (s *String) BetweenFirst(start, end string) *String {
// 	if start == "" || end == "" {
// 		return s
// 	}
// 	return s.Before(end).After(start)
// }

// // Camel returns the String instance in camel case.
// func (s *String) Camel() *String {
// 	return s.Studly().LcFirst()
// }

// // ChopEnd remove the given string(s) if it exists at the end of the haystack.
// func (s *String) ChopEnd(needle string, more ...string) *String {
// 	more = append([]string{needle}, more...)

// 	for _, v := range more {
// 		if s.EndsWith(v) {
// 			s.value = strings.TrimRight(s.value, v)
// 			break
// 		}
// 	}
// 	return s
// }

// // ChopStart remove the given string(s) if it exists at the start of the haystack.
// func (s *String) ChopStart(needle string, more ...string) *String {
// 	more = append([]string{needle}, more...)

// 	for _, v := range more {
// 		if s.StartsWith(v) {
// 			s.value = strings.TrimLeft(s.value, v)
// 			break
// 		}
// 	}
// 	return s
// }

// // Contains returns true if the string contains the given value or any of the values.
// func (s *String) Contains(values ...string) bool {
// 	for _, value := range values {
// 		if value != "" && strings.Contains(s.value, value) {
// 			return true
// 		}
// 	}

// 	return false
// }

// // ContainsAll returns true if the string contains all of the given values.
// func (s *String) ContainsAll(values ...string) bool {
// 	for _, value := range values {
// 		if !strings.Contains(s.value, value) {
// 			return false
// 		}
// 	}

// 	return true
// }

// // Dirname returns the String instance with the directory name of the current file path string.
// func (s *String) Dirname(levels ...int) *String {
// 	defaultLevels := 1
// 	if len(levels) > 0 {
// 		defaultLevels = levels[0]
// 	}

// 	dir := s.value
// 	for i := 0; i < defaultLevels; i++ {
// 		dir = filepath.Dir(dir)
// 	}

// 	s.value = dir
// 	return s
// }

// // EndsWith returns true if the string ends with the given value or any of the values.
// func (s *String) EndsWith(values ...string) bool {
// 	for _, value := range values {
// 		if value != "" && strings.HasSuffix(s.value, value) {
// 			return true
// 		}
// 	}

// 	return false
// }

// // Exactly returns true if the string is exactly the given value.
// func (s *String) Exactly(value string) bool {
// 	return s.value == value
// }

// // Excerpt returns the String instance truncated to the given length.
// func (s *String) Excerpt(phrase string, options ...ExcerptOption) *String {
// 	defaultOptions := ExcerptOption{
// 		Radius:   100,
// 		Omission: "...",
// 	}

// 	if len(options) > 0 {
// 		if options[0].Radius != 0 {
// 			defaultOptions.Radius = options[0].Radius
// 		}
// 		if options[0].Omission != "" {
// 			defaultOptions.Omission = options[0].Omission
// 		}
// 	}

// 	radius := maximum(0, defaultOptions.Radius)
// 	omission := defaultOptions.Omission

// 	regex := regexp.MustCompile(`(.*?)(` + regexp.QuoteMeta(phrase) + `)(.*)`)
// 	matches := regex.FindStringSubmatch(s.value)

// 	if len(matches) == 0 {
// 		return s
// 	}

// 	start := strings.TrimRight(matches[1], "")
// 	end := strings.TrimLeft(matches[3], "")

// 	end = Of(Substr(end, 0, radius)).LTrim("").
// 		Unless(func(s *String) bool {
// 			return s.Exactly(end)
// 		}, func(s *String) *String {
// 			return s.Append(omission)
// 		}).String()

// 	s.value = Of(Substr(start, maximum(len(start)-radius, 0), radius)).LTrim("").
// 		Unless(func(s *String) bool {
// 			return s.Exactly(start)
// 		}, func(s *String) *String {
// 			return s.Prepend(omission)
// 		}).Append(matches[2], end).String()

// 	return s
// }

// // Explode splits the string by given delimiter string.
// func (s *String) Explode(delimiter string, limit ...int) []string {
// 	defaultLimit := 1
// 	isLimitSet := false
// 	if len(limit) > 0 && limit[0] != 0 {
// 		defaultLimit = limit[0]
// 		isLimitSet = true
// 	}
// 	tempExplode := strings.Split(s.value, delimiter)
// 	if !isLimitSet || len(tempExplode) <= defaultLimit {
// 		return tempExplode
// 	}

// 	if defaultLimit > 0 {
// 		return append(tempExplode[:defaultLimit-1], strings.Join(tempExplode[defaultLimit-1:], delimiter))
// 	}

// 	if defaultLimit < 0 && len(tempExplode) <= -defaultLimit {
// 		return []string{}
// 	}

// 	return tempExplode[:len(tempExplode)+defaultLimit]
// }

// // Finish returns the String instance with the given value appended.
// // If the given value already ends with the suffix, it will not be added twice.
// func (s *String) Finish(value string) *String {
// 	quoted := regexp.QuoteMeta(value)
// 	reg := regexp.MustCompile("(?:" + quoted + ")+$")
// 	s.value = reg.ReplaceAllString(s.value, "") + value
// 	return s
// }

// // Headline returns the String instance in headline case.
// func (s *String) Headline() *String {
// 	parts := s.Explode(" ")

// 	if len(parts) > 1 {
// 		return s.Title()
// 	}

// 	parts = Of(strings.Join(parts, "_")).Studly().UcSplit()
// 	collapsed := Of(strings.Join(parts, "_")).
// 		Replace("-", "_").
// 		Replace(" ", "_").
// 		Replace("_", "_").Explode("_")

// 	s.value = strings.Join(collapsed, " ")
// 	return s
// }

// // Kebab returns the String instance in kebab case.
// func (s *String) Kebab() *String {
// 	return s.Snake("-")
// }

// // LcFirst returns the String instance with the first character lowercased.
// func (s *String) LcFirst() *String {
// 	if s.Length() == 0 {
// 		return s
// 	}
// 	s.value = strings.ToLower(Substr(s.value, 0, 1)) + Substr(s.value, 1)
// 	return s
// }

// // Length returns the length of the string.
// func (s *String) Length() int {
// 	return utf8.RuneCountInString(s.value)
// }

// // Limit returns the String instance truncated to the given length.
// func (s *String) Limit(limit int, end ...string) *String {
// 	defaultEnd := "..."
// 	if len(end) > 0 {
// 		defaultEnd = end[0]
// 	}

// 	if s.Length() <= limit {
// 		return s
// 	}
// 	s.value = Substr(s.value, 0, limit) + defaultEnd
// 	return s
// }

// // Lower returns the String instance in lower case.
// func (s *String) Lower() *String {
// 	s.value = strings.ToLower(s.value)
// 	return s
// }

// // Ltrim returns the String instance with the leftmost occurrence of the given value removed.
// func (s *String) LTrim(characters ...string) *String {
// 	if len(characters) == 0 {
// 		s.value = strings.TrimLeft(s.value, " ")
// 		return s
// 	}

// 	s.value = strings.TrimLeft(s.value, characters[0])
// 	return s
// }

// // Mask returns the String instance with the given character masking the specified number of characters.
// func (s *String) Mask(character string, index int, length ...int) *String {
// 	// Check if the character is empty, if so, return the original string.
// 	if character == "" {
// 		return s
// 	}

// 	segment := Substr(s.value, index, length...)

// 	// Check if the segment is empty, if so, return the original string.
// 	if segment == "" {
// 		return s
// 	}

// 	strLen := utf8.RuneCountInString(s.value)
// 	startIndex := index

// 	// Check if the start index is out of bounds.
// 	if index < 0 {
// 		if index < -strLen {
// 			startIndex = 0
// 		} else {
// 			startIndex = strLen + index
// 		}
// 	}

// 	start := Substr(s.value, 0, startIndex)
// 	segmentLen := utf8.RuneCountInString(segment)
// 	end := Substr(s.value, startIndex+segmentLen)

// 	s.value = start + strings.Repeat(Substr(character, 0, 1), segmentLen) + end
// 	return s
// }

// // Match returns the String instance with the first occurrence of the given pattern.
// func (s *String) Match(pattern string) *String {
// 	if pattern == "" {
// 		return s
// 	}
// 	reg := regexp.MustCompile(pattern)
// 	s.value = reg.FindString(s.value)
// 	return s
// }

// // MatchAll returns all matches for the given regular expression.
// func (s *String) MatchAll(pattern string) []string {
// 	if pattern == "" {
// 		return []string{s.value}
// 	}
// 	reg := regexp.MustCompile(pattern)
// 	return reg.FindAllString(s.value, -1)
// }

// // NewLine appends one or more new lines to the current string.
// func NewLine(str string, count ...int) string {
// 	if len(count) == 0 {
// 		str += "\n"
// 		return str
// 	}

// 	str += strings.Repeat("\n", count[0])
// 	return str
// }

// // PadBoth returns the String instance padded to the left and right sides of the given length.
// func (s *String) PadBoth(length int, pad ...string) *String {
// 	defaultPad := " "
// 	if len(pad) > 0 {
// 		defaultPad = pad[0]
// 	}
// 	short := maximum(0, length-s.Length())
// 	left := short / 2
// 	right := short/2 + short%2

// 	s.value = Substr(strings.Repeat(defaultPad, left), 0, left) + s.value + Substr(strings.Repeat(defaultPad, right), 0, right)

// 	return s
// }

// // PadLeft returns the String instance padded to the left side of the given length.
// func (s *String) PadLeft(length int, pad ...string) *String {
// 	defaultPad := " "
// 	if len(pad) > 0 {
// 		defaultPad = pad[0]
// 	}
// 	short := maximum(0, length-s.Length())

// 	s.value = Substr(strings.Repeat(defaultPad, short), 0, short) + s.value
// 	return s
// }

// // PadRight returns the String instance padded to the right side of the given length.
// func (s *String) PadRight(length int, pad ...string) *String {
// 	defaultPad := " "
// 	if len(pad) > 0 {
// 		defaultPad = pad[0]
// 	}
// 	short := maximum(0, length-s.Length())

// 	s.value = s.value + Substr(strings.Repeat(defaultPad, short), 0, short)
// 	return s
// }

// // Pipe passes the string to the given callback and returns the result.
// func (s *String) Pipe(callback func(s string) string) *String {
// 	s.value = callback(s.value)
// 	return s
// }

// // Prepend one or more strings to the current string.
// func (s *String) Prepend(values ...string) *String {
// 	s.value = strings.Join(values, "") + s.value
// 	return s
// }

// // Remove returns the String instance with the first occurrence of the given value removed.
// func (s *String) Remove(values ...string) *String {
// 	for _, value := range values {
// 		s.value = strings.ReplaceAll(s.value, value, "")
// 	}

// 	return s
// }

// // Repeat returns the String instance repeated the given number of times.
// func (s *String) Repeat(times int) *String {
// 	s.value = strings.Repeat(s.value, times)
// 	return s
// }

// // Replace returns the String instance with all occurrences of the search string replaced by the given replacement string.
// func (s *String) Replace(search string, replace string, caseSensitive ...bool) *String {
// 	caseSensitive = append(caseSensitive, true)
// 	if len(caseSensitive) > 0 && !caseSensitive[0] {
// 		s.value = regexp.MustCompile("(?i)"+search).ReplaceAllString(s.value, replace)
// 		return s
// 	}
// 	s.value = strings.ReplaceAll(s.value, search, replace)
// 	return s
// }

// // ReplaceEnd returns the String instance with the last occurrence of the given value replaced.
// func (s *String) ReplaceEnd(search string, replace string) *String {
// 	if search == "" {
// 		return s
// 	}

// 	if s.EndsWith(search) {
// 		return s.ReplaceLast(search, replace)
// 	}

// 	return s
// }

// // ReplaceFirst returns the String instance with the first occurrence of the given value replaced.
// func (s *String) ReplaceFirst(search string, replace string) *String {
// 	if search == "" {
// 		return s
// 	}
// 	s.value = strings.Replace(s.value, search, replace, 1)
// 	return s
// }

// // ReplaceLast returns the String instance with the last occurrence of the given value replaced.
// func (s *String) ReplaceLast(search string, replace string) *String {
// 	if search == "" {
// 		return s
// 	}
// 	index := strings.LastIndex(s.value, search)
// 	if index != -1 {
// 		s.value = s.value[:index] + replace + s.value[index+len(search):]
// 		return s
// 	}

// 	return s
// }

// // ReplaceMatches returns the String instance with all occurrences of the given pattern
// // replaced by the given replacement string.
// func (s *String) ReplaceMatches(pattern string, replace string) *String {
// 	s.value = regexp.MustCompile(pattern).ReplaceAllString(s.value, replace)
// 	return s
// }

// // ReplaceStart returns the String instance with the first occurrence of the given value replaced.
// func (s *String) ReplaceStart(search string, replace string) *String {
// 	if search == "" {
// 		return s
// 	}

// 	if s.StartsWith(search) {
// 		return s.ReplaceFirst(search, replace)
// 	}

// 	return s
// }

// // RTrim returns the String instance with the right occurrences of the given value removed.
// func (s *String) RTrim(characters ...string) *String {
// 	if len(characters) == 0 {
// 		s.value = strings.TrimRight(s.value, " ")
// 		return s
// 	}

// 	s.value = strings.TrimRight(s.value, characters[0])
// 	return s
// }

// // Snake returns the String instance in snake case.
// func (s *String) Snake(delimiter ...string) *String {
// 	defaultDelimiter := "_"
// 	if len(delimiter) > 0 {
// 		defaultDelimiter = delimiter[0]
// 	}
// 	words := fieldsFunc(s.value, func(r rune) bool {
// 		return r == ' ' || r == ',' || r == '.' || r == '-' || r == '_'
// 	}, func(r rune) bool {
// 		return unicode.IsUpper(r)
// 	})

// 	casesLower := cases.Lower(language.Und)
// 	var studlyWords []string
// 	for _, word := range words {
// 		studlyWords = append(studlyWords, casesLower.String(word))
// 	}

// 	s.value = strings.Join(studlyWords, defaultDelimiter)
// 	return s
// }

// // Split splits the string by given pattern string.
// func (s *String) Split(pattern string, limit ...int) []string {
// 	r := regexp.MustCompile(pattern)
// 	defaultLimit := -1
// 	if len(limit) != 0 {
// 		defaultLimit = limit[0]
// 	}

// 	return r.Split(s.value, defaultLimit)
// }

// // Squish returns the String instance with consecutive whitespace characters collapsed into a single space.
// func (s *String) Squish() *String {
// 	leadWhitespace := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
// 	insideWhitespace := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
// 	s.value = leadWhitespace.ReplaceAllString(s.value, "")
// 	s.value = insideWhitespace.ReplaceAllString(s.value, " ")
// 	return s
// }

// // Start returns the String instance with the given value prepended.
// func (s *String) Start(prefix string) *String {
// 	quoted := regexp.QuoteMeta(prefix)
// 	re := regexp.MustCompile(`^(` + quoted + `)+`)
// 	s.value = prefix + re.ReplaceAllString(s.value, "")
// 	return s
// }

// // StartsWith returns true if the string starts with the given value or any of the values.
// func (s *String) StartsWith(values ...string) bool {
// 	for _, value := range values {
// 		if strings.HasPrefix(s.value, value) {
// 			return true
// 		}
// 	}

// 	return false
// }

// // Studly returns the String instance in studly case.
// func (s *String) Studly() *String {
// 	words := fieldsFunc(s.value, func(r rune) bool {
// 		return r == '_' || r == ' ' || r == '-' || r == ',' || r == '.'
// 	}, func(r rune) bool {
// 		return unicode.IsUpper(r)
// 	})

// 	casesTitle := cases.Title(language.Und)
// 	var studlyWords []string
// 	for _, word := range words {
// 		studlyWords = append(studlyWords, casesTitle.String(word))
// 	}

// 	s.value = strings.Join(studlyWords, "")
// 	return s
// }

// // Swap replaces all occurrences of the search string with the given replacement string.
// func (s *String) Swap(replacements map[string]string) *String {
// 	if len(replacements) == 0 {
// 		return s
// 	}

// 	oldNewPairs := make([]string, 0, len(replacements)*2)
// 	for k, v := range replacements {
// 		if k == "" {
// 			return s
// 		}
// 		oldNewPairs = append(oldNewPairs, k, v)
// 	}

// 	s.value = strings.NewReplacer(oldNewPairs...).Replace(s.value)
// 	return s
// }

// // Tap passes the string to the given callback and returns the string.
// func (s *String) Tap(callback func(String)) *String {
// 	callback(*s)
// 	return s
// }

// // Test returns true if the string matches the given pattern.
// func (s *String) Test(pattern string) bool {
// 	return s.IsMatch(pattern)
// }

// // Title returns the String instance in title case.
// func (s *String) Title() *String {
// 	casesTitle := cases.Title(language.Und)
// 	s.value = casesTitle.String(s.value)
// 	return s
// }

// // Trim returns the String instance with trimmed characters from the left and right sides.
// func (s *String) Trim(characters ...string) *String {
// 	if len(characters) == 0 {
// 		s.value = strings.TrimSpace(s.value)
// 		return s
// 	}

// 	s.value = strings.Trim(s.value, characters[0])
// 	return s
// }

// // Unless returns the String instance with the given fallback applied if the given condition is false.
// func (s *String) Unless(callback func(*String) bool, fallback func(*String) *String) *String {
// 	if !callback(s) {
// 		return fallback(s)
// 	}

// 	return s
// }

// // When returns the String instance with the given callback applied if the given condition is true.
// // If the condition is false, the fallback callback is applied (if provided).
// func (s *String) When(condition bool, callback ...func(*String) *String) *String {
// 	if condition {
// 		if len(callback) > 0 && callback[0] != nil {
// 			return callback[0](s)
// 		}
// 	} else {
// 		if len(callback) > 1 && callback[1] != nil {
// 			return callback[1](s)
// 		}
// 	}

// 	return s
// }

// // WhenContains returns the String instance with the given callback applied if the string contains the given value.
// func (s *String) WhenContains(value string, callback ...func(*String) *String) *String {
// 	return s.When(s.Contains(value), callback...)
// }

// // WhenContainsAll returns the String instance with the given callback applied if the string contains all the given values.
// func (s *String) WhenContainsAll(values []string, callback ...func(*String) *String) *String {
// 	return s.When(s.ContainsAll(values...), callback...)
// }

// // WhenEmpty returns the String instance with the given callback applied if the string is empty.
// func (s *String) WhenEmpty(callback ...func(*String) *String) *String {
// 	return s.When(s.IsEmpty(), callback...)
// }

// // WhenIsAscii returns the String instance with the given callback applied if the string contains only ASCII characters.
// func (s *String) WhenIsAscii(callback ...func(*String) *String) *String {
// 	return s.When(s.IsAscii(), callback...)
// }

// // WhenNotEmpty returns the String instance with the given callback applied if the string is not empty.
// func (s *String) WhenNotEmpty(callback ...func(*String) *String) *String {
// 	return s.When(s.IsNotEmpty(), callback...)
// }

// // WhenStartsWith returns the String instance with the given callback applied if the string starts with the given value.
// func (s *String) WhenStartsWith(value []string, callback ...func(*String) *String) *String {
// 	return s.When(s.StartsWith(value...), callback...)
// }

// // WhenEndsWith returns the String instance with the given callback applied if the string ends with the given value.
// func (s *String) WhenEndsWith(value []string, callback ...func(*String) *String) *String {
// 	return s.When(s.EndsWith(value...), callback...)
// }

// // WhenExactly returns the String instance with the given callback applied if the string is exactly the given value.
// func (s *String) WhenExactly(value string, callback ...func(*String) *String) *String {
// 	return s.When(s.Exactly(value), callback...)
// }

// // WhenNotExactly returns the String instance with the given callback applied if the string is not exactly the given value.
// func (s *String) WhenNotExactly(value string, callback ...func(*String) *String) *String {
// 	return s.When(!s.Exactly(value), callback...)
// }

// // WhenIs returns the String instance with the given callback applied if the string matches any of the given patterns.
// func (s *String) WhenIs(value string, callback ...func(*String) *String) *String {
// 	return s.When(s.Is(value), callback...)
// }

// // WhenIsUlid returns the String instance with the given callback applied if the string is a valid ULID.
// func (s *String) WhenIsUlid(callback ...func(*String) *String) *String {
// 	return s.When(s.IsUlid(), callback...)
// }

// // WhenIsUuid returns the String instance with the given callback applied if the string is a valid UUID.
// func (s *String) WhenIsUuid(callback ...func(*String) *String) *String {
// 	return s.When(s.IsUuid(), callback...)
// }

// // WhenTest returns the String instance with the given callback applied if the string matches the given pattern.
// func (s *String) WhenTest(pattern string, callback ...func(*String) *String) *String {
// 	return s.When(s.Test(pattern), callback...)
// }

// func Random(length int) string {
// 	b := make([]byte, length)
// 	_, err := rand.Read(b)
// 	if err != nil {
// 		panic(err)
// 	}
// 	letters := "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	for i, v := range b {
// 		b[i] = letters[v%byte(len(letters))]
// 	}

// 	return string(b)
// }

// type Buffer struct {
// 	*bytes.Buffer
// }

// func NewBuffer() *Buffer {
// 	return &Buffer{Buffer: new(bytes.Buffer)}
// }

// func (b *Buffer) Append(i any) *Buffer {
// 	switch val := i.(type) {
// 	case int:
// 		b.append(strconv.Itoa(val))
// 	case int64:
// 		b.append(strconv.FormatInt(val, 10))
// 	case uint:
// 		b.append(strconv.FormatUint(uint64(val), 10))
// 	case uint64:
// 		b.append(strconv.FormatUint(val, 10))
// 	case string:
// 		b.append(val)
// 	case []byte:
// 		b.Write(val)
// 	case rune:
// 		b.WriteRune(val)
// 	}
// 	return b
// }

// func (b *Buffer) append(s string) *Buffer {
// 	b.WriteString(s)

// 	return b
// }

// // fieldsFunc splits the input string into words with preservation, following the rules defined by
// // the provided functions f and preserveFunc.
// func fieldsFunc(s string, f func(rune) bool, preserveFunc ...func(rune) bool) []string {
// 	var fields []string
// 	var currentField strings.Builder

// 	shouldPreserve := func(r rune) bool {
// 		for _, preserveFn := range preserveFunc {
// 			if preserveFn(r) {
// 				return true
// 			}
// 		}
// 		return false
// 	}

// 	for _, r := range s {
// 		if f(r) {
// 			if currentField.Len() > 0 {
// 				fields = append(fields, currentField.String())
// 				currentField.Reset()
// 			}
// 		} else if shouldPreserve(r) {
// 			if currentField.Len() > 0 {
// 				fields = append(fields, currentField.String())
// 				currentField.Reset()
// 			}
// 			currentField.WriteRune(r)
// 		} else {
// 			currentField.WriteRune(r)
// 		}
// 	}

// 	if currentField.Len() > 0 {
// 		fields = append(fields, currentField.String())
// 	}

// 	return fields
// }

// // maximum returns the largest of x or y.
// func maximum[T constraints.Ordered](x T, y T) T {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }
