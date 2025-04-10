package wf_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	// Import the package being tested
	"github.com/dracory/base/wf" // Adjust this path if your module path is different
)

func TestPipelineVisualization_Empty(t *testing.T) {
	// Test empty pipeline
	emptyPipeline := wf.NewPipeline()
	dot := emptyPipeline.Visualize()
	if !strings.Contains(dot, "digraph") {
		t.Error("Empty pipeline visualization should contain 'digraph'")
	}
	if !strings.Contains(dot, "rankdir = \"LR\"") {
		t.Error("Empty pipeline visualization should have left-to-right layout")
	}
	// Check that there are no nodes or edges defined beyond the basic structure
	if strings.Contains(dot, "[label=") || strings.Contains(dot, "->") {
		t.Error("Empty pipeline visualization should not contain nodes or edges")
	}
}

func TestPipelineVisualization(t *testing.T) {
	// Create a pipeline with three steps
	pipeline := wf.NewPipeline()
	step1 := wf.NewStep()
	step2 := wf.NewStep()
	step3 := wf.NewStep()

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

	// Test initial pipeline visualization (should be default colors)
	dot := pipeline.Visualize()

	// Basic checks
	if !strings.Contains(dot, "digraph") {
		t.Error("Pipeline visualization should contain 'digraph'")
	}
	if !strings.Contains(dot, "rankdir = \"LR\"") {
		t.Error("Pipeline visualization should have left-to-right layout")
	}

	// Check if all steps are present by label
	if !strings.Contains(dot, `label="Step 1"`) || !strings.Contains(dot, `label="Step 2"`) || !strings.Contains(dot, `label="Step 3"`) {
		t.Error("Pipeline visualization should contain all step names as labels")
	}
	// Check if all steps are present by ID (node definition)
	if !strings.Contains(dot, `"`+step1.GetID()+`"`) || !strings.Contains(dot, `"`+step2.GetID()+`"`) || !strings.Contains(dot, `"`+step3.GetID()+`"`) {
		t.Error("Pipeline visualization should contain all step IDs as nodes")
	}

	// Check if edges are present
	edge12 := fmt.Sprintf(`"%s" -> "%s"`, step1.GetID(), step2.GetID())
	edge23 := fmt.Sprintf(`"%s" -> "%s"`, step2.GetID(), step3.GetID())
	if !strings.Contains(dot, edge12) {
		t.Errorf("Pipeline visualization should contain edge: %s", edge12)
	}
	if !strings.Contains(dot, edge23) {
		t.Errorf("Pipeline visualization should contain edge: %s", edge23)
	}

	// --- Test visualization with current step (Running) ---
	// Create a new state for this test case to ensure isolation
	runningState := wf.NewState()
	runningState.SetStatus(wf.StateStatusRunning)
	runningState.SetCurrentStepID(step2.GetID())
	// Add completed steps if necessary for the running state visualization logic
	// runningState.AddCompletedStep(step1.GetID()) // Example if needed
	pipeline.SetState(runningState) // Apply the specific state

	dot = pipeline.Visualize()
	step2NodeDef := fmt.Sprintf(`"%s" [label="Step 2" shape=box style=filled tooltip="Step: Step 2" fillcolor="#2196F3" fontcolor="white"]`, step2.GetID())
	if !strings.Contains(dot, step2NodeDef) {
		t.Errorf("Current step (Step 2) should be colored blue. Expected substring: %s\nGot DOT:\n%s", step2NodeDef, dot)
	}
	// Ensure others are default
	step1NodeDefDefault := fmt.Sprintf(`"%s" [label="Step 1" shape=box style=solid tooltip="Step: Step 1" fillcolor="#ffffff" ]`, step1.GetID())
	if !strings.Contains(dot, step1NodeDefDefault) {
		t.Errorf("Non-current step (Step 1) should have default style. Expected substring: %s\nGot DOT:\n%s", step1NodeDefDefault, dot)
	}
	// Check edge colors for running state
	edge12Running := fmt.Sprintf(`"%s" -> "%s" [style=solid tooltip="From Step 1 to Step 2" color="#4CAF50"]`, step1.GetID(), step2.GetID()) // Edge before current might be green
	edge23Running := fmt.Sprintf(`"%s" -> "%s" [style=solid tooltip="From Step 2 to Step 3" color="#9E9E9E"]`, step2.GetID(), step3.GetID()) // Edge after current should be grey
	if !strings.Contains(dot, edge12Running) {
		t.Errorf("Edge before current step (1->2) should be green. Expected substring: %s\nGot DOT:\n%s", edge12Running, dot)
	}
	if !strings.Contains(dot, edge23Running) {
		t.Errorf("Edge after current step (2->3) should be grey. Expected substring: %s\nGot DOT:\n%s", edge23Running, dot)
	}

	// --- Test visualization with completed steps ---
	// Create a new state for this test case
	completedState := wf.NewState()
	completedState.SetStatus(wf.StateStatusComplete)
	completedState.SetCurrentStepID("") // No current step when complete
	completedState.AddCompletedStep(step1.GetID())
	completedState.AddCompletedStep(step2.GetID())
	completedState.AddCompletedStep(step3.GetID())
	pipeline.SetState(completedState) // Apply the specific state

	dot = pipeline.Visualize()
	step1NodeDefComplete := fmt.Sprintf(`"%s" [label="Step 1" shape=box style=filled tooltip="Step: Step 1" fillcolor="#4CAF50" fontcolor="white"]`, step1.GetID())
	step2NodeDefComplete := fmt.Sprintf(`"%s" [label="Step 2" shape=box style=filled tooltip="Step: Step 2" fillcolor="#4CAF50" fontcolor="white"]`, step2.GetID())
	// Based on visualization.go: `i < len(p.nodes)-1`, the last node won't be green on complete.
	step3NodeDefComplete := fmt.Sprintf(`"%s" [label="Step 3" shape=box style=solid tooltip="Step: Step 3" fillcolor="#ffffff" ]`, step3.GetID())

	if !strings.Contains(dot, step1NodeDefComplete) {
		t.Errorf("Completed step (Step 1) should be colored green. Expected substring: %s\nGot DOT:\n%s", step1NodeDefComplete, dot)
	}
	if !strings.Contains(dot, step2NodeDefComplete) {
		t.Errorf("Completed step (Step 2) should be colored green. Expected substring: %s\nGot DOT:\n%s", step2NodeDefComplete, dot)
	}
	if !strings.Contains(dot, step3NodeDefComplete) {
		t.Errorf("Last step (Step 3) should not be green when pipeline complete. Expected substring: %s\nGot DOT:\n%s", step3NodeDefComplete, dot)
	}
	// Check edge colors on complete
	edge12Complete := fmt.Sprintf(`"%s" -> "%s" [style=solid tooltip="From Step 1 to Step 2" color="#4CAF50"]`, step1.GetID(), step2.GetID())
	edge23Complete := fmt.Sprintf(`"%s" -> "%s" [style=solid tooltip="From Step 2 to Step 3" color="#4CAF50"]`, step2.GetID(), step3.GetID())
	if !strings.Contains(dot, edge12Complete) {
		t.Errorf("Completed edge (1->2) should be colored green. Expected substring: %s\nGot DOT:\n%s", edge12Complete, dot)
	}
	if !strings.Contains(dot, edge23Complete) {
		t.Errorf("Completed edge (2->3) should be colored green. Expected substring: %s\nGot DOT:\n%s", edge23Complete, dot)
	}

	// --- Test visualization with failed step ---
	// Create a new state for this test case
	failedState := wf.NewState()
	// Need a valid prior state to transition from, e.g., Running
	failedState.SetStatus(wf.StateStatusRunning)
	failedState.SetStatus(wf.StateStatusFailed) // Now transition to Failed
	failedState.SetCurrentStepID(step2.GetID()) // Failed at step 2
	failedState.AddCompletedStep(step1.GetID()) // Only step 1 completed
	pipeline.SetState(failedState)              // Apply the specific state

	dot = pipeline.Visualize()
	step2NodeDefFailed := fmt.Sprintf(`"%s" [label="Step 2" shape=box style=filled tooltip="Step: Step 2" fillcolor="#F44336" fontcolor="white"]`, step2.GetID())
	if !strings.Contains(dot, step2NodeDefFailed) {
		t.Errorf("Failed step (Step 2) should be colored red. Expected substring: %s\nGot DOT:\n%s", step2NodeDefFailed, dot)
	}
	// Check edge colors on fail (should be default grey)
	edge12Failed := fmt.Sprintf(`"%s" -> "%s" [style=solid tooltip="From Step 1 to Step 2" color="#9E9E9E"]`, step1.GetID(), step2.GetID())
	edge23Failed := fmt.Sprintf(`"%s" -> "%s" [style=solid tooltip="From Step 2 to Step 3" color="#9E9E9E"]`, step2.GetID(), step3.GetID())
	if !strings.Contains(dot, edge12Failed) {
		t.Errorf("Edge leading to failed step (1->2) should be default grey. Expected substring: %s\nGot DOT:\n%s", edge12Failed, dot)
	}
	if !strings.Contains(dot, edge23Failed) {
		t.Errorf("Edge after failed step (2->3) should be default grey. Expected substring: %s\nGot DOT:\n%s", edge23Failed, dot)
	}

	// --- Test visualization with paused step ---
	// Create a new state for this test case
	pausedState := wf.NewState()
	// Need a valid prior state to transition from, e.g., Running
	pausedState.SetStatus(wf.StateStatusRunning)
	pausedState.SetStatus(wf.StateStatusPaused) // Now transition to Paused
	pausedState.SetCurrentStepID(step2.GetID()) // Paused at step 2
	pausedState.AddCompletedStep(step1.GetID()) // Only step 1 completed
	pipeline.SetState(pausedState)              // Apply the specific state

	dot = pipeline.Visualize()
	step2NodeDefPaused := fmt.Sprintf(`"%s" [label="Step 2" shape=box style=filled tooltip="Step: Step 2" fillcolor="#FFC107" fontcolor="white"]`, step2.GetID())
	if !strings.Contains(dot, step2NodeDefPaused) {
		t.Errorf("Paused step (Step 2) should be colored yellow. Expected substring: %s\nGot DOT:\n%s", step2NodeDefPaused, dot)
	}
	// Check edge colors on pause (should be default grey)
	edge12Paused := fmt.Sprintf(`"%s" -> "%s" [style=solid tooltip="From Step 1 to Step 2" color="#9E9E9E"]`, step1.GetID(), step2.GetID())
	edge23Paused := fmt.Sprintf(`"%s" -> "%s" [style=solid tooltip="From Step 2 to Step 3" color="#9E9E9E"]`, step2.GetID(), step3.GetID())
	if !strings.Contains(dot, edge12Paused) {
		t.Errorf("Edge leading to paused step (1->2) should be default grey. Expected substring: %s\nGot DOT:\n%s", edge12Paused, dot)
	}
	if !strings.Contains(dot, edge23Paused) {
		t.Errorf("Edge after paused step (2->3) should be default grey. Expected substring: %s\nGot DOT:\n%s", edge23Paused, dot)
	}
}

