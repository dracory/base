package config

import "strings"

// LoadAccumulator centralizes validation error collection while building a
// configuration instance. Helper methods mirror the existing RequireString
// and RequireWhen primitives so callers stay concise.
type LoadAccumulator struct {
	errs []error
}

// Add appends err to the accumulator when it is non-nil.
func (a *LoadAccumulator) Add(err error) {
	if err != nil {
		a.errs = append(a.errs, err)
	}
}

// MustString returns the value for key via RequireString, while recording any
// resulting error for later inspection.
func (a *LoadAccumulator) MustString(key, context string) string {
	value, err := RequireString(key, context)
	a.Add(err)
	return value
}

// MustWhen delegates to RequireWhen and records any error produced under the
// supplied condition.
func (a *LoadAccumulator) MustWhen(condition bool, key, context, value string) {
	if err := RequireWhen(condition, key, context, value); err != nil {
		a.Add(err)
	}
}

// Err returns a ValidationError wrapping all collected issues. Nil is returned
// when no errors were recorded.
func (a *LoadAccumulator) Err() error {
	if len(a.errs) == 0 {
		return nil
	}
	return ValidationError{errs: a.errs}
}

// ValidationError aggregates multiple missing/invalid environment errors while
// preserving the existing error semantics.
type ValidationError struct {
	errs []error
}

func (e ValidationError) Error() string {
	if len(e.errs) == 0 {
		return "config: validation failed"
	}

	var builder strings.Builder
	builder.WriteString("config: validation failed:\n")
	for i, err := range e.errs {
		if i > 0 {
			builder.WriteByte('\n')
		}
		builder.WriteString(" - ")
		builder.WriteString(err.Error())
	}
	return builder.String()
}

// Errors exposes the accumulated error slice. A defensive copy is returned to
// avoid callers mutating internal state.
func (e ValidationError) Errors() []error {
	return append([]error(nil), e.errs...)
}
