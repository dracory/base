# Steps Examples

This directory contains example code demonstrating various features of the steps package. Each example is in its own folder with a main program and tests.

## Examples

### Basic Usage
- [examples/basic_usage/main.go](basic_usage/main.go): Shows how to create and run a simple step
- [examples/basic_usage/basic_usage_test.go](basic_usage/basic_usage_test.go): Tests for the basic usage example

### Dependencies
- [examples/dependencies/main.go](dependencies/main.go): Demonstrates how to create steps with dependencies
- [examples/dependencies/dependencies_test.go](dependencies/dependencies_test.go): Tests for the dependencies example

### Error Handling
- [examples/error_handling/main.go](error_handling/main.go): Shows how errors are propagated through the step chain
- [examples/error_handling/error_handling_test.go](error_handling/error_handling_test.go): Tests for the error handling example

## Running Examples

To run an example, navigate to its directory and use:

```bash
# Run the main program
go run main.go

# Run the tests
go test -v