// ... (rest of the file remains the same) ...

func TestDagVisualization_Empty(t *testing.T) {
	// Test empty DAG
	emptyDag := wf.NewDag()
	dot := emptyDag.Visualize()
	if !strings.Contains(dot, "digraph") {
		t.Error("Empty DAG visualization should contain 'digraph'")
	}
	if !strings.Contains(dot, "rankdir = \"LR\"") {
		t.Error("Empty DAG visualization should have left-to-right layout")
	}
	// Check that there are no nodes or edges defined beyond the basic structure
	if strings.Contains(dot, "[label=") || strings.Contains(dot, "->") {
		t.Error("Empty DAG visualization should not contain nodes or edges")
	}
}

// Helper function to create a standard DAG for testing visualization states
func createTestDag() (wf.DagInterface, wf.StepInterface, wf.StepInterface, wf.StepInterface, wf.StepInterface, wf.StepInterface) {
	dag := wf.NewDag()
	step1 := wf.NewStep()
	step2 := wf.NewStep()
	step3 := wf.NewStep()
	step4 := wf.NewStep()
	step5 := wf.NewStep()

	step1.SetName("Step 1")
	step2.SetName("Step 2")
	step3.SetName("Step 3")
	step4.SetName("Step 4")
	step5.SetName("Step 5")

	// Set up handlers for each step
	handler := func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		return ctx, data, nil
	}
	step1.SetHandler(handler)
	step2.SetHandler(handler)
	step3.SetHandler(handler)
	step4.SetHandler(handler)
	step5.SetHandler(handler)

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

	return dag, step1, step2, step3, step4, step5
}

