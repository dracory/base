# Array Package

The `arr` package provides a comprehensive set of utilities for working with arrays, slices, and maps in Go. It offers a collection of functions to simplify common array operations and manipulations.

## Features

- **Array Manipulation**
  - Filter arrays based on conditions
  - Map arrays to new values
  - Reverse array order
  - Shuffle array elements
  - Split arrays into chunks
  - Merge multiple arrays
  - Remove elements by index
  - Move elements up/down by index

- **Array Analysis**
  - Find minimum and maximum values
  - Calculate sum of numeric arrays
  - Count elements
  - Count elements by condition
  - Check if array contains specific values
  - Check if array contains all values
  - Find index of elements
  - Get unique values

- **Map Operations**
  - Extract keys from maps
  - Extract values from maps
  - Group array elements by key
  - Check map equality

- **Iteration**
  - Iterate over array elements
  - Filter empty values

## Installation

```bash
go get github.com/dracory/base/arr
```

## Usage Examples

### Filtering Arrays

```go
package main

import (
    "fmt"
    "github.com/dracory/base/arr"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    // Filter even numbers
    evenNumbers := arr.Filter(numbers, func(n int) bool {
        return n%2 == 0
    })
    
    fmt.Println(evenNumbers) // [2 4 6 8 10]
}
```

### Mapping Arrays

```go
package main

import (
    "fmt"
    "github.com/dracory/base/arr"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    
    // Double each number
    doubled := arr.Map(numbers, func(n int) int {
        return n * 2
    })
    
    fmt.Println(doubled) // [2 4 6 8 10]
}
```

### Array Manipulation

```go
package main

import (
    "fmt"
    "github.com/dracory/base/arr"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    
    // Reverse the array
    reversed := arr.Reverse(numbers)
    fmt.Println(reversed) // [5 4 3 2 1]
    
    // Shuffle the array
    shuffled := arr.Shuffle(numbers)
    fmt.Println(shuffled) // Random order
    
    // Split into chunks
    chunks := arr.Split(numbers, 2)
    fmt.Println(chunks) // [[1 2] [3 4] [5]]
}
```

### Map Operations

```go
package main

import (
    "fmt"
    "github.com/dracory/base/arr"
)

func main() {
    userMap := map[string]string{
        "name": "John",
        "age": "30",
        "email": "john@example.com",
    }
    
    // Get all keys
    keys := arr.Keys(userMap)
    fmt.Println(keys) // ["name" "age" "email"]
    
    // Get all values
    values := arr.Values(userMap)
    fmt.Println(values) // ["John" "30" "john@example.com"]
}
```

### Array Analysis

```go
package main

import (
    "fmt"
    "github.com/dracory/base/arr"
)

func main() {
    numbers := []int{5, 2, 9, 1, 7, 3, 8, 4, 6}
    
    // Find minimum
    min := arr.Min(numbers)
    fmt.Println(min) // 1
    
    // Find maximum
    max := arr.Max(numbers)
    fmt.Println(max) // 9
    
    // Calculate sum
    sum := arr.Sum(numbers)
    fmt.Println(sum) // 45
    
    // Count elements
    count := arr.Count(numbers)
    fmt.Println(count) // 9
    
    // Get unique values
    duplicates := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
    unique := arr.Unique(duplicates)
    fmt.Println(unique) // [1 2 3 4]
}
```

## Error Handling

Most functions in the `arr` package handle errors gracefully:

- Empty arrays return appropriate zero values
- Index out of bounds errors are prevented
- Type mismatches are handled safely

## Performance Considerations

- The package is designed for convenience rather than maximum performance
- For large arrays, consider using standard Go operations for better performance
- Some operations like `Shuffle` and `Reverse` modify arrays in-place when possible

## License

This package is part of the dracory/base project and is licensed under the same terms. See the main [README.md](../README.md) for license information. 