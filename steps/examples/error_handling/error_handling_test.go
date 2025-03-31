package main

import (
	"fmt"
	"testing"

	"github.com/dracory/base/steps"
)

func TestErrorHandling(t *testing.T) {
	// Create steps that may fail
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

	// Create a step that will fail
	stepWithError := steps.NewStep(func(ctx steps.StepContextInterface) (steps.StepContextInterface, error) {
		return ctx, fmt.Errorf("intentional error")
	})

	// Create a DAG with the steps
	dag := steps.NewDag()
	dag.AddStep(step1)
	dag.AddStep(step2)
	dag.AddStep(stepWithError)

	// Run the DAG
	ctx := steps.NewStepContext()
	ctx, err := dag.Run(ctx)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	// Verify the error message
	if err.Error() != "intentional error" {
		t.Errorf("Expected error 'intentional error', got '%v'", err)
	}

	// Verify the value was still processed
	value := ctx.Get("value")
	if value != 2 {
		t.Errorf("Expected value 2, got %v", value)
	}
}