func TestDagVisualization_Basic(t *testing.T) {
	dag, step1, step2, step3, step4, step5 := createTestDag()

	// Test DAG visualization (initial state)
	dot := dag.Visualize()

	// Basic checks
	if !strings.Contains(dot, "digraph") {
		t.Error("DAG visualization should contain 'digraph'")
	}
	if !strings.Contains(dot, "rankdir = \"LR\"") {
		t.Error("DAG visualization should have left-to-right layout")
	}

	// Check if all steps are present by label
	if !strings.Contains(dot, `label="Step 1"`) || !strings.Contains(dot, `label="Step 2"`) ||
		!strings.Contains(dot, `label="Step 3"`) || !strings.Contains(dot, `label="Step 4"`) ||
		!strings.Contains(dot, `label="Step 5"`) {
		t.Error("DAG visualization should contain all step names as labels")
	}
	// Check if all steps are present by ID (node definition)
	if !strings.Contains(dot, `"`+step1.GetID()+`"`) || !strings.Contains(dot, `"`+step2.GetID()+`"`) ||
		!strings.Contains(dot, `"`+step3.GetID()+`"`) || !strings.Contains(dot, `"`+step4.GetID()+`"`) ||
		!strings.Contains(dot, `"`+step5.GetID()+`"`) {
		t.Error("DAG visualization should contain all step IDs as nodes")
	}

	// Check if dependencies are represented as edges
	if !strings.Contains(dot, "->") {
		t.Error("DAG visualization should contain edges for dependencies")
	}

	// Verify specific dependencies are present in the visualization
	dependencies := []struct{ from, to wf.StepInterface }{
		{step1, step2},
		{step1, step3},
		{step2, step4},
		{step3, step4},
		{step2, step5},
		{step3, step5},
	}

	for _, dep := range dependencies {
		edge := fmt.Sprintf(`"%s" -> "%s"`, dep.from.GetID(), dep.to.GetID())
		tooltip := fmt.Sprintf(`tooltip="From %s to %s"`, dep.from.GetName(), dep.to.GetName())
		color := `color="#9E9E9E"` // Default edge color
		if !strings.Contains(dot, edge) {
			t.Errorf("DAG visualization should contain dependency edge: %s", edge)
		}
		if !strings.Contains(dot, tooltip) {
			t.Errorf("DAG visualization edge should contain tooltip: %s", tooltip)
		}
		if !strings.Contains(dot, color) {
			t.Errorf("DAG visualization edge should contain default color: %s", color)
		}
	}
}

