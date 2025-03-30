# Steps Package

The steps package provides a flexible and extensible framework for defining and executing sequential operations in Go. It implements a DAG (Directed Acyclic Graph) pattern for managing dependencies between operations.

## Key Features

- **Step-based Execution**: Define operations as individual steps that can be executed in sequence
- **Dependency Management**: Steps can depend on other steps, ensuring proper execution order
- **Cycle Detection**: Automatically detects and prevents circular dependencies
- **Context Management**: Each step receives a context object that can be used to share data between steps
- **Error Propagation**: Errors are properly propagated through the step chain
- **Testable**: Designed with testing in mind, using real database connections and avoiding mocks

## Interfaces

### StepContextInterface
```go
type StepContextInterface interface {
    object.SerializablePropertyObjectInterface
}
```
- Base interface for step execution context
- Extends `object.SerializablePropertyObjectInterface` for data management

### StepInterface
```go
type StepInterface interface {
    object.SerializablePropertyObjectInterface

    // GetExecute returns the function that implements the step's execution logic.
    GetExecute() func(ctx StepContextInterface) error

    // SetExecute allows setting or modifying the step's execution logic.
    SetExecute(fn func(ctx StepContextInterface) error) StepInterface

    // AddDependency adds a single dependency on another step.
    AddDependency(step StepInterface) StepInterface

    // AddDependencies adds multiple dependencies at once.
    AddDependencies(steps ...StepInterface) StepInterface

    // GetDependencies returns a list of all steps that this step depends on.
    GetDependencies() []StepInterface

    // Run executes the step's logic within the given context.
    Run(ctx StepContextInterface) error

    // Name returns the name of the step.
    Name() string

    // SetName sets the name of the step.
    SetName(name string) StepInterface
}
```
- Defines a single execution step
- Supports dependency management
- Provides execution and error handling

### DagInterface
```go
type DagInterface interface {
    // AddStep adds a single step to the DAG.
    AddStep(step StepInterface)

    // AddSteps adds multiple steps to the DAG at once.
    AddSteps(steps ...StepInterface)

    // RemoveStep removes a step from the DAG.
    RemoveStep(step StepInterface) bool

    // GetSteps returns all steps in the DAG.
    GetSteps() []StepInterface

    // Run executes all steps in the DAG within the given context.
    Run(ctx StepContextInterface) error
}
```
- Manages a collection of steps
- Executes steps in dependency order
- Handles error propagation

## Usage

### Creating Steps
```go
// Create a simple step
step := NewStep(func(ctx StepContextInterface) error {
    ctx.Set("key", "value")
    return nil
})

// Create a step with dependencies
step1 := NewStep(func(ctx StepContextInterface) error {
    ctx.Set("A", 1)
    return nil
})

step2 := NewStep(func(ctx StepContextInterface) error {
    value := ctx.Get("A").(int)
    ctx.Set("B", value*2)
    return nil
}).AddDependency(step1)
```

### Executing Steps
```go
// Create a DAG
dag := Dag()

// Add steps
dag.AddStep(step1)
dag.AddStep(step2)

// Create a context
ctx := NewStepContext()

// Execute all steps
err := dag.Run(ctx)
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

- `github.com/dracory/base/object`: For property object and serialization functionality

## Best Practices

1. Always use the provided interfaces for type safety
2. Handle errors appropriately in step implementations
3. Use the context for data sharing between steps
4. Define dependencies when steps must be executed in a specific order
5. Avoid creating circular dependencies between steps

## Example

```go
// Create a test context
type MyContext struct {
    object.SerializablePropertyObject
    value int
}

// Create steps
step1 := NewStep(func(ctx StepContextInterface) error {
    ctx.(*MyContext).value = 1
    return nil
})

step2 := NewStep(func(ctx StepContextInterface) error {
    ctx.(*MyContext).value = 2
    return nil
})

// Set up dependencies
step2.AddDependency(step1) // step2 depends on step1

// Create and run steps
dag := Dag()
dag.AddStep(step1)
dag.AddStep(step2)

ctx := &MyContext{}
err := dag.Run(ctx)
if err != nil {
    // Handle error
}
// ctx.value will be 2 after execution
```

## Error Handling

The package will return errors in the following cases:
- If a cycle is detected in the dependency graph
- If any step execution fails
- If a step is added multiple times
- If dependencies are not properly defined
