package wf

import (
	"fmt"
	"slices" // Keep slices import as it's used in other Visualize methods
	"strings"
	// Keep bytes import as it's used in other Visualize methods
)

// DotNodeSpec represents a node in the DOT graph
type DotNodeSpec struct {
	Name        string
	DisplayName string
	Tooltip     string
	Shape       string
	Style       string
	FillColor   string
	// Add FontColor for easier string building if needed, or handle conditionally
}

// DotEdgeSpec represents an edge in the DOT graph
type DotEdgeSpec struct {
	FromNodeName string
	ToNodeName   string
	Tooltip      string
	Style        string
	Color        string
}

// Helper function to escape double quotes in DOT labels/tooltips
func escapeDotString(s string) string {
	return strings.ReplaceAll(s, `"`, `\"`)
}

// getNodeStyleAndColor determines the node's style and color based on its state,
// *only if* it is the current step. Otherwise, returns defaults.
func getNodeStyleAndColor(state StateInterface, isCurrentStep bool) (string, string) {
	nodeStyle := "solid"
	fillColor := "#ffffff" // Default white

	// Only apply special styling if this node is the current step
	if state != nil && isCurrentStep {
		status := state.GetStatus()
		// Check if the current node should be filled based on the workflow status
		shouldBeFilled := status == StateStatusRunning || status == StateStatusComplete || status == StateStatusFailed || status == StateStatusPaused

		if shouldBeFilled {
			nodeStyle = "filled"
			switch status {
			case StateStatusFailed:
				fillColor = "#F44336" // Red
			case StateStatusPaused:
				fillColor = "#FFC107" // Yellow
			case StateStatusComplete:
				// If the workflow is complete, the "current" step (which is likely the last one)
				// might need specific styling, but often it's just default.
				// Let the specific Visualize method handle this override if needed.
				// For now, let's assume the last step in a completed workflow is default.
				// Revert to default if the status is Complete.
				nodeStyle = "solid"
				fillColor = "#ffffff"
			case StateStatusRunning:
				fillColor = "#2196F3" // Blue
			default:
				// Fallback for unexpected status when filled
				fillColor = "#2196F3" // Default blue
			}
		}
		// If status is not one that implies filling (e.g., initial state ""),
		// it remains solid white.
	}
	return nodeStyle, fillColor
}

// dotTemplateFuncs generates the DOT graph string directly from node and edge specs.
func dotTemplateFuncs(nodes []*DotNodeSpec, edges []*DotEdgeSpec) string {
	var sb strings.Builder

	// Write graph header
	sb.WriteString("digraph {\n")
	sb.WriteString("\trankdir = \"LR\"\n")
	sb.WriteString("\tnode [fontname=\"Arial\"]\n")
	sb.WriteString("\tedge [fontname=\"Arial\"]\n")

	// Write node definitions
	for _, node := range nodes {
		sb.WriteString(fmt.Sprintf("\t\"%s\" [label=\"%s\" shape=%s style=%s tooltip=\"%s\" fillcolor=\"%s\"",
			escapeDotString(node.Name),
			escapeDotString(node.DisplayName),
			node.Shape,
			node.Style,
			escapeDotString(node.Tooltip),
			node.FillColor,
		))
		// Add fontcolor conditionally for filled nodes
		if node.Style == "filled" {
			sb.WriteString(" fontcolor=\"white\"")
		}
		sb.WriteString("]\n")
	}

	// Write edge definitions
	for _, edge := range edges {
		sb.WriteString(fmt.Sprintf("\t\"%s\" -> \"%s\" [style=%s tooltip=\"%s\" color=\"%s\"]\n",
			escapeDotString(edge.FromNodeName),
			escapeDotString(edge.ToNodeName),
			edge.Style,
			escapeDotString(edge.Tooltip),
			edge.Color,
		))
	}

	// Write graph footer
	sb.WriteString("}\n")

	return sb.String()
}

// Visualize returns a DOT graph representation of the workflow using the helper function
func (p *pipelineImplementation) Visualize() string {
	// Handle empty pipeline
	if len(p.nodes) == 0 {
		// Return the basic empty graph structure
		return dotTemplateFuncs([]*DotNodeSpec{}, []*DotEdgeSpec{})
	}

	nodes := make([]*DotNodeSpec, 0, len(p.nodes))
	edges := make([]*DotEdgeSpec, 0, len(p.nodes)-1)

	status := p.state.GetStatus()
	currentStepID := p.state.GetCurrentStepID()

	// Create nodes
	for i, node := range p.nodes {
		isCurrentStep := currentStepID == node.GetID()
		nodeStyle, fillColor := getNodeStyleAndColor(p.state, isCurrentStep) // Get style only if current

		// --- Logic for NON-CURRENT steps ---
		if !isCurrentStep {
			// Pipeline specific logic: Completed steps (except the last one) are green *if* pipeline is complete
			if status == StateStatusComplete && i < len(p.nodes)-1 {
				nodeStyle = "filled"
				fillColor = "#4CAF50" // Green
			} else {
				// Otherwise, non-current steps are default
				nodeStyle = "solid"
				fillColor = "#ffffff" // Default white
			}
		}
		// --- End Logic for NON-CURRENT steps ---

		// If it *is* the current step, getNodeStyleAndColor already determined the style/color

		nodes = append(nodes, &DotNodeSpec{
			Name:        node.GetID(),
			DisplayName: node.GetName(),
			Tooltip:     fmt.Sprintf("Step: %s", node.GetName()),
			Shape:       "box",
			Style:       nodeStyle,
			FillColor:   fillColor,
		})

		// Create edges between steps
		if i > 0 {
			edgeStyle := "solid"
			edgeColor := "#9E9E9E" // Default grey

			// Find the index of the current step if it exists
			currentStepIndex := -1
			if currentStepID != "" {
				for idx, n := range p.nodes {
					if n.GetID() == currentStepID {
						currentStepIndex = idx
						break
					}
				}
			}

			// Color edges green if pipeline is complete OR if the source step is before/at the current running step
			if status == StateStatusComplete ||
				(status == StateStatusRunning && currentStepIndex != -1 && i <= currentStepIndex) {
				edgeColor = "#4CAF50" // Green
			}

			edges = append(edges, &DotEdgeSpec{
				FromNodeName: p.nodes[i-1].GetID(),
				ToNodeName:   node.GetID(),
				Style:        edgeStyle,
				Color:        edgeColor,
				Tooltip:      fmt.Sprintf("From %s to %s", p.nodes[i-1].GetName(), node.GetName()),
			})
		}
	}

	// Use the helper function to generate the DOT string
	return dotTemplateFuncs(nodes, edges)
}