func TestDagVisualization_Running(t *testing.T) {
	dag, step1, step2, step3, step4, _ := createTestDag()

	// Simulate running state: Step 3 is current, Step 1 is done.
	dag.GetState().SetStatus(wf.StateStatusRunning)
	dag.GetState().AddCompletedStep(step1.GetID())
	dag.GetState().SetCurrentStepID(step3.GetID()) // Step 3 is now the current one

	dot := dag.Visualize()
	t.Logf("Running DAG Viz:\n%s", dot)

	// Check Step 1 (Completed)
	step1NodeDef := fmt.Sprintf(`"%s" [label="Step 1" shape=box style=filled tooltip="Step: Step 1" fillcolor="#4CAF50" fontcolor="white"]`, step1.GetID())
	if !strings.Contains(dot, step1NodeDef) {
		t.Errorf("Completed step (Step 1) should be green. Expected substring: %s", step1NodeDef)
	}

	// Check Step 3 (Current/Running)
	step3NodeDef := fmt.Sprintf(`"%s" [label="Step 3" shape=box style=filled tooltip="Step: Step 3" fillcolor="#2196F3" fontcolor="white"]`, step3.GetID())
	if !strings.Contains(dot, step3NodeDef) {
		t.Errorf("Current running step (Step 3) should be blue. Expected substring: %s", step3NodeDef)
	}

	// Check Step 2 (Waiting, depends on completed Step 1) - Should be default
	step2NodeDef := fmt.Sprintf(`"%s" [label="Step 2" shape=box style=solid tooltip="Step: Step 2" fillcolor="#ffffff" ]`, step2.GetID())
	if !strings.Contains(dot, step2NodeDef) {
		t.Errorf("Waiting step (Step 2) should be default. Expected substring: %s", step2NodeDef)
	}

	// Check Step 4 (Waiting) - Should be default
	step4NodeDef := fmt.Sprintf(`"%s" [label="Step 4" shape=box style=solid tooltip="Step: Step 4" fillcolor="#ffffff" ]`, step4.GetID())
	if !strings.Contains(dot, step4NodeDef) {
		t.Errorf("Waiting step (Step 4) should be default. Expected substring: %s", step4NodeDef)
	}

	// Check Edges
	edge12 := fmt.Sprintf(`"%s" -> "%s" [style=solid tooltip="From Step 1 to Step 2" color="#4CAF50"]`, step1.GetID(), step2.GetID()) // From completed = green
	edge13 := fmt.Sprintf(`"%s" -> "%s" [style=solid tooltip="From Step 1 to Step 3" color="#4CAF50"]`, step1.GetID(), step3.GetID()) // From completed = green
	edge24 := fmt.Sprintf(`"%s" -> "%s" [style=solid tooltip="From Step 2 to Step 4" color="#9E9E9E"]`, step2.GetID(), step4.GetID()) // Default
	edge34 := fmt.Sprintf(`"%s" -> "%s" [style=solid tooltip="From Step 3 to Step 4" color="#9E9E9E"]`, step3.GetID(), step4.GetID()) // Default

	if !strings.Contains(dot, edge12) {
		t.Errorf("Edge from completed step (1->2) should be green. Expected substring: %s", edge12)
	}
	if !strings.Contains(dot, edge13) {
		t.Errorf("Edge from completed step (1->3) should be green. Expected substring: %s", edge13)
	}
	if !strings.Contains(dot, edge24) {
		t.Errorf("Edge from waiting step (2->4) should be default grey. Expected substring: %s", edge24)
	}
	if !strings.Contains(dot, edge34) {
		t.Errorf("Edge from current step (3->4) should be default grey. Expected substring: %s", edge34)
	}
}

