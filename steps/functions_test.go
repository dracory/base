package steps

import (
	"context"
	"testing"
)

func Test_VisitNode(t *testing.T) {
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

func Test_TopologicalSort(t *testing.T) {
	// Create test steps
	step1 := NewStep()
	step1.SetName("Step1")
	step2 := NewStep()
	step2.SetName("Step2")
	step3 := NewStep()
	step3.SetName("Step3")
	step4 := NewStep()
	step4.SetName("Step4")

	// Create graph with dependencies
	graph := map[RunnableInterface][]RunnableInterface{
		step1: {},
		step2: {step1},
		step3: {step2},
		step4: {step1},
	}

	// Test case 1: Regular dependency chain
	result, err := topologicalSort(graph)
	if err != nil {
		t.Errorf("topologicalSort failed: %v", err)
	}

	// Verify result order
	if len(result) != 4 {
		t.Errorf("Expected 4 nodes in result, got %d", len(result))
	}

	// Verify step1 is first since it has no dependencies
	if result[0] != step1 {
		t.Errorf("Expected step1 to be first, got %s", result[0].GetName())
	}

	// Verify step2 comes after step1
	if result[1] != step2 {
		t.Errorf("Expected step2 to be second, got %s", result[1].GetName())
	}

	// Verify step3 comes after step2
	if result[2] != step3 {
		t.Errorf("Expected step3 to be third, got %s", result[2].GetName())
	}

	// Verify step4 comes after step1
	if result[3] != step4 {
		t.Errorf("Expected step4 to be fourth, got %s", result[3].GetName())
	}

	// Test case 2: Circular dependency
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

	// Test case 3: Multiple independent chains
	graphWithMultipleChains := map[RunnableInterface][]RunnableInterface{
		step1: {},
		step2: {step1},
		step3: {},
		step4: {step3},
	}

	result, err = topologicalSort(graphWithMultipleChains)
	if err != nil {
		t.Errorf("topologicalSort failed: %v", err)
	}

	// Verify result order - independent chains should be sorted by name
	if len(result) != 4 {
		t.Errorf("Expected 4 nodes in result, got %d", len(result))
	}

	// Verify step1 is first since it has no dependencies
	if result[0] != step1 {
		t.Errorf("Expected step1 to be first, got %s", result[0].GetName())
	}

	// Verify step2 comes after step1
	if result[1] != step2 {
		t.Errorf("Expected step2 to be second, got %s", result[1].GetName())
	}

	// Verify step3 is next since it has no dependencies
	if result[2] != step3 {
		t.Errorf("Expected step3 to be third, got %s", result[2].GetName())
	}

	// Verify step4 comes after step3
	if result[3] != step4 {
		t.Errorf("Expected step4 to be fourth, got %s", result[3].GetName())
	}
}

func Test_Func_BuildDependencyGraph_BasicChain(t *testing.T) {
	// Create test steps
	step1 := NewStep()
	step1.SetName("Step1")
	step2 := NewStep()
	step2.SetName("Step2")
	step3 := NewStep()
	step3.SetName("Step3")

	// Create runnables map
	runnables := map[string]RunnableInterface{
		"step1": step1,
		"step2": step2,
		"step3": step3,
	}

	// Create dependencies
	dependencies := map[string][]string{
		"step2": {"step1"},
		"step3": {"step2"},
	}

	conditionalDependencies := map[string]map[string]func(context.Context, map[string]any) bool{}

	ctx := context.Background()
	data := make(map[string]any)
	graph := buildDependencyGraph(runnables, dependencies, conditionalDependencies, ctx, data)

	// Verify graph structure
	if len(graph) != 3 {
		t.Errorf("Expected 3 nodes in graph, got %d", len(graph))
	}

	// Verify regular dependencies
	if len(graph[step1]) != 0 {
		t.Errorf("Expected step1 to have 0 dependencies, got %d", len(graph[step1]))
	}

	if len(graph[step2]) != 1 || graph[step2][0] != step1 {
		t.Errorf("Expected step2 to depend on step1")
	}

	if len(graph[step3]) != 1 || graph[step3][0] != step2 {
		t.Errorf("Expected step3 to depend on step2")
	}
}

func Test_Func_BuildDependencyGraph_ConditionalDependencies_NotMet(t *testing.T) {
	// Create test steps
	step1 := NewStep()
	step1.SetName("Step1")
	step2 := NewStep()
	step2.SetName("Step2")
	step3 := NewStep()
	step3.SetName("Step3")

	// Create runnables map
	runnables := map[string]RunnableInterface{
		"step1": step1,
		"step2": step2,
		"step3": step3,
	}

	// Create dependencies
	dependencies := map[string][]string{
		"step2": {"step1"},
	}

	// Create conditional dependencies
	conditionalDependencies := map[string]map[string]func(context.Context, map[string]any) bool{
		"step3": {
			"step2": func(ctx context.Context, data map[string]any) bool {
				return data["condition"] == true
			},
		},
	}

	ctx := context.Background()
	data := map[string]any{
		"condition": false,
	}
	graph := buildDependencyGraph(runnables, dependencies, conditionalDependencies, ctx, data)

	// Verify regular dependencies
	if len(graph[step2]) != 1 || graph[step2][0] != step1 {
		t.Errorf("Expected step2 to depend on step1")
	}

	// Verify conditional dependencies not added
	if len(graph[step3]) != 0 {
		t.Errorf("Expected step3 to have 0 dependencies when condition is false")
	}
}

func Test_Func_BuildDependencyGraph_ConditionalDependencies_Met(t *testing.T) {
	// Create test steps
	step1 := NewStep()
	step1.SetName("Step1")
	step2 := NewStep()
	step2.SetName("Step2")
	step3 := NewStep()
	step3.SetName("Step3")

	// Create runnables map
	runnables := map[string]RunnableInterface{
		"step1": step1,
		"step2": step2,
		"step3": step3,
	}

	// Create dependencies
	dependencies := map[string][]string{
		"step2": {"step1"},
	}

	// Create conditional dependencies
	conditionalDependencies := map[string]map[string]func(context.Context, map[string]any) bool{
		"step3": {
			"step2": func(ctx context.Context, data map[string]any) bool {
				return data["condition"] == true
			},
		},
	}

	ctx := context.Background()
	data := map[string]any{
		"condition": true,
	}
	graph := buildDependencyGraph(runnables, dependencies, conditionalDependencies, ctx, data)

	// Verify conditional dependencies added
	if len(graph[step3]) != 1 || graph[step3][0] != step2 {
		t.Errorf("Expected step3 to depend on step2 when condition is true")
	}
}

func Test_Func_BuildDependencyGraph_CircularDependencies(t *testing.T) {
	// Create test steps
	step1 := NewStep()
	step1.SetName("Step1")
	step2 := NewStep()
	step2.SetName("Step2")

	// Create runnables map
	runnables := map[string]RunnableInterface{
		"step1": step1,
		"step2": step2,
	}

	// Create circular dependencies
	dependencies := map[string][]string{
		"step1": {"step2"},
		"step2": {"step1"},
	}

	conditionalDependencies := map[string]map[string]func(context.Context, map[string]any) bool{}

	ctx := context.Background()
	data := make(map[string]any)
	graph := buildDependencyGraph(runnables, dependencies, conditionalDependencies, ctx, data)

	// Verify graph structure
	if len(graph) != 2 {
		t.Errorf("Expected 2 nodes in graph, got %d", len(graph))
	}

	// Verify circular dependencies
	if len(graph[step1]) != 1 || graph[step1][0] != step2 {
		t.Errorf("Expected step1 to depend on step2")
	}

	if len(graph[step2]) != 1 || graph[step2][0] != step1 {
		t.Errorf("Expected step2 to depend on step1")
	}
}
