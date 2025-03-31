package basic_usage

import (
	"testing"
	"github.com/dracory/base/steps"
)

func TestBasicUsage(t *testing.T) {
	// Create and run the DAG
	ctx := steps.NewStepContext()
	ctx.Set("value", 0)

	dag := steps.NewDag()
	dag.AddStep(NewIncrementStep())
	dag.AddStep(NewIncrementStep())
	dag.AddStep(NewIncrementStep())
	dag.AddStep(NewIncrementStep())

	ctx, err := dag.Run(ctx)
	if err != nil {
		t.Errorf("Error running DAG: %v", err)
		return
	}

	// Verify the value
	value := ctx.Get("value")
	if value != 4 {
		t.Errorf("Expected value 4, got %v", value)
	}
}
