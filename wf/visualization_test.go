package wf

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

func TestPipelineVisualization(t *testing.T) {
	// Create a pipeline with three steps
	pipeline := NewPipeline()
	step1 := NewStep()
	step2 := NewStep()
	step3 := NewStep()

	step1.SetName("Step 1")
	step2.SetName("Step 2")
	step3.SetName("Step 3")

	step1.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		return ctx, data, nil
	})
	step2.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		return ctx, data, nil
	})
	step3.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		return ctx, data, nil
	})

	pipeline.RunnableAdd(step1, step2, step3)

	// Test empty pipeline
	emptyPipeline := NewPipeline()
	dot := emptyPipeline.Visualize()
	if !strings.Contains(dot, "digraph") {
		t.Error("Empty pipeline visualization should contain 'digraph'")
	}

	// Test pipeline visualization
	dot = pipeline.Visualize()

	// Basic checks
	if !strings.Contains(dot, "digraph") {
		t.Error("Pipeline visualization should contain 'digraph'")
	}
	if !strings.Contains(dot, "rankdir = \"LR\"") {
		t.Error("Pipeline visualization should have left-to-right layout")
	}

	// Check if all steps are present
	if !strings.Contains(dot, "Step 1") || !strings.Contains(dot, "Step 2") || !strings.Contains(dot, "Step 3") {
		t.Error("Pipeline visualization should contain all step names")
	}

	// Check if edges are present
	if !strings.Contains(dot, "->") {
		t.Error("Pipeline visualization should contain edges")
	}

	// Test visualization with current step
	pipeline.GetState().SetCurrentStepID(step2.GetID())
	dot = pipeline.Visualize()
	if !strings.Contains(dot, "#2196F3") {
		t.Error("Current step should be colored blue")
	}

	// Test visualization with completed steps
	pipeline.GetState().SetStatus(StateStatusComplete)
	dot = pipeline.Visualize()
	if !strings.Contains(dot, "#4CAF50") {
		t.Error("Completed steps should be colored green")
	}

	// t.Log(dot)
}