func TestDagVisualization_Completed(t *testing.T) {
	dag, step1, step2, step3, step4, step5 := createTestDag()

	// Simulate completed state
	dag.GetState().SetStatus(wf.StateStatusComplete)
	dag.GetState().AddCompletedStep(step1.GetID())
	dag.GetState().AddCompletedStep(step2.GetID())
	dag.GetState().AddCompletedStep(step3.GetID())
	dag.GetState().AddCompletedStep(step4.GetID())
	dag.GetState().AddCompletedStep(step5.GetID())
	dag.GetState().SetCurrentStepID("") // No current step

	dot := dag.Visualize()
	t.Logf("Completed DAG Viz:\n%s", dot)

	// Check Nodes (All should be default color when DAG is complete, based on visualization.go logic)
	expectedNodeStyle := `style=solid tooltip="Step: Step %d" fillcolor="#ffffff"`
	for i, step := range []wf.StepInterface{step1, step2, step3, step4, step5} {
		nodeDef := fmt.Sprintf(`"%s" [label="Step %d" shape=box %s ]`, step.GetID(), i+1, fmt.Sprintf(expectedNodeStyle, i+1))
		if !strings.Contains(dot, nodeDef) {
			t.Errorf("Node (Step %d) should have default style when DAG complete. Expected substring: %s", i+1, nodeDef)
		}
	}

	// Check Edges (All should be green when DAG is complete)
	dependencies := []struct{ from, to wf.StepInterface }{
		{step1, step2}, {step1, step3}, {step2, step4}, {step3, step4}, {step2, step5}, {step3, step5},
	}
	for _, dep := range dependencies {
		edge := fmt.Sprintf(`"%s" -> "%s" [style=solid tooltip="From %s to %s" color="#4CAF50"]`, dep.from.GetID(), dep.to.GetID(), dep.from.GetName(), dep.to.GetName())
		if !strings.Contains(dot, edge) {
			t.Errorf("Edge (%s->%s) should be green when DAG complete. Expected substring: %s", dep.from.GetName(), dep.to.GetName(), edge)
		}
	}
}

