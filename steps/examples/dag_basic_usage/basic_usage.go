package basic_usage

import (
	"context"

	"github.com/dracory/base/steps"
)

// NewSetValueStep creates a new step that sets a value
func NewSetValueStep() steps.StepInterface {
	step := steps.NewStep()
	step.SetName("Set Value")
	step.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		data["value"] = 42
		return ctx, data, nil
	})
	return step
}

// NewIncrementStep creates a new step that increments a value
func NewIncrementStep() steps.StepInterface {
	step := steps.NewStep()
	step.SetName("Increment Value")
	step.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		value := data["value"].(int)
		value++
		data["value"] = value
		return ctx, data, nil
	})
	return step
}

// NewMultipleIncrementDag creates a DAG with multiple increment steps
func NewMultipleIncrementDag() steps.DagInterface {
	dag := steps.NewDag()
	dag.SetName("Multiple Increment DAG")
	
	// Add 4 increment steps
	for i := 0; i < 4; i++ {
		dag.RunnableAdd(NewIncrementStep())
	}
	
	return dag
}
