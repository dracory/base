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

	// Create nodes
	for i, node := range p.nodes {
		nodeStyle := "solid"
		fillColor := "#ffffff"

		// Current step is filled blue
		if p.state.GetCurrentStepID() == node.GetID() {
			nodeStyle = "filled"
			fillColor = "#2196F3"
		} else if p.state.GetStatus() == StateStatusComplete && i < len(p.nodes)-1 {
			// Completed steps are filled green
			nodeStyle = "filled"
			fillColor = "#4CAF50"
		} else if p.state.GetStatus() == StateStatusFailed && p.state.GetCurrentStepID() == node.GetID() {
			// Failed step is filled red
			nodeStyle = "filled"
			fillColor = "#F44336"
		} else if p.state.GetStatus() == StateStatusPaused && p.state.GetCurrentStepID() == node.GetID() {
			// Paused step is filled yellow
			nodeStyle = "filled"
			fillColor = "#FFC107"
		}

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
			edgeColor := "#9E9E9E"

			// Highlight the path up to the current step
			if p.state.GetStatus() == StateStatusComplete ||
				(p.state.GetStatus() == StateStatusRunning && i < len(p.nodes)-1) {
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
		return fmt.Sprintf("Error generating DOT graph: %v", err)
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

	// Create nodes
	for _, node := range d.runnables {
		nodeStyle := "solid"
		fillColor := "#ffffff"

		// First check if this is the current step
		if d.state.GetCurrentStepID() == node.GetID() {
			nodeStyle = "filled"
			switch d.state.GetStatus() {
			case StateStatusFailed:
				fillColor = "#F44336" // red
			case StateStatusPaused:
				fillColor = "#FFC107" // yellow
			case StateStatusRunning:
				fillColor = "#2196F3" // blue
			}
		} else if d.state.GetStatus() == StateStatusRunning &&
			slices.Contains(d.state.GetCompletedSteps(), node.GetID()) {
			// Only color completed steps green when workflow is running
			nodeStyle = "filled"
			fillColor = "#4CAF50" // green
		}

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
				continue
			}

			edgeStyle := "solid"
			edgeColor := "#9E9E9E"

			// Highlight completed dependencies
			if d.state.GetStatus() == StateStatusComplete ||
				(d.state.GetStatus() == StateStatusRunning &&
					slices.Contains(d.state.GetCompletedSteps(), dependencyID)) {
				edgeColor = "#4CAF50"
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
		return fmt.Sprintf("Error generating DOT graph: %v", err)
	}

	return buf.String()
}

// Visualize returns a DOT graph representation of the step
func (s *stepImplementation) Visualize() string {
	nodeStyle := "solid"
	fillColor := "#ffffff"

	// Set node style based on state
	if s.state != nil {
		nodeStyle = "filled"
		switch s.state.GetStatus() {
		case StateStatusFailed:
			fillColor = "#F44336" // red
		case StateStatusPaused:
			fillColor = "#FFC107" // yellow
		case StateStatusComplete:
			fillColor = "#4CAF50" // green
		case StateStatusRunning:
			fillColor = "#2196F3" // blue
		default:
			nodeStyle = "solid"
			fillColor = "#ffffff" // white
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
		Edges []*DotEdgeSpec
	}{
		Nodes: nodes,
		Edges: []*DotEdgeSpec{},
	})

	if err != nil {
		return fmt.Sprintf("Error generating DOT graph: %v", err)
	}

	return buf.String()
}