func TestDagVisualization_Failed(t *testing.T) {
	dag, step1, step2, step3, step4, _ := createTestDag()

	// Simulate failed state: Failed at Step 4. Steps 1, 2, 3 completed.
	dag.GetState().SetStatus(wf.StateStatusFailed)
	dag.GetState().AddCompletedStep(step1.GetID())
	dag.GetState().AddCompletedStep(step2.GetID())
	dag.GetState().AddCompletedStep(step3.GetID())
	dag.GetState().SetCurrentStepID(step4.GetID()) // Failed step

	dot := dag.Visualize()
	t.Logf("Failed DAG Viz:\n%s", dot)

	// Check Step 4 (Failed)
	step4NodeDef := fmt.Sprintf(`"%s" [label="Step 4" shape=box style=filled tooltip="Step: Step 4" fillcolor="#F44336" fontcolor="white"]`, step4.GetID())
	if !strings.Contains(dot, step4NodeDef) {
		t.Errorf("Failed step (Step 4) should be red. Expected substring: %s", step4NodeDef)
	}

	// Check Completed Steps (Should be default color when DAG failed, based on visualization.go)
	expectedNodeStyle := `style=solid tooltip="Step: Step %d" fillcolor="#ffffff"`
	for i, step := range []wf.StepInterface{step1, step2, step3} {
		nodeDef := fmt.Sprintf(`"%s" [label="Step %d" shape=box %s ]`, step.GetID(), i+1, fmt.Sprintf(expectedNodeStyle, i+1))
		if !strings.Contains(dot, nodeDef) {
			t.Errorf("Completed node (Step %d) should have default style when DAG failed. Expected substring: %s", i+1, nodeDef)
		}
	}

	// Check Edges (Should be default grey when DAG failed)
	dependencies := []struct{ from, to wf.StepInterface }{
		{step1, step2}, {step1, step3}, {step2, step4}, {step3, step4},
	}
	for _, dep := range dependencies {
		edge := fmt.Sprintf(`"%s" -> "%s" [style=solid tooltip="From %s to %s" color="#9E9E9E"]`, dep.from.GetID(), dep.to.GetID(), dep.from.GetName(), dep.to.GetName())
		if !strings.Contains(dot, edge) {
			t.Errorf("Edge (%s->%s) should be default grey when DAG failed. Expected substring: %s", dep.from.GetName(), dep.to.GetName(), edge)
		}
	}
}

func TestDagVisualization_Paused(t *testing.T) {
	dag, step1, step2, step3, _, _ := createTestDag()

	// Simulate paused state: Paused at Step 3. Step 1 completed.
	dag.GetState().SetStatus(wf.StateStatusPaused)
	dag.GetState().AddCompletedStep(step1.GetID())
	dag.GetState().SetCurrentStepID(step3.GetID()) // Paused step

	dot := dag.Visualize()
	t.Logf("Paused DAG Viz:\n%s", dot)

	// Check Step 3 (Paused)
	step3NodeDef := fmt.Sprintf(`"%s" [label="Step 3" shape=box style=filled tooltip="Step: Step 3" fillcolor="#FFC107" fontcolor="white"]`, step3.GetID())
	if !strings.Contains(dot, step3NodeDef) {
		t.Errorf("Paused step (Step 3) should be yellow. Expected substring: %s", step3NodeDef)
	}

	// Check Completed Step 1 (Should be default color when DAG paused)
	step1NodeDef := fmt.Sprintf(`"%s" [label="Step 1" shape=box style=solid tooltip="Step: Step 1" fillcolor="#ffffff" ]`, step1.GetID())
	if !strings.Contains(dot, step1NodeDef) {
		t.Errorf("Completed node (Step 1) should have default style when DAG paused. Expected substring: %s", step1NodeDef)
	}

	// Check Waiting Step 2 (Should be default color)
	step2NodeDef := fmt.Sprintf(`"%s" [label="Step 2" shape=box style=solid tooltip="Step: Step 2" fillcolor="#ffffff" ]`, step2.GetID())
	if !strings.Contains(dot, step2NodeDef) {
		t.Errorf("Waiting node (Step 2) should have default style when DAG paused. Expected substring: %s", step2NodeDef)
	}

	// Check Edges (Should be default grey when DAG paused)
	dependencies := []struct{ from, to wf.StepInterface }{
		{step1, step2}, {step1, step3},
	}
	for _, dep := range dependencies {
		edge := fmt.Sprintf(`"%s" -> "%s" [style=solid tooltip="From %s to %s" color="#9E9E9E"]`, dep.from.GetID(), dep.to.GetID(), dep.from.GetName(), dep.to.GetName())
		if !strings.Contains(dot, edge) {
			t.Errorf("Edge (%s->%s) should be default grey when DAG paused. Expected substring: %s", dep.from.GetName(), dep.to.GetName(), edge)
		}
	}
}

