# Error Handling Example

This example demonstrates advanced error handling in a DAG using the steps package. It shows how errors can be caught, handled gracefully, and how conditional dependencies can be used to control step execution.

## Overview

The example consists of three steps:

1. `InitialValueStep`: Sets an initial value of 1
2. `ProcessDataStep`: Multiplies the value by 2
3. `ErrorStep`: Intentionally fails with an error

The steps are arranged in a DAG with dependencies:

```
InitialValueStep -> ProcessDataStep -> ErrorStep
```

## Key Concepts

- **Advanced Error Handling**: Shows how errors are propagated through the DAG and handled gracefully
- **Conditional Dependencies**: Demonstrates how `DependencyAddIf` can be used to control step execution based on conditions
- **Step Dependencies**: Shows how steps depend on each other
- **Context Management**: Shows how data is passed between steps
- **Graceful Failure**: Demonstrates how the DAG can fail gracefully while still completing successful steps

## Implementation Details

The example showcases several important features:

1. **Error Propagation**: Errors are properly propagated through the DAG while allowing successful steps to complete
2. **Conditional Dependencies**: The `DependencyAddIf` function is used to create a conditional dependency between steps
3. **Data Flow**: Data is passed between steps using the context and data map
4. **Step Isolation**: Each step is isolated and handles its own errors independently

## Example Code

```go
// Create steps
dag := steps.NewDag()
dag.SetName("Error Handling Example DAG")

// Create steps
initialStep := NewInitialValueStep()
processStep := NewProcessDataStep()
errorStep := NewErrorStep()

// Add steps to DAG
dag.RunnableAdd(initialStep, processStep, errorStep)

// Set up dependencies
dag.DependencyAdd(processStep, initialStep)
dag.DependencyAdd(errorStep, processStep)

// Add conditional error handling
dag.DependencyAddIf(errorStep, processStep, func(ctx context.Context, data map[string]any) bool {
    return true // Always allow error step to run
})
```

## Running the Example

To run the example:

```bash
go run main.go
```

The program will output:
```
Error: intentional error
```

The error occurs in the `ErrorStep`, but the previous steps (`InitialValueStep` and `ProcessDataStep`) still complete successfully. This demonstrates how the DAG can handle errors gracefully while still allowing successful steps to complete their work.

## Testing

To run the tests:

```bash
go test -v
```

The tests verify that:
1. The error is properly propagated
2. The value is correctly processed by the successful steps
3. The error message matches expectations
4. Conditional dependencies work as expected

## Best Practices

This example demonstrates several best practices for error handling in DAGs:

1. Use `DependencyAddIf` for conditional step execution
2. Handle errors at the step level
3. Use proper type assertions for data handling
4. Maintain clear step dependencies
5. Keep error messages descriptive and specific
