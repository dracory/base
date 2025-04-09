# String Package

The `str` package provides a comprehensive set of string manipulation utilities for the Dracory framework.
It offers a wide range of functions for string operations, validation, transformation, and formatting.

## Overview

This package includes utilities for:

1. **String Manipulation**: Functions for modifying, extracting, and transforming strings
2. **String Validation**: Functions for checking string properties and patterns
3. **String Formatting**: Functions for formatting strings in various ways
4. **String Conversion**: Functions for converting between different string formats and encodings
5. **String Generation**: Functions for generating random strings and hashes

## Key Features

- **Pattern Matching**: Check if strings match patterns using glob syntax
- **String Extraction**: Extract substrings based on various criteria
- **String Transformation**: Convert strings to different formats (camelCase, snake_case, etc.)
- **String Validation**: Check if strings are empty, UUIDs, ULIDs, etc.
- **String Encoding**: Encode and decode strings using various encodings (Base32, Base64, etc.)
- **String Hashing**: Generate hashes from strings (MD5, BCrypt, etc.)
- **String Formatting**: Format strings with padding, truncation, etc.
- **String Generation**: Generate random strings and identifiers

## Usage Examples

### String Validation

```go
import "github.com/dracory/base/str"

// Check if a string is empty
isEmpty := str.IsEmpty("")  // true

// Check if a string is not empty
isNotEmpty := str.IsNotEmpty("Hello")  // true

// Check if a string is a UUID
isUUID := str.IsUUID("123e4567-e89b-12d3-a456-426614174000")  // true

// Check if a string is a ULID
isULID := str.IsULID("01H9Z8K2P3M4N5Q6R7S8T9U0V")  // true

// Check if a string matches a pattern
matches := str.Is("hello.txt", "*.txt")  // true
```

### String Manipulation

```go
import "github.com/dracory/base/str"

// Extract substring between two strings
between := str.Between("Hello [World] Test", "[", "]")  // "World"

// Extract substring before a string
before := str.Before("Hello World", "World")  // "Hello "

// Extract substring after a string
after := str.After("Hello World", "Hello ")  // "World"

// Extract substring before the last occurrence of a string
beforeLast := str.BeforeLast("Hello World World", "World")  // "Hello "

// Extract substring after the last occurrence of a string
afterLast := str.AfterLast("Hello World World", "World")  // ""

// Extract substring from the left
leftFrom := str.LeftFrom("Hello World", 5)  // "Hello"

// Extract substring from the right
rightFrom := str.RightFrom("Hello World", 5)  // "World"

// Truncate a string
truncated := str.Truncate("Hello World", 8, "...")  // "Hello..."

// Convert to snake_case
snake := str.ToSnake("HelloWorld")  // "hello_world"

// Convert to camelCase
camel := str.ToCamel("hello_world")  // "helloWorld"

// Convert first character to uppercase
ucFirst := str.UcFirst("hello")  // "Hello"

// Convert string to uppercase
upper := str.Upper("hello")  // "HELLO"

// Split string into words
words := str.Words("Hello World")  // ["Hello", "World"]

// Count words in a string
wordCount := str.WordCount("Hello World")  // 2

// Create a URL-friendly slug
slug := str.Slugify("Hello World!", '-')  // "hello-world"
```

### String Encoding and Hashing

```go
import "github.com/dracory/base/str"

// Encode string to Base64
base64 := str.Base64Encode("Hello World")  // "SGVsbG8gV29ybGQ="

// Decode Base64 string
decoded := str.Base64Decode("SGVsbG8gV29ybGQ=")  // "Hello World"

// Encode string to Base32 Extended
base32 := str.Base32ExtendedEncode("Hello World")  // "91IMOR3FCPBI41"

// Decode Base32 Extended string
decoded := str.Base32ExtendedDecode("91IMOR3FCPBI41")  // "Hello World"

// Generate MD5 hash
md5 := str.MD5("Hello World")  // "b10a8db164e0754105b7a99be72e3fe5"

// Generate BCrypt hash
bcrypt := str.ToBcryptHash("password")  // "$2a$10$..."

// Compare password with BCrypt hash
matches := str.BcryptHashCompare("password", bcrypt)  // true
```

### String Generation

```go
import "github.com/dracory/base/str"

// Generate a random string
random := str.Random(10)  // "a1b2c3d4e5"

// Generate a random string from a gamma distribution
randomGamma := str.RandomFromGamma(10, 2.0, 1.0)  // "a1b2c3d4e5"

// Convert integer to Base32
base32 := str.IntToBase32(12345)  // "3RP"

// Convert integer to Base36
base36 := str.IntToBase36(12345)  // "9IX"
```

## Best Practices

1. **Use Appropriate Functions**: Choose the most specific function for your task
2. **Handle Errors**: Check for errors when using functions that can fail
3. **Consider Performance**: Some functions may be more efficient than others for your use case
4. **Validate Input**: Always validate input strings before processing them
5. **Use Constants**: Use constants for repeated string values

## License

This package is part of the dracory/base project and is licensed under the same terms. 