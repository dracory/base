# Steps Package

The steps package provides a flexible and extensible framework for defining and executing sequential operations in Go. It implements a DAG (Directed Acyclic Graph) pattern for managing dependencies between operations.

## Key Features

- **Simple Step Definitions**: Easily define individual operations as reusable steps
- **Organized Pipelines**: Group related operations into logical pipelines for better maintainability
- **Flexible Dependencies**: Create complex workflows with step dependencies
- **Cycle Detection**: Automatically detects and prevents circular dependencies
- **Context Management**: Share data between steps using a context object
- **Error Handling**: Proper error propagation through the entire workflow
- **Testable**: Designed with testing in mind

## Core Components

- [Step](https://github.com/dracory/base/blob/main/steps/step.go): Represents a single execution step with unique ID, name, and execution handler
- [Pipeline](https://github.com/dracory/base/blob/main/steps/pipeline.go): Groups related steps into a logical unit that can be treated as a single step
- [Dag](https://github.com/dracory/base/blob/main/steps/dag.go): Manages a collection of steps and their dependencies, executing them in the correct order

## Component Hierarchy

```
Runnable
├── Step (basic unit of work, single operation)
├── Pipeline (runs a set of runnables in the sequence they are added)
└── Dag (advanced workflow manager with dependencies between runnables)
```

## Usage Examples

### Creating Steps
```go
// Create a step with an execution function
step := NewStep()
step.SetName("My Step")
step.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
    data["key"] = "value"
    return ctx, data, nil
})
```

### Creating a Pipeline
```go
// Create steps for a pipeline
step1 := NewStep()
step1.SetName("Process Data")
step1.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
    data["processed"] = true
    return ctx, data, nil
})

step2 := NewStep()
step2.SetName("Validate Data")
step2.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
    if !data["processed"].(bool) {
        return ctx, data, errors.New("data not processed")
    }
    return ctx, data, nil
})

// Create a pipeline
pipeline := NewPipeline()
pipeline.SetName("Data Processing Pipeline")

// Add steps to pipeline
pipeline.RunnableAdd(step1, step2)
```

### Creating a DAG
```go
// Create a DAG
dag := NewDag()
dag.SetName("My DAG")

// Add steps
dag.RunnableAdd(step1, step2)

// Add dependencies
dag.DependencyAdd(step2, step1) // step2 depends on step1
```

### Using a Pipeline in a DAG
```go
// Create a pipeline with steps
pipeline := NewPipeline()
pipeline.SetName("Data Processing Pipeline")

// Add steps to pipeline
pipeline.RunnableAdd(step1, step2)

// Create a DAG
dag := NewDag()
dag.SetName("My DAG")

// Add pipeline to DAG
dag.RunnableAdd(pipeline)

// Add other steps that depend on the pipeline
dag.RunnableAdd(step3)
dag.DependencyAdd(step3, pipeline)
```

### Executing Steps
```go
// Create a context and data map
ctx := context.Background()
data := make(map[string]any)

// Execute all steps
_, data, err := dag.Run(ctx, data)
if err != nil {
    // Handle error
}
```

## Testing

The package includes comprehensive tests that verify:
- Successful step execution
- Error propagation
- Dependency handling
- Context data sharing
- Cycle detection
- Parallel execution
- Serialization

## Dependencies

- `github.com/gouniverse/uid`: For generating unique IDs

## Best Practices

1. Always use the provided interfaces for type safety
2. Handle errors appropriately in step implementations
3. Use the context for data sharing between steps
4. Define dependencies when steps must be executed in a specific order
5. Avoid creating circular dependencies between steps
6. Use pipelines to group related steps into logical units
7. Implement proper error handling in each step

## Examples

The package includes several examples demonstrating different use cases:

### Basic Usage Example
- Shows how to create and execute simple steps
- Demonstrates basic step dependencies
- Location: [examples/dag_basic_usage](examples/dag_basic_usage)

### Conditional Logic Example
- Demonstrates how to implement conditional logic using DAGs and pipelines
- Shows different step chains for different scenarios
- Location: [examples/dag_conditional_logic](examples/dag_conditional_logic)

### Dependencies Example
- Shows how to create steps with complex dependencies
- Demonstrates proper execution order through dependencies
- Location: [examples/dag_dependencies](examples/dag_dependencies)

### Error Handling Example
- Demonstrates error handling in a DAG
- Shows how errors are propagated through the DAG
- Location: [examples/dag_error_handling](examples/dag_error_handling)

## Error Handling

The package will return errors in the following cases:
- If a cycle is detected in the dependency graph
- If any step execution fails
- If a step is added multiple times
- If dependencies are not properly defined
- If pipeline execution fails
- If conditional logic conditions are not met
