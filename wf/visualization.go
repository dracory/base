package wf

import (
	"bytes"
	"fmt"
	"slices"
	"text/template"
)

// DotNodeSpec represents a node in the DOT graph
type DotNodeSpec struct {
	Name        string
	DisplayName string
	Tooltip     string
	Shape       string
	Style       string
	FillColor   string
}

// DotEdgeSpec represents an edge in the DOT graph
type DotEdgeSpec struct {
	FromNodeName string
	ToNodeName   string
	Tooltip      string
	Style        string
	Color        string
}

const dotTemplateText = `digraph {
	rankdir = "LR"
	node [fontname="Arial"]
	edge [fontname="Arial"]
{{ range $node := $.Nodes}}	"{{$node.Name}}" [label="{{$node.DisplayName}}" shape={{$node.Shape}} style={{$node.Style}} tooltip="{{$node.Tooltip}}" fillcolor="{{$node.FillColor}}" {{if eq $node.Style "filled"}}fontcolor="white"{{end}}]
{{ end }}
{{ range $edge := $.Edges}}	"{{$edge.FromNodeName}}" -> "{{$edge.ToNodeName}}" [style={{$edge.Style}} tooltip="{{$edge.Tooltip}}" color="{{$edge.Color}}"]
{{ end }}}`

var dotTemplate = template.Must(template.New("digraph").Parse(dotTemplateText))

// Visualize returns a DOT graph representation of the workflow
func (p *pipelineImplementation) Visualize() string {
	// Handle empty pipeline
	if len(p.nodes) == 0 {
		return `digraph {
	rankdir = "LR"
	node [fontname="Arial"]
	edge [fontname="Arial"]
}`
	}

	nodes := make([]*DotNodeSpec, 0, len(p.nodes))
	edges := make([]*DotEdgeSpec, 0, len(p.nodes)-1)

	status := p.state.GetStatus()
	currentStepID := p.state.GetCurrentStepID()

	// Create nodes
	for i, node := range p.nodes {
		nodeStyle := "solid"
		fillColor := "#ffffff" // Default white
		isCurrentStep := currentStepID == node.GetID()

		if isCurrentStep {
			nodeStyle = "filled" // Current steps are always filled
			switch status {
			case StateStatusFailed:
				fillColor = "#F44336" // Red
			case StateStatusPaused:
				fillColor = "#FFC107" // Yellow
			case StateStatusRunning:
				fillColor = "#2196F3" // Blue
			default:
				// Default to blue if it's the current step but status is unexpected
				// or if the logic implies current step should be highlighted regardless.
				fillColor = "#2196F3"
			}
		} else if status == StateStatusComplete && i < len(p.nodes)-1 {
			// Completed steps (except the last one in a completed pipeline) are green
			// This condition assumes completed steps are only relevant *if* the pipeline itself is complete.
			// If you want completed steps to be green even in a running/paused/failed pipeline,
			// you might need: `slices.Contains(p.state.GetCompletedSteps(), node.GetID())`
			// However, the original logic only colored them green on overall completion.
			nodeStyle = "filled"
			fillColor = "#4CAF50" // Green
		}
		// else: Keep default white/solid style for non-current, non-completed steps

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

			// Highlight edges green if the pipeline is complete OR
			// if the pipeline is running and the edge leads *up to* the current step's index.
			// Note: The original logic colored edges green if running and i < len(p.nodes)-1,
			// which seems slightly off. Let's refine to color edges green if the *source* node is completed.
			// A simpler approach might be green edges only on overall completion.
			// Let's stick closer to the original intent for now: green edges on complete or up to current on running.

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

			if status == StateStatusComplete ||
				(status == StateStatusRunning && currentStepIndex != -1 && i <= currentStepIndex) {
				// Color edge green if pipeline is complete, or if running and the edge's target node index (i)
				// is less than or equal to the current step's index.
				edgeColor = "#4CAF50"
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

	buf := new(bytes.Buffer)
	err := dotTemplate.Execute(buf, struct {
		Nodes []*DotNodeSpec
		Edges []*DotEdgeSpec
	}{
		Nodes: nodes,
		Edges: edges,
	})

	if err != nil {
		// Consider logging the error instead of returning it in the DOT string
		// log.Printf("Error generating DOT graph: %v", err)
		return fmt.Sprintf("digraph { error [label=\"Error generating DOT graph: %v\"]; }", err)
	}

	return buf.String()
}

// Visualize returns a DOT graph representation of the DAG
func (d *Dag) Visualize() string {
	// Handle empty DAG
	if len(d.runnables) == 0 {
		return `digraph {
	rankdir = "LR"
	node [fontname="Arial"]
	edge [fontname="Arial"]
}`
	}

	nodes := make([]*DotNodeSpec, 0, len(d.runnables))
	edges := make([]*DotEdgeSpec, 0)

	status := d.state.GetStatus()
	currentStepID := d.state.GetCurrentStepID()
	completedSteps := d.state.GetCompletedSteps()

	// Create nodes
	for _, node := range d.runnables {
		nodeStyle := "solid"
		fillColor := "#ffffff" // Default white
		isCurrentStep := currentStepID == node.GetID()

		if isCurrentStep {
			nodeStyle = "filled"
			switch status {
			case StateStatusFailed:
				fillColor = "#F44336" // red
			case StateStatusPaused:
				fillColor = "#FFC107" // yellow
			case StateStatusRunning:
				fillColor = "#2196F3" // blue
			default:
				// Default to blue if current but status unknown/other
				fillColor = "#2196F3"
			}
		} else if status == StateStatusRunning && slices.Contains(completedSteps, node.GetID()) {
			// Only color completed steps green when workflow is running
			nodeStyle = "filled"
			fillColor = "#4CAF50" // green
		}
		// else: Keep default white/solid style

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

			// Highlight completed dependencies green if DAG is complete OR
			// if DAG is running and the source dependency is completed.
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

	buf := new(bytes.Buffer)
	err := dotTemplate.Execute(buf, struct {
		Nodes []*DotNodeSpec
		Edges []*DotEdgeSpec
	}{
		Nodes: nodes,
		Edges: edges,
	})

	if err != nil {
		// Consider logging the error
		// log.Printf("Error generating DOT graph: %v", err)
		return fmt.Sprintf("digraph { error [label=\"Error generating DOT graph: %v\"]; }", err)
	}

	return buf.String()
}

// Visualize returns a DOT graph representation of the step
func (s *stepImplementation) Visualize() string {
	nodeStyle := "solid"
	fillColor := "#ffffff" // Default white

	// Set node style based on state
	if s.state != nil {
		status := s.state.GetStatus()
		// Only fill if not in a default/waiting state
		if status == StateStatusRunning || status == StateStatusComplete || status == StateStatusFailed || status == StateStatusPaused {
			nodeStyle = "filled"
			switch status {
			case StateStatusFailed:
				fillColor = "#F44336" // red
			case StateStatusPaused:
				fillColor = "#FFC107" // yellow
			case StateStatusComplete:
				fillColor = "#4CAF50" // green
			case StateStatusRunning:
				fillColor = "#2196F3" // blue
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

	buf := new(bytes.Buffer)
	err := dotTemplate.Execute(buf, struct {
		Nodes []*DotNodeSpec
		Edges []*DotEdgeSpec // Steps have no edges in their own visualization
	}{
		Nodes: nodes,
		Edges: []*DotEdgeSpec{},
	})

	if err != nil {
		// Consider logging the error
		// log.Printf("Error generating DOT graph: %v", err)
		return fmt.Sprintf("digraph { error [label=\"Error generating DOT graph: %v\"]; }", err)
	}

	return buf.String()
}
