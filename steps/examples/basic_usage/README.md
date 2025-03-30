# Basic Usage Example

This example demonstrates the basic usage of the steps package by creating and running multiple steps in sequence.

## Overview

The example shows:
1. How to create a step using the `Step` function
2. How to create a steps collection using `Steps()`
3. How to add multiple steps to the collection
4. How to run the steps with a context

## Key Concepts

- Step Creation: Using the `Step` function to create a step that increments a value
- Steps Collection: Using `Steps()` to create a collection of steps
- Context: Using a context to share data between steps
- Execution: Running multiple steps in sequence

## Running the Example

To run this example:

```bash
# Run the main program
go run main.go

# Run the tests
go test -v
```

## Expected Output

The program will output:
```
Final value: 5
