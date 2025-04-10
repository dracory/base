package wf

import (
	"fmt"
	"slices"
	"strings"
)

// --- Constants (Color, Node Style, Edge Style) remain the same ---

// Color constants for DOT graph visualization
const (
	colorWhite  = "#ffffff" // Default node fill
	colorRed    = "#F44336" // Failed status
	colorYellow = "#FFC107" // Paused status
	colorBlue   = "#2196F3" // Running status
	colorGreen  = "#4CAF50" // Completed status/edges
	colorGrey   = "#9E9E9E" // Default edge, fallback fill
)

// Node style constants
const (
	nodeStyleSolid  = "solid"
	nodeStyleFilled = "filled"
)

// Edge style constants
const (
	edgeStyleSolid = "solid"
	// edgeStyleDashed = "dashed" // Example if needed later
)

// --- Structs (DotNodeSpec, DotEdgeSpec) remain the same ---

// DotNodeSpec represents a node in the DOT graph
type DotNodeSpec struct {
	Name        string
	DisplayName string
	Tooltip     string
	Shape       string
	Style       string // Use nodeStyleSolid or nodeStyleFilled
	FillColor   string
	// FontColor is handled conditionally by dotTemplateFuncs
}

// DotEdgeSpec represents an edge in the DOT graph
type DotEdgeSpec struct {
	FromNodeName string
	ToNodeName   string
	Tooltip      string
	Style        string // Use edgeStyleSolid, etc.
	Color        string
}

// --- Helper Functions ---

// escapeDotString remains the same
func escapeDotString(s string) string {
	quoted := fmt.Sprintf("%q", s)
	if len(quoted) >= 2 {
		return quoted[1 : len(quoted)-1]
	}
	return ""
}

// getWorkflowStateInfo safely extracts common state information.
func getWorkflowStateInfo(state StateInterface) (status StateStatus, currentStepID string, completedSteps []string) {
	status = StateStatus("") // Default empty status
	currentStepID = ""
	completedSteps = []string{}
	if state != nil {
		status = state.GetStatus()
		currentStepID = state.GetCurrentStepID()
		completedSteps = state.GetCompletedSteps() // Get completed steps only if state exists
	}
	return status, currentStepID, completedSteps
}

// getNodeStyleAndColor remains the same (handles CURRENT node styling based on status)
func getNodeStyleAndColor(state StateInterface, isCurrentStep bool) (style, fillColor string) {
	style = nodeStyleSolid
	fillColor = colorWhite

	if state == nil || !isCurrentStep {
		return style, fillColor
	}

	status := state.GetStatus()
	shouldBeFilled := status == StateStatusRunning || status == StateStatusFailed || status == StateStatusPaused

	if shouldBeFilled {
		style = nodeStyleFilled
		switch status {
		case StateStatusFailed:
			fillColor = colorRed
		case StateStatusPaused:
			fillColor = colorYellow
		case StateStatusRunning:
			fillColor = colorBlue
		default:
			fillColor = colorGrey
		}
	}
	return style, fillColor
}

// getPipelineNodeStyleAndColor determines style/color for a pipeline node.
func getPipelineNodeStyleAndColor(state StateInterface, nodeID string, index, nodeCount int, isCurrentStep bool) (style, fillColor string) {
	if isCurrentStep {
		// Delegate to the common helper for current step styling
		return getNodeStyleAndColor(state, true)
	}

	// Default for non-current
	style = nodeStyleSolid
	fillColor = colorWhite

	// Pipeline specific logic for NON-CURRENT steps:
	status, _, _ := getWorkflowStateInfo(state)
	if status == StateStatusComplete && index < nodeCount-1 {
		style = nodeStyleFilled
		fillColor = colorGreen
	}
	return style, fillColor
}

