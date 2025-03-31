package steps

import (
	"context"
)

type StepHandler func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error)

// RunnableInterface represents a single unit of work, that can be executed
// within a given context, and specified data. It can work wuth the data
// and return the result of the work.
//
// It can be used as a single step, or combined with other nodes to form
// a Pipeline, Workflow or DAG.
type RunnableInterface interface {
	GetID() string
	SetID(id string)
	GetName() string
	SetName(name string)
	Run(ctx context.Context, data map[string]any) (context.Context, map[string]any, error)
}

// StepInterface represents a single node in a Pipeline, Workflow or DAG.
// A step is a unit of work that can be executed within a given context.
// A step is executed by a Pipeline, Workflow or DAG which manages
// its dependencies and execution order.
type StepInterface interface {
	RunnableInterface

	// GetHandler returns the function that implements the step's execution logic.
	GetHandler() StepHandler

	// SetHandler allows setting or modifying the step's execution logic.
	SetHandler(handler StepHandler)
}

// DagInterface represents a Directed Acyclic Graph (DAG) of steps that can be executed in a specific order.
// It manages the dependencies between steps and ensures they are executed in the correct sequence.
type DagInterface interface {
	RunnableInterface

	// RunnableAdd adds a single node to the DAG.
	// Runnable nodes can be added in any order, as their execution order will be determined by their dependencies.
	RunnableAdd(node ...RunnableInterface)

	// RunnableRemove removes a node from the DAG.
	// Returns true if the node was found and removed, false if it wasn't found.
	RunnableRemove(node RunnableInterface) bool

	// RunnableList returns all runnable nodes in the DAG.
	// The order of nodes in the returned slice is not guaranteed to be their execution order.
	// Use Run() to execute nodes in the correct order based on their dependencies.
	RunnableList() []RunnableInterface

	// DependencyAdd adds a dependency between two nodes.
	// The dependent node will only execute after the dependency node has completed successfully.
	DependencyAdd(dependent RunnableInterface, dependency ...RunnableInterface)

	// DependencyAddIf adds a conditional dependency between nodes.
	// The dependency will only exist if the condition function returns true.
	DependencyAddIf(dependent RunnableInterface, dependency RunnableInterface, condition func(context.Context, map[string]any) bool)

	// DependencyList returns all dependencies for a given node.
	// The actual dependencies may vary based on the context and any conditional dependencies.
	DependencyList(ctx context.Context, node RunnableInterface, data map[string]any) []RunnableInterface
}

// PipelineInterface represents a sequence of runnable nodes
// that will be executed in the sequence they are added.
type PipelineInterface interface {
	RunnableInterface

	// RunnableAdd adds a runnable node(s) to the pipeline.
	RunnableAdd(node ...RunnableInterface)

	// RunnableRemove removes a runnable node from the pipeline.
	RunnableRemove(node RunnableInterface) bool

	// RunnableList returns all runnable nodes in the pipeline.
	// The order of nodes in the returned slice is the order they were added.
	RunnableList() []RunnableInterface
}
