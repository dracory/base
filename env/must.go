package env

// Must retrieves the value of an environment variable, panicking if not set.
func Must(key string) string {
	return ValueOrPanic(key)
}
