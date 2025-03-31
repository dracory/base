package main

import (
	"fmt"
	"testing"

	"github.com/dracory/base/steps"
)

func TestDependencies(t *testing.T) {
	// Create steps
	step1 := steps.NewStep(func(ctx steps.StepContextInterface) (steps.StepContextInterface, error) {
		ctx.Set("value", 1)
		return ctx, nil
	})

	step2 := steps.NewStep(func(ctx steps.StepContextInterface) (steps.StepContextInterface, error) {
		if !ctx.Has("value") {
			return ctx, fmt.Errorf("value not found")
		}
		value := ctx.Get("value").(int)
		ctx.Set("value", value*2)
		return ctx, nil
	})

	step3 := steps.NewStep(func(ctx steps.StepContextInterface) (steps.StepContextInterface, error) {
		if !ctx.Has("value") {
			return ctx, fmt.Errorf("value not found")
		}
		value := ctx.Get("value").(int)
		ctx.Set("value", value*3)
		return ctx, nil
	})

	// Create a DAG and add dependencies
	dag := steps.NewDag()
	dag.AddStep(step1)
	dag.AddStep(step2)
	dag.AddStep(step3)
	dag.AddDependency(step2, step1)
	dag.AddDependency(step3, step2)

	// Run the DAG
	ctx := steps.NewStepContext()
	ctx, err := dag.Run(ctx)
	if err != nil {
		t.Errorf("Error running DAG: %v", err)
		return
	}

	// Verify the value
	value := ctx.Get("value")
	if value != 6 {
		t.Errorf("Expected value 6, got %v", value)
	}
}
