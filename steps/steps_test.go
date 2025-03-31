package steps

import (
	"context"
	"errors"
	"testing"
)

func TestDag_Run(t *testing.T) {
	// Test successful steps
	step1 := NewStep()
	step1.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		data["value"] = 1
		return ctx, data, nil
	})

	step2 := NewStep()
	step2.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		if _, ok := data["value"]; !ok {
			return ctx, data, errors.New("value not found")
		}
		value := data["value"].(int)
		data["value"] = value * 2
		return ctx, data, nil
	})

	// Test DAG with multiple steps
	t.Run("Test successful DAG execution", func(t *testing.T) {
		dag := NewDag()
		dag.RunnableAdd(step1, step2)
		dag.DependencyAdd(step2, step1) // Ensure step2 runs after step1

		ctx := context.Background()
		data := make(map[string]any)
		ctx, data, err := dag.Run(ctx, data)
		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}
		if data["value"].(int) != 2 {
			t.Errorf("expected value 2, got: %d", data["value"].(int))
		}
	})

	// Test error propagation
	t.Run("Test error propagation", func(t *testing.T) {
		stepWithError := NewStep()
		stepWithError.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
			return ctx, data, errors.New("test error")
		})

		dagWithErr := NewDag()
		dagWithErr.RunnableAdd(stepWithError)

		ctx := context.Background()
		data := make(map[string]any)
		_, _, err := dagWithErr.Run(ctx, data)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	// Test dependency ordering
	t.Run("Test dependency ordering", func(t *testing.T) {
		stepA := NewStep()
		stepA.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
			data["A"] = 1
			return ctx, data, nil
		})

		stepB := NewStep()
		stepB.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
			if _, ok := data["A"]; !ok {
				return ctx, data, errors.New("A not found")
			}
			data["B"] = data["A"].(int) * 2
			return ctx, data, nil
		})

		stepC := NewStep()
		stepC.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
			if _, ok := data["B"]; !ok {
				return ctx, data, errors.New("B not found")
			}
			data["C"] = data["B"].(int) * 3
			return ctx, data, nil
		})

		dag := NewDag()
		dag.RunnableAdd(stepA)
		dag.RunnableAdd(stepB)
		dag.RunnableAdd(stepC)
		dag.DependencyAdd(stepB, stepA)
		dag.DependencyAdd(stepC, stepB)

		ctx := context.Background()
		data := make(map[string]any)
		ctx, data, err := dag.Run(ctx, data)
		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}

		if data["A"].(int) != 1 {
			t.Errorf("expected A=1, got: %d", data["A"].(int))
		}
		if data["B"].(int) != 2 {
			t.Errorf("expected B=2, got: %d", data["B"].(int))
		}
		if data["C"].(int) != 6 {
			t.Errorf("expected C=6, got: %d", data["C"].(int))
		}
	})

	// Test circular dependency
	t.Run("Test circular dependency", func(t *testing.T) {
		step1 := NewStep()
		step1.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
			return ctx, data, nil
		})

		step2 := NewStep()
		step2.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
			return ctx, data, nil
		})

		dag := NewDag()
		dag.RunnableAdd(step1)
		dag.RunnableAdd(step2)
		dag.DependencyAdd(step2, step1)
		dag.DependencyAdd(step1, step2) // Create circular dependency

		ctx := context.Background()
		data := make(map[string]any)
		ctx, data, err := dag.Run(ctx, data)
		if err == nil {
			t.Error("expected error for circular dependency, got nil")
		}
	})

	// Test dependencies
	t.Run("Test step dependencies", func(t *testing.T) {
		stepWithDependency := NewStep()
		stepWithDependency.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
			return ctx, data, nil
		})

		dagWithDependencies := NewDag()
		dagWithDependencies.RunnableAdd(step1)
		dagWithDependencies.RunnableAdd(stepWithDependency)
		dagWithDependencies.DependencyAdd(stepWithDependency, step1)

		ctx := context.Background()
		data := make(map[string]any)
		ctx, data, err := dagWithDependencies.Run(ctx, data)
		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}
	})

	// Test parallel execution
	t.Run("Test parallel step execution", func(t *testing.T) {
		stepParallel1 := NewStep()
		stepParallel1.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
			if _, ok := data["value"]; !ok {
				return ctx, data, errors.New("value not found")
			}
			data["value"] = data["value"].(int) + 1
			return ctx, data, nil
		})

		stepParallel2 := NewStep()
		stepParallel2.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
			if _, ok := data["value"]; !ok {
				return ctx, data, errors.New("value not found")
			}
			data["value"] = data["value"].(int) + 2
			return ctx, data, nil
		})

		dagParallel := NewDag()
		dagParallel.RunnableAdd(stepParallel1)
		dagParallel.RunnableAdd(stepParallel2)

		ctx := context.Background()
		data := make(map[string]any)
		data["value"] = 0 // Initialize the value in the context
		ctx, data, err := dagParallel.Run(ctx, data)
		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}
		if data["value"].(int) != 3 {
			t.Errorf("expected value 3, got: %d", data["value"].(int))
		}
	})

	// Test dependency resolution
	t.Run("Test dependency resolution", func(t *testing.T) {
		stepA_SetOne := NewStep()
		stepA_SetOne.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
			data["A"] = 1
			return ctx, data, nil
		})

		stepB_MultiplyByTwo := NewStep()
		stepB_MultiplyByTwo.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
			if _, ok := data["A"]; !ok {
				return ctx, data, errors.New("A not found")
			}
			data["B"] = data["A"].(int) * 2
			return ctx, data, nil
		})

		stepC_MultiplyByThree := NewStep()
		stepC_MultiplyByThree.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
			if _, ok := data["B"]; !ok {
				return ctx, data, errors.New("B not found")
			}
			data["C"] = data["B"].(int) * 3
			return ctx, data, nil
		})

		dag := NewDag()
		dag.RunnableAdd(stepA_SetOne)
		dag.RunnableAdd(stepB_MultiplyByTwo)
		dag.RunnableAdd(stepC_MultiplyByThree)
		dag.DependencyAdd(stepB_MultiplyByTwo, stepA_SetOne)
		dag.DependencyAdd(stepC_MultiplyByThree, stepB_MultiplyByTwo)

		ctx := context.Background()
		data := make(map[string]any)
		ctx, data, err := dag.Run(ctx, data)
		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}

		if data["A"].(int) != 1 {
			t.Errorf("expected A=1, got: %d", data["A"].(int))
		}
		if data["B"].(int) != 2 {
			t.Errorf("expected B=2, got: %d", data["B"].(int))
		}
		if data["C"].(int) != 6 {
			t.Errorf("expected C=6, got: %d", data["C"].(int))
		}
	})
}