// Visualize returns a DOT graph representation of the DAG using the helper function
func (d *Dag) Visualize() string {
	// Handle empty DAG
	if len(d.runnables) == 0 {
		// Return the basic empty graph structure
		return dotTemplateFuncs([]*DotNodeSpec{}, []*DotEdgeSpec{})
	}

	nodes := make([]*DotNodeSpec, 0, len(d.runnables))
	edges := make([]*DotEdgeSpec, 0)

	status := d.state.GetStatus()
	currentStepID := d.state.GetCurrentStepID()
	completedSteps := d.state.GetCompletedSteps()

	// Create nodes
	for _, node := range d.runnables {
		isCurrentStep := currentStepID == node.GetID()
		nodeStyle, fillColor := getNodeStyleAndColor(d.state, isCurrentStep) // Get style only if current

		// --- Logic for NON-CURRENT steps ---
		if !isCurrentStep {
			// DAG specific logic: Completed steps are green *only* when the DAG is *running*
			if status == StateStatusRunning && slices.Contains(completedSteps, node.GetID()) {
				nodeStyle = "filled"
				fillColor = "#4CAF50" // green
			} else {
				// Otherwise, non-current steps are default (including completed steps when DAG is not running)
				nodeStyle = "solid"
				fillColor = "#ffffff" // Default white
			}
		}
		// --- End Logic for NON-CURRENT steps ---

		// If it *is* the current step, getNodeStyleAndColor already determined the style/color

		nodes = append(nodes, &DotNodeSpec{
			Name:        node.GetID(),
			DisplayName: node.GetName(),
			Tooltip:     fmt.Sprintf("Step: %s", node.GetName()),
			Shape:       "box",
			Style:       nodeStyle,
			FillColor:   fillColor,
		})
	}

	// Create edges based on dependencies
	for dependentID, dependencyIDs := range d.dependencies {
		for _, dependencyID := range dependencyIDs {
			dependent, depExists := d.runnables[dependentID]
			dependency, depExists2 := d.runnables[dependencyID]

			if !depExists || !depExists2 {
				continue // Should not happen if DAG is consistent
			}

			edgeStyle := "solid"
			edgeColor := "#9E9E9E" // Default grey

			// Color edges green if DAG is complete OR if DAG is running and the source dependency is completed.
			if status == StateStatusComplete ||
				(status == StateStatusRunning && slices.Contains(completedSteps, dependencyID)) {
				edgeColor = "#4CAF50" // Green
			}

			edges = append(edges, &DotEdgeSpec{
				FromNodeName: dependencyID,
				ToNodeName:   dependentID,
				Style:        edgeStyle,
				Color:        edgeColor,
				Tooltip:      fmt.Sprintf("From %s to %s", dependency.GetName(), dependent.GetName()),
			})
		}
	}

	// Use the helper function to generate the DOT string
	return dotTemplateFuncs(nodes, edges)
}

// Visualize returns a DOT graph representation of the step using the helper function
func (s *stepImplementation) Visualize() string {
	// For a single step, it's always the "current" step in its own visualization context.
	// However, we need to handle the 'Complete' status specifically for single steps.
	nodeStyle, fillColor := "solid", "#ffffff" // Start with default
	if s.state != nil {
		status := s.state.GetStatus()
		shouldBeFilled := status == StateStatusRunning || status == StateStatusComplete || status == StateStatusFailed || status == StateStatusPaused
		if shouldBeFilled {
			nodeStyle = "filled"
			switch status {
			case StateStatusFailed:
				fillColor = "#F44336" // Red
			case StateStatusPaused:
				fillColor = "#FFC107" // Yellow
			case StateStatusComplete: // Handle complete specifically for single step
				fillColor = "#4CAF50" // Green
			case StateStatusRunning:
				fillColor = "#2196F3" // Blue
			default:
				fillColor = "#2196F3" // Fallback blue
			}
		}
	}

	nodes := []*DotNodeSpec{
		{
			Name:        s.GetID(),
			DisplayName: s.GetName(),
			Tooltip:     fmt.Sprintf("Step: %s", s.GetName()),
			Shape:       "box",
			Style:       nodeStyle,
			FillColor:   fillColor,
		},
	}

	// Steps have no edges in their own visualization
	edges := []*DotEdgeSpec{}

	// Use the helper function to generate the DOT string
	return dotTemplateFuncs(nodes, edges)
}