// getDagNodeStyleAndColor remains the same (handles DAG node styling)
func getDagNodeStyleAndColor(state StateInterface, nodeID string, isCurrentStep bool, completedSteps []string) (style, fillColor string) {
	if isCurrentStep {
		return getNodeStyleAndColor(state, true)
	}

	style = nodeStyleSolid
	fillColor = colorWhite

	if state == nil {
		return style, fillColor
	}

	status := state.GetStatus()
	if status == StateStatusRunning && slices.Contains(completedSteps, nodeID) {
		style = nodeStyleFilled
		fillColor = colorGreen
	}
	return style, fillColor
}

// createDotNodeSpec creates a DotNodeSpec with common defaults.
func createDotNodeSpec(node RunnableInterface, style, fillColor string) *DotNodeSpec {
	return &DotNodeSpec{
		Name:        node.GetID(),
		DisplayName: node.GetName(),
		Shape:       "box", // Common shape
		Style:       style,
		FillColor:   fillColor,
		// Tooltip is handled by dotTemplateFuncs if empty
	}
}

// createDotEdgeSpec creates a DotEdgeSpec.
func createDotEdgeSpec(fromNode, toNode RunnableInterface, style, color string) *DotEdgeSpec {
	return &DotEdgeSpec{
		FromNodeName: fromNode.GetID(),
		ToNodeName:   toNode.GetID(),
		Style:        style,
		Color:        color,
		Tooltip:      fmt.Sprintf("From %s to %s", fromNode.GetName(), toNode.GetName()),
	}
}

// dotTemplateFuncs remains the same (generates the final DOT string)
func dotTemplateFuncs(nodes []*DotNodeSpec, edges []*DotEdgeSpec) string {
	var sb strings.Builder

	sb.WriteString("digraph {\n")
	sb.WriteString("\trankdir = \"LR\";\n")
	sb.WriteString("\tnode [fontname=\"Arial\"];\n")
	sb.WriteString("\tedge [fontname=\"Arial\"];\n")

	for _, node := range nodes {
		displayName := node.DisplayName
		if displayName == "" {
			displayName = node.Name
		}
		tooltip := node.Tooltip
		if tooltip == "" {
			tooltip = fmt.Sprintf("Step: %s", displayName)
		}

		sb.WriteString(fmt.Sprintf("\t\"%s\" [label=\"%s\" shape=%s style=%s tooltip=\"%s\" fillcolor=\"%s\"",
			escapeDotString(node.Name),
			escapeDotString(displayName),
			node.Shape,
			node.Style,
			escapeDotString(tooltip),
			node.FillColor,
		))
		if node.Style == nodeStyleFilled {
			sb.WriteString(" fontcolor=\"white\"")
		}
		sb.WriteString("];\n")
	}

	for _, edge := range edges {
		sb.WriteString(fmt.Sprintf("\t\"%s\" -> \"%s\" [style=%s tooltip=\"%s\" color=\"%s\"];\n",
			escapeDotString(edge.FromNodeName),
			escapeDotString(edge.ToNodeName),
			edge.Style,
			escapeDotString(edge.Tooltip),
			edge.Color,
		))
	}

	sb.WriteString("}\n")
	return sb.String()
}

// --- Visualize Methods ---

// Visualize returns a DOT graph representation of the pipeline.
func (p *pipelineImplementation) Visualize() string {
	if len(p.nodes) == 0 {
		return dotTemplateFuncs([]*DotNodeSpec{}, []*DotEdgeSpec{})
	}

	nodes := make([]*DotNodeSpec, 0, len(p.nodes))
	edges := make([]*DotEdgeSpec, 0, len(p.nodes)-1)

	status, currentStepID, _ := getWorkflowStateInfo(p.state) // Use helper

	// Determine current step index once for edge coloring
	currentStepIndex := -1
	if currentStepID != "" {
		currentStepIndex = slices.IndexFunc(p.nodes, func(n RunnableInterface) bool {
			return n.GetID() == currentStepID
		})
	}

	// --- Create Node Specs ---
	for i, node := range p.nodes {
		isCurrentStep := currentStepID == node.GetID()
		// Use pipeline-specific node styling helper
		nodeStyle, fillColor := getPipelineNodeStyleAndColor(p.state, node.GetID(), i, len(p.nodes), isCurrentStep)
		nodes = append(nodes, createDotNodeSpec(node, nodeStyle, fillColor)) // Use helper
	}

	// --- Create Edge Specs ---
	for i := 1; i < len(p.nodes); i++ {
		fromNode := p.nodes[i-1]
		toNode := p.nodes[i]
		edgeStyle := edgeStyleSolid
		edgeColor := colorGrey // Default

		// Edge coloring logic (remains specific to pipeline)
		if status == StateStatusComplete ||
			(status == StateStatusRunning && currentStepIndex != -1 && i <= currentStepIndex) {
			edgeColor = colorGreen
		}
		edges = append(edges, createDotEdgeSpec(fromNode, toNode, edgeStyle, edgeColor)) // Use helper
	}

	return dotTemplateFuncs(nodes, edges)
}

