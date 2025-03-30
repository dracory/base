package steps

import (
	"github.com/dracory/base/object"
)

// StepContextInterface represents the context in which steps are executed.
// It must implement DataObjectInterface to provide data access capabilities.
type StepContextInterface interface {
	object.SerializablePropertyObjectInterface
	
	// Name returns the name of the context
	Name() string
	
	// SetName sets the name of the context
	SetName(name string) StepContextInterface
}

// StepInterface represents a single node in a DAG.
// A step is a unit of work that can be executed within a given context.
// Steps can have dependencies on other steps and maintain their own execution logic.
type StepInterface interface {
	object.SerializablePropertyObjectInterface

	// GetExecute returns the function that implements the step's execution logic.
	// This function takes a StepContextInterface and returns an error if the step fails.
	GetExecute() func(ctx StepContextInterface) error

	// SetExecute allows setting or modifying the step's execution logic.
	// Returns the step itself to support method chaining.
	SetExecute(fn func(ctx StepContextInterface) error) StepInterface

	// AddDependency adds a single dependency on another step.
	// A step will only execute after all its dependencies have completed successfully.
	// Returns the step itself to support method chaining.
	AddDependency(step StepInterface) StepInterface

	// AddDependencies adds multiple dependencies at once.
	// This is more efficient than calling AddDependency multiple times.
	// Returns the step itself to support method chaining.
	AddDependencies(steps ...StepInterface) StepInterface

	// AddDependencyIf adds a dependency that only exists if the condition is true.
	// Returns the step itself to support method chaining.
	AddDependencyIf(step StepInterface, condition func(ctx StepContextInterface) bool) StepInterface

	// GetDependencies returns a list of all steps that this step depends on.
	// The actual dependencies may vary based on the context and any conditional dependencies.
	GetDependencies(ctx StepContextInterface) []StepInterface

	// Run executes the step's logic within the given context.
	// Returns an error if the step fails to execute.
	Run(ctx StepContextInterface) error

	// Name returns the name of the step.
	// This is used for identifying steps in the DAG.
	Name() string

	// SetName sets the name of the step.
	// Returns the step itself to support method chaining.
	SetName(name string) StepInterface
}

// DagInterface represents a Directed Acyclic Graph (DAG) of steps that can be executed in a specific order.
// It manages the dependencies between steps and ensures they are executed in the correct sequence.
type DagInterface interface {
	object.SerializablePropertyObjectInterface

	// AddStep adds a single step to the DAG.
	// Steps can be added in any order, as their execution order will be determined by their dependencies.
	AddStep(step StepInterface)

	// AddSteps adds multiple steps to the DAG at once.
	// This is more efficient than calling AddStep multiple times.
	// Steps can be added in any order, as their execution order will be determined by their dependencies.
	AddSteps(steps ...StepInterface)

	// RemoveStep removes a step from the DAG.
	// Returns true if the step was found and removed, false if it wasn't found.
	RemoveStep(step StepInterface) bool

	// GetSteps returns all steps in the DAG.
	// The order of steps in the returned slice is not guaranteed to be their execution order.
	// Use Run() to execute steps in the correct order based on their dependencies.
	GetSteps() []StepInterface

	// Run executes all steps in the DAG within the given context.
	// Steps are executed in topological order based on their dependencies.
	// If any step fails, the execution stops and the error is returned.
	Run(ctx StepContextInterface) error
}