func TestDagVisualization(t *testing.T) {
	// Test empty DAG
	emptyDag := NewDag()
	dot := emptyDag.Visualize()
	if !strings.Contains(dot, "digraph") {
		t.Error("Empty DAG visualization should contain 'digraph'")
	}

	// Create a DAG with multiple steps and complex dependencies
	dag := NewDag()
	step1 := NewStep()
	step2 := NewStep()
	step3 := NewStep()
	step4 := NewStep()
	step5 := NewStep()

	step1.SetName("Step 1")
	step2.SetName("Step 2")
	step3.SetName("Step 3")
	step4.SetName("Step 4")
	step5.SetName("Step 5")

	// Set up handlers for each step
	step1.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		return ctx, data, nil
	})
	step2.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		return ctx, data, nil
	})
	step3.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		return ctx, data, nil
	})
	step4.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		return ctx, data, nil
	})
	step5.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		return ctx, data, nil
	})

	// Add all steps to the DAG
	dag.RunnableAdd(step1, step2, step3, step4, step5)

	// Create a complex dependency structure:
	// Step1 -> Step2 -> Step4
	// Step1 -> Step3 -> Step4
	// Step2 -> Step5
	// Step3 -> Step5
	dag.DependencyAdd(step2, step1)
	dag.DependencyAdd(step3, step1)
	dag.DependencyAdd(step4, step2, step3)
	dag.DependencyAdd(step5, step2, step3)

	// Test DAG visualization
	dot = dag.Visualize()

	// Basic checks
	if !strings.Contains(dot, "digraph") {
		t.Error("DAG visualization should contain 'digraph'")
	}
	if !strings.Contains(dot, "rankdir = \"LR\"") {
		t.Error("DAG visualization should have left-to-right layout")
	}

	// Check if all steps are present
	if !strings.Contains(dot, "Step 1") || !strings.Contains(dot, "Step 2") ||
		!strings.Contains(dot, "Step 3") || !strings.Contains(dot, "Step 4") ||
		!strings.Contains(dot, "Step 5") {
		t.Error("DAG visualization should contain all step names")
	}

	// Check if dependencies are represented as edges
	if !strings.Contains(dot, "->") {
		t.Error("DAG visualization should contain edges for dependencies")
	}

	// Verify specific dependencies are present in the visualization
	dependencies := []string{
		fmt.Sprintf("\"%s\" -> \"%s\"", step1.GetID(), step2.GetID()),
		fmt.Sprintf("\"%s\" -> \"%s\"", step1.GetID(), step3.GetID()),
		fmt.Sprintf("\"%s\" -> \"%s\"", step2.GetID(), step4.GetID()),
		fmt.Sprintf("\"%s\" -> \"%s\"", step3.GetID(), step4.GetID()),
		fmt.Sprintf("\"%s\" -> \"%s\"", step2.GetID(), step5.GetID()),
		fmt.Sprintf("\"%s\" -> \"%s\"", step3.GetID(), step5.GetID()),
	}

	for _, dep := range dependencies {
		if !strings.Contains(dot, dep) {
			t.Errorf("DAG visualization should contain dependency: %s", dep)
		}
	}

	// Test visualization with current step
	dag.GetState().SetCurrentStepID(step2.GetID())
	dot = dag.Visualize()
	if !strings.Contains(dot, "#2196F3") {
		t.Error("Current step should be colored blue")
	}

	t.Log(dot)

	// Test visualization with completed steps
	dag.GetState().SetStatus(StateStatusComplete)
	dot = dag.Visualize()
	if !strings.Contains(dot, "#4CAF50") {
		t.Error("Completed steps should be colored green")
	}

	// Test visualization with a failed step
	dag.GetState().SetStatus(StateStatusFailed)
	dag.GetState().SetCurrentStepID(step4.GetID())
	dag.GetState().AddCompletedStep(step1.GetID())
	dag.GetState().AddCompletedStep(step2.GetID())
	dag.GetState().AddCompletedStep(step3.GetID())
	dot = dag.Visualize()
	if !strings.Contains(dot, "#F44336") {
		t.Error("Failed step should be colored red")
	}

	// Test visualization with a paused step
	dag.GetState().SetStatus(StateStatusPaused)
	dag.GetState().SetCurrentStepID(step3.GetID())
	dag.GetState().AddCompletedStep(step1.GetID())
	dag.GetState().AddCompletedStep(step2.GetID())
	dot = dag.Visualize()
	if !strings.Contains(dot, "#FFC107") {
		t.Error("Paused step should be colored yellow")
	}

	t.Log(dot)
}

func TestStepVisualization(t *testing.T) {
	// Create a step
	step := NewStep()
	step.SetName("Step 1")
	step.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		return ctx, data, nil
	})

	// Test step visualization
	dot := step.Visualize()

	// Basic checks
	if !strings.Contains(dot, "digraph") {
		t.Error("Step visualization should contain 'digraph'")
	}
	if !strings.Contains(dot, "rankdir = \"LR\"") {
		t.Error("Step visualization should have left-to-right layout")
	}

	// Check if step name is present
	if !strings.Contains(dot, "Step 1") {
		t.Error("Step visualization should contain step name")
	}

	// Test visualization with different states
	step.GetState().SetStatus(StateStatusRunning)
	dot = step.Visualize()
	if !strings.Contains(dot, "#2196F3") {
		t.Error("Running step should be colored blue")
	}

	step.GetState().SetStatus(StateStatusComplete)
	dot = step.Visualize()
	if !strings.Contains(dot, "#4CAF50") {
		t.Error("Completed step should be colored green")
	}

	step.GetState().SetStatus(StateStatusFailed)
	dot = step.Visualize()
	if !strings.Contains(dot, "#F44336") {
		t.Error("Failed step should be colored red")
	}

	step.GetState().SetStatus(StateStatusPaused)
	dot = step.Visualize()
	if !strings.Contains(dot, "#FFC107") {
		t.Error("Paused step should be colored yellow")
	}
}