func TestStepVisualization(t *testing.T) {
	// Create a step
	step := wf.NewStep()
	step.SetName("My Step")
	step.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		return ctx, data, nil
	})

	// Test step visualization (initial state - should be default white/solid)
	// Note: A new step's initial state might be 'Running' based on NewState(), adjust if needed.
	// Let's explicitly set it to an "empty" status for the initial check if possible,
	// or check against the actual initial status (likely Running or empty string).
	// Based on NewState(), it starts as Running. Let's test that first.
	initialState := wf.NewState() // Get a fresh state
	step.SetState(initialState)   // Set it
	dot := step.Visualize()
	t.Logf("Initial Step Viz (Expecting Running):\n%s", dot)

	// Check initial state (should be Running/Blue according to NewState and step Visualize logic)
	stepNodeDefInitial := fmt.Sprintf(`"%s" [label="My Step" shape=box style=filled tooltip="Step: My Step" fillcolor="#2196F3" fontcolor="white"]`, step.GetID())
	if !strings.Contains(dot, stepNodeDefInitial) {
		t.Errorf("Initial step state should be Running (blue). Expected substring: %s\nGot DOT:\n%s", stepNodeDefInitial, dot)
	}

	// Basic checks
	if !strings.Contains(dot, "digraph") {
		t.Error("Step visualization should contain 'digraph'")
	}
	if !strings.Contains(dot, "rankdir = \"LR\"") {
		t.Error("Step visualization should have left-to-right layout")
	}

	// Test visualization with different states explicitly set

	// Running (already tested as initial, but good to be explicit)
	runningState := wf.NewState()
	runningState.SetStatus(wf.StateStatusRunning)
	step.SetState(runningState)
	dot = step.Visualize()
	stepNodeDefRunning := fmt.Sprintf(`"%s" [label="My Step" shape=box style=filled tooltip="Step: My Step" fillcolor="#2196F3" fontcolor="white"]`, step.GetID())
	if !strings.Contains(dot, stepNodeDefRunning) {
		t.Errorf("Running step should be colored blue. Expected substring: %s\nGot DOT:\n%s", stepNodeDefRunning, dot)
	}

	// Completed
	completedState := wf.NewState()
	completedState.SetStatus(wf.StateStatusRunning) // Valid transition from Running
	completedState.SetStatus(wf.StateStatusComplete)
	step.SetState(completedState)
	dot = step.Visualize()
	stepNodeDefComplete := fmt.Sprintf(`"%s" [label="My Step" shape=box style=filled tooltip="Step: My Step" fillcolor="#4CAF50" fontcolor="white"]`, step.GetID())
	if !strings.Contains(dot, stepNodeDefComplete) {
		t.Errorf("Completed step should be colored green. Expected substring: %s\nGot DOT:\n%s", stepNodeDefComplete, dot)
	}

	// Failed
	failedState := wf.NewState()
	failedState.SetStatus(wf.StateStatusRunning) // Valid transition from Running
	failedState.SetStatus(wf.StateStatusFailed)
	step.SetState(failedState)
	dot = step.Visualize()
	stepNodeDefFailed := fmt.Sprintf(`"%s" [label="My Step" shape=box style=filled tooltip="Step: My Step" fillcolor="#F44336" fontcolor="white"]`, step.GetID())
	if !strings.Contains(dot, stepNodeDefFailed) {
		t.Errorf("Failed step should be colored red. Expected substring: %s\nGot DOT:\n%s", stepNodeDefFailed, dot)
	}

	// Paused
	pausedState := wf.NewState()
	pausedState.SetStatus(wf.StateStatusRunning) // Valid transition from Running
	pausedState.SetStatus(wf.StateStatusPaused)
	step.SetState(pausedState)
	dot = step.Visualize()
	stepNodeDefPaused := fmt.Sprintf(`"%s" [label="My Step" shape=box style=filled tooltip="Step: My Step" fillcolor="#FFC107" fontcolor="white"]`, step.GetID())
	if !strings.Contains(dot, stepNodeDefPaused) {
		t.Errorf("Paused step should be colored yellow. Expected substring: %s\nGot DOT:\n%s", stepNodeDefPaused, dot)
	}
}