// Visualize returns a DOT graph representation of the DAG.
func (d *Dag) Visualize() string {
	if len(d.runnables) == 0 {
		return dotTemplateFuncs([]*DotNodeSpec{}, []*DotEdgeSpec{})
	}

	nodes := make([]*DotNodeSpec, 0, len(d.runnables))
	edges := make([]*DotEdgeSpec, 0)

	status, currentStepID, completedSteps := getWorkflowStateInfo(d.state) // Use helper

	// --- Create Node Specs ---
	for _, node := range d.runnables {
		isCurrentStep := currentStepID == node.GetID()
		// Use DAG-specific node styling helper
		nodeStyle, fillColor := getDagNodeStyleAndColor(d.state, node.GetID(), isCurrentStep, completedSteps)
		nodes = append(nodes, createDotNodeSpec(node, nodeStyle, fillColor)) // Use helper
	}

	// --- Create Edge Specs ---
	for dependentID, dependencyIDs := range d.dependencies {
		dependent, depExists := d.runnables[dependentID]
		if !depExists {
			continue
		}

		for _, dependencyID := range dependencyIDs {
			dependency, depExists2 := d.runnables[dependencyID]
			if !depExists2 {
				continue
			}

			edgeStyle := edgeStyleSolid
			edgeColor := colorGrey // Default

			// Edge coloring logic (remains specific to DAG)
			isSourceCompleted := slices.Contains(completedSteps, dependencyID)
			if status == StateStatusComplete || (status == StateStatusRunning && isSourceCompleted) {
				edgeColor = colorGreen
			}
			edges = append(edges, createDotEdgeSpec(dependency, dependent, edgeStyle, edgeColor)) // Use helper (Note: dependency -> dependent)
		}
	}

	return dotTemplateFuncs(nodes, edges)
}

// Visualize returns a DOT graph representation of the step.
// (This one is quite specific, less benefit from further abstraction here)
func (s *stepImplementation) Visualize() string {
	nodeStyle, fillColor := nodeStyleSolid, colorWhite // Default

	status, _, _ := getWorkflowStateInfo(s.state) // Use helper for consistency

	// Single step visualization logic is simple and direct
	switch status {
	case StateStatusRunning:
		nodeStyle = nodeStyleFilled
		fillColor = colorBlue
	case StateStatusComplete: // Single step complete is green
		nodeStyle = nodeStyleFilled
		fillColor = colorGreen
	case StateStatusFailed:
		nodeStyle = nodeStyleFilled
		fillColor = colorRed
	case StateStatusPaused:
		nodeStyle = nodeStyleFilled
		fillColor = colorYellow
	}

	// Build node spec directly (using helper is minimal gain here)
	nodeSpec := &DotNodeSpec{
		Name:        s.GetID(),
		DisplayName: s.GetName(),
		Shape:       "box",
		Style:       nodeStyle,
		FillColor:   fillColor,
	}

	edges := []*DotEdgeSpec{} // Steps have no edges

	return dotTemplateFuncs([]*DotNodeSpec{nodeSpec}, edges)
}