func TestDag_IDHandling(t *testing.T) {
	// Create a new DAG
	dag := NewDag()
	
	// Create a step without ID
	step1 := NewStep()
	step1.SetName("Step 1")
	
	// Add step - should automatically generate an ID
	dag.RunnableAdd(step1)
	if step1.GetID() == "" {
		t.Errorf("Expected step to have an ID after adding to DAG")
	}
	
	// Create another step with the same name
	step2 := NewStep()
	step2.SetName("Step 1")
	
	// Add step - should generate a different ID
	dag.RunnableAdd(step2)
	if step1.GetID() == step2.GetID() {
		t.Errorf("Expected different IDs for steps with same name")
	}
	
	// Try to add a step with empty ID
	step3 := NewStep()
	step3.SetID("")
	step3.SetName("Step 3")
	dag.RunnableAdd(step3)
	if step3.GetID() == "" {
		t.Errorf("Expected step to have an ID after adding to DAG")
	}
}

func TestDag_DependencyManagement(t *testing.T) {
	// Create a new DAG
	dag := NewDag()
	
	// Create steps
	step1 := NewStep()
	step1.SetName("Step 1")
	step2 := NewStep()
	step2.SetName("Step 2")
	step3 := NewStep()
	step3.SetName("Step 3")
	
	// Add steps to DAG
	dag.RunnableAdd(step1, step2, step3)
	
	// Add dependencies
	dag.DependencyAdd(step2, step1) // step2 depends on step1
	dag.DependencyAdd(step3, step1, step2) // step3 depends on both step1 and step2
	
	// Verify dependencies
	deps1 := dag.DependencyList(context.Background(), step1, make(map[string]any))
	if len(deps1) != 0 {
		t.Errorf("Expected no dependencies for step1, got %d", len(deps1))
	}
	
	deps2 := dag.DependencyList(context.Background(), step2, make(map[string]any))
	if len(deps2) != 1 {
		t.Errorf("Expected 1 dependency for step2, got %d", len(deps2))
	}
	if deps2[0].GetName() != "Step 1" {
		t.Errorf("Expected dependency to be Step 1")
	}
	
	deps3 := dag.DependencyList(context.Background(), step3, make(map[string]any))
	if len(deps3) != 2 {
		t.Errorf("Expected 2 dependencies for step3, got %d", len(deps3))
	}
	
	// Test conditional dependency
	alwaysTrue := func(ctx context.Context, data map[string]any) bool { return true }
	dag.DependencyAddIf(step3, step1, alwaysTrue)
	
	// Test duplicate dependency (should be ignored)
	dag.DependencyAddIf(step3, step1, alwaysTrue)
	
	// Test invalid dependency (should be ignored)
	invalidStep := NewStep()
	invalidStep.SetID("")
	dag.DependencyAddIf(step3, invalidStep, alwaysTrue)
}

func TestDag_Removal(t *testing.T) {
	// Create a new DAG
	dag := NewDag()
	
	// Create steps
	step1 := NewStep()
	step1.SetName("Step 1")
	step2 := NewStep()
	step2.SetName("Step 2")
	
	// Add steps to DAG
	dag.RunnableAdd(step1, step2)
	
	// Add dependency
	dag.DependencyAdd(step2, step1)
	
	// Remove step1
	removed := dag.RunnableRemove(step1)
	if !removed {
		t.Errorf("Expected step1 to be removed")
	}
	
	// Verify step1 is removed from list
	runners := dag.RunnableList()
	found := false
	for _, runner := range runners {
		if runner.GetID() == step1.GetID() {
			found = true
			break
		}
	}
	if found {
		t.Errorf("Expected step1 to be removed from runnables")
	}
	
	// Verify step2's dependency on step1 is removed
	deps := dag.DependencyList(context.Background(), step2, make(map[string]any))
	if len(deps) != 0 {
		t.Errorf("Expected step2 to have no dependencies after step1 was removed")
	}
}
