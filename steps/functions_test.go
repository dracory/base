package steps

import (
	"context"
	"testing"
)

func TestVisitNode(t *testing.T) {
	// Create test steps
	step1 := NewStep()
	step2 := NewStep()
	step3 := NewStep()
	step1.SetName("Step1")
	step2.SetName("Step2")
	step3.SetName("Step3")

	// Set handlers for steps
	step1.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		return ctx, data, nil
	})
	step2.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		return ctx, data, nil
	})
	step3.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		return ctx, data, nil
	})

	// Test regular dependency chain
	graph := map[RunnableInterface][]RunnableInterface{
		step1: {step2},
		step2: {step3},
		step3: {},
	}

	visited := make(map[RunnableInterface]bool)
	tempMark := make(map[RunnableInterface]bool)
	result := []RunnableInterface{}

	if err := visitNode(step1, graph, visited, tempMark, &result); err != nil {
		t.Errorf("visitNode failed for regular dependency chain: %v", err)
	}

	if len(result) != 3 {
		t.Errorf("Expected 3 nodes in result, got %d", len(result))
	}

	// Test cycle detection
	graphWithCycle := map[RunnableInterface][]RunnableInterface{
		step1: {step2},
		step2: {step3},
		step3: {step1},
	}

	visited = make(map[RunnableInterface]bool)
	tempMark = make(map[RunnableInterface]bool)
	result = []RunnableInterface{}

	if err := visitNode(step1, graphWithCycle, visited, tempMark, &result); err == nil {
		t.Error("Expected cycle detection error, got nil")
	} else if err.Error() != "cycle detected" {
		t.Errorf("Expected cycle detected error, got: %v", err)
	}

	// Test conditional dependencies
	graphWithConditional := map[RunnableInterface][]RunnableInterface{
		step1: {step2},
		step2: {},
		step3: {},
	}

	visited = make(map[RunnableInterface]bool)
	tempMark = make(map[RunnableInterface]bool)
	result = []RunnableInterface{}

	if err := visitNode(step1, graphWithConditional, visited, tempMark, &result); err != nil {
		t.Errorf("visitNode failed for conditional dependencies: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 nodes in result, got %d", len(result))
	}
}

func TestTopologicalSort(t *testing.T) {
	// Create test steps
	step1 := NewStep()
	step2 := NewStep()
	step3 := NewStep()
	step4 := NewStep()
	step1.SetName("Step1")
	step2.SetName("Step2")
	step3.SetName("Step3")
	step4.SetName("Step4")

	// Set handlers for steps
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

	// Test simple linear dependency
	graph := map[RunnableInterface][]RunnableInterface{
		step1: {step2},
		step2: {step3},
		step3: {},
	}

	result, err := topologicalSort(graph)
	if err != nil {
		t.Errorf("topologicalSort failed for simple linear dependency: %v", err)
	}

	if len(result) != 3 {
		t.Errorf("Expected 3 nodes in result, got %d", len(result))
	}

	// Test parallel dependencies
	graphParallel := map[RunnableInterface][]RunnableInterface{
		step1: {step2, step3},
		step2: {step4},
		step3: {step4},
		step4: {},
	}

	result, err = topologicalSort(graphParallel)
	if err != nil {
		t.Errorf("topologicalSort failed for parallel dependencies: %v", err)
	}

	if len(result) != 4 {
		t.Errorf("Expected 4 nodes in result, got %d", len(result))
	}

	// Test cycle detection
	graphWithCycle := map[RunnableInterface][]RunnableInterface{
		step1: {step2},
		step2: {step3},
		step3: {step1},
	}

	_, err = topologicalSort(graphWithCycle)
	if err == nil {
		t.Error("Expected cycle detection error, got nil")
	} else if err.Error() != "cycle detected" {
		t.Errorf("Expected cycle detected error, got: %v", err)
	}

	// Test disconnected nodes
	graphDisconnected := map[RunnableInterface][]RunnableInterface{
		step1: {step2},
		step3: {},
		step4: {},
	}

	result, err = topologicalSort(graphDisconnected)
	if err != nil {
		t.Errorf("topologicalSort failed for disconnected nodes: %v", err)
	}

	if len(result) != 4 {
		t.Errorf("Expected 4 nodes in result, got %d", len(result))
	}
}
