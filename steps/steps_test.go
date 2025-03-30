package steps

import (
	"errors"
	"testing"
)

func TestDag_Run(t *testing.T) {
	// Test successful steps
	step1 := NewStep(func(ctx StepContextInterface) error {
		ctx.Set("value", 1)
		return nil
	})

	step2 := NewStep(func(ctx StepContextInterface) error {
		if !ctx.Has("value") {
			return errors.New("value not found")
		}
		value := ctx.Get("value").(int)
		ctx.Set("value", value*2)
		return nil
	})

	// Test DAG with multiple steps
	t.Run("Test successful DAG execution", func(t *testing.T) {
		dag := NewDag()
		dag.AddStep(step1)
		dag.AddStep(step2)

		ctx := NewStepContext()
		err := dag.Run(ctx)
		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}
		if ctx.Get("value").(int) != 2 {
			t.Errorf("expected value 2, got: %d", ctx.Get("value").(int))
		}
	})

	// Test error propagation
	t.Run("Test error propagation", func(t *testing.T) {
		stepWithError := NewStep(func(ctx StepContextInterface) error {
			return errors.New("test error")
		})

		dagWithErr := NewDag()
		dagWithErr.AddStep(stepWithError)

		ctx := NewStepContext()
		err := dagWithErr.Run(ctx)
		if err == nil {
			t.Error("expected error, got nil")
		}
		if err.Error() != "test error" {
			t.Errorf("expected error 'test error', got: %v", err)
		}
	})

	// Test dependency ordering
	t.Run("Test dependency ordering", func(t *testing.T) {
		stepA := NewStep(func(ctx StepContextInterface) error {
			ctx.Set("A", 1)
			return nil
		})

		stepB := NewStep(func(ctx StepContextInterface) error {
			if !ctx.Has("A") {
				return errors.New("A not found")
			}
			ctx.Set("B", ctx.Get("A").(int)*2)
			return nil
		}).AddDependency(stepA)

		stepC := NewStep(func(ctx StepContextInterface) error {
			if !ctx.Has("B") {
				return errors.New("B not found")
			}
			ctx.Set("C", ctx.Get("B").(int)*3)
			return nil
		}).AddDependency(stepB)

		dag := NewDag()
		dag.AddStep(stepA)
		dag.AddStep(stepB)
		dag.AddStep(stepC)

		ctx := NewStepContext()
		err := dag.Run(ctx)
		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}

		if ctx.Get("A").(int) != 1 {
			t.Errorf("expected A=1, got: %d", ctx.Get("A").(int))
		}
		if ctx.Get("B").(int) != 2 {
			t.Errorf("expected B=2, got: %d", ctx.Get("B").(int))
		}
		if ctx.Get("C").(int) != 6 {
			t.Errorf("expected C=6, got: %d", ctx.Get("C").(int))
		}
	})

	// Test circular dependency detection
	t.Run("Test circular dependency", func(t *testing.T) {
		step1 := NewStep(func(ctx StepContextInterface) error {
			return nil
		})

		step2 := NewStep(func(ctx StepContextInterface) error {
			return nil
		}).AddDependency(step1)

		// Create circular dependency
		step1.AddDependency(step2)

		dag := NewDag()
		dag.AddStep(step1)
		dag.AddStep(step2)

		ctx := NewStepContext()
		err := dag.Run(ctx)
		if err == nil {
			t.Error("expected error for circular dependency, got nil")
		}
	})

	// Test dependencies
	t.Run("Test step dependencies", func(t *testing.T) {
		stepWithDependency := NewStep(func(ctx StepContextInterface) error {
			return nil
		})

		stepWithDependency.AddDependency(step1)

		dagWithDependencies := NewDag()
		dagWithDependencies.AddStep(step1)
		dagWithDependencies.AddStep(stepWithDependency)

		ctx := NewStepContext()
		err := dagWithDependencies.Run(ctx)
		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}
	})

	// Test serialization
	t.Run("Test step serialization", func(t *testing.T) {
		stepToSerialize := NewStep(func(ctx StepContextInterface) error {
			return nil
		})

		jsonBytes, err := stepToSerialize.ToJSON()
		if err != nil {
			t.Errorf("expected no error serializing, got: %v", err)
		}

		newStep := NewStep(func(ctx StepContextInterface) error {
			return nil
		})
		err = newStep.FromJSON(jsonBytes)
		if err != nil {
			t.Errorf("expected no error deserializing, got: %v", err)
		}

		if newStep.GetID() != stepToSerialize.GetID() {
			t.Errorf("expected same ID after serialization, got: %s", newStep.GetID())
		}
	})

	// Test parallel execution
	t.Run("Test parallel step execution", func(t *testing.T) {
		stepParallel1 := NewStep(func(ctx StepContextInterface) error {
			if !ctx.Has("value") {
				return errors.New("value not found")
			}
			ctx.Set("value", ctx.Get("value").(int)+1)
			return nil
		})

		stepParallel2 := NewStep(func(ctx StepContextInterface) error {
			if !ctx.Has("value") {
				return errors.New("value not found")
			}
			ctx.Set("value", ctx.Get("value").(int)+2)
			return nil
		})

		dagParallel := NewDag()
		dagParallel.AddStep(stepParallel1)
		dagParallel.AddStep(stepParallel2)

		ctx := NewStepContext()
		ctx.Set("value", 0) // Initialize the value in the context
		err := dagParallel.Run(ctx)
		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}
		if ctx.Get("value").(int) != 3 {
			t.Errorf("expected value 3, got: %d", ctx.Get("value").(int))
		}
	})

	// Test dependency resolution
	t.Run("Test dependency resolution", func(t *testing.T) {
		stepA_SetOne := NewStep(func(ctx StepContextInterface) error {
			ctx.Set("A", 1)
			return nil
		})

		stepB_MultiplyByTwo := NewStep(func(ctx StepContextInterface) error {
			if !ctx.Has("A") {
				return errors.New("A not found")
			}
			ctx.Set("B", ctx.Get("A").(int)*2)
			return nil
		}).AddDependency(stepA_SetOne)

		stepC_MultiplyByThree := NewStep(func(ctx StepContextInterface) error {
			if !ctx.Has("B") {
				return errors.New("B not found")
			}
			ctx.Set("C", ctx.Get("B").(int)*3)
			return nil
		}).AddDependency(stepB_MultiplyByTwo)

		dag := NewDag()
		dag.AddStep(stepA_SetOne)
		dag.AddStep(stepB_MultiplyByTwo)
		dag.AddStep(stepC_MultiplyByThree)

		ctx := NewStepContext()
		err := dag.Run(ctx)
		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}

		if ctx.Get("A").(int) != 1 {
			t.Errorf("expected A=1, got: %d", ctx.Get("A").(int))
		}
		if ctx.Get("B").(int) != 2 {
			t.Errorf("expected B=2, got: %d", ctx.Get("B").(int))
		}
		if ctx.Get("C").(int) != 6 {
			t.Errorf("expected C=6, got: %d", ctx.Get("C").(int))
		}
	})
}
