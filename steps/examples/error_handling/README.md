# Error Handling Example

This example demonstrates how to handle errors in a steps chain.

## Overview

The example shows:
1. How to create steps that may fail
2. How errors are propagated through the chain
3. How to track and count errors

## Key Concepts

- Error Propagation: Errors from any step are propagated through the chain
- Early Termination: When a step fails, no subsequent steps are executed
- Error Tracking: Tracking the number of errors that occurred

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
Error: data verification failed
Value: 200
Steps completed: [SetInitialValue ProcessData VerifyData]
Error count: 1
```

## Step Operations

The steps perform the following operations:
1. `SetInitialValue`: Sets the initial value to 100
2. `ProcessData`: Multiplies the value by 2
3. `VerifyData`: Simulates a data verification failure

The example demonstrates that:
1. Steps execute in order until an error occurs
2. Error count is tracked
3. Final value is preserved even if verification fails
