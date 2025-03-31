package steps

import (
	"context"
	"fmt"
	"sort"

	"github.com/gouniverse/uid"
)

type Step struct {
	id      string
	name    string
	data    map[string]any
	handler StepHandler
}

// NewStep creates a new step with the given execution function and optional ID.
func NewStep() StepInterface {
	step := &Step{
		data: make(map[string]any),
	}

	step.SetID(uid.HumanUid())
	step.SetName("")

	return step
}

func (s *Step) GetID() string {
	return s.id
}

func (s *Step) SetID(id string) {
	s.id = id
}

func (s *Step) GetName() string {
	return s.name
}

func (s *Step) SetName(name string) {
	s.name = name
}

// GetHandler returns the step's execution function
func (s *Step) GetHandler() StepHandler {
	return s.handler
}

// SetHandler sets the step's execution function
func (s *Step) SetHandler(fn StepHandler) {
	s.handler = fn
}

// Run executes the step's function with the given context
func (s *Step) Run(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
	return s.handler(ctx, data)
}

type Dag struct {
	// id of the dag
	id string

	// name of the dag
	name string

	// runnable sequence (IDs)
	runnableSequence []string

	// runnables (ID, RunnableInterface)
	runnables map[string]RunnableInterface

	// dependencies (DependentID, DependencyIDs []string)
	dependencies map[string][]string
}

// NewDag creates a new DAG
func NewDag() DagInterface {
	dag := &Dag{}
	dag.SetName("New DAG")
	dag.id = uid.HumanUid()
	dag.runnables = make(map[string]RunnableInterface)
	dag.dependencies = make(map[string][]string)
	return dag
}

func (d *Dag) GetID() string {
	return d.id
}

func (d *Dag) SetID(id string) {
	d.id = id
}

func (d *Dag) GetName() string {
	return d.name
}

func (d *Dag) SetName(name string) {
	d.name = name
}

// RunnableAdd adds a single node to the DAG.
func (d *Dag) RunnableAdd(node ...RunnableInterface) {
	for _, n := range node {
		if n == nil {
			continue
		}
		id := n.GetID()
		if id == "" {
			id = uid.HumanUid()
			n.SetID(id)
		}

		// Check for duplicate ID
		if _, exists := d.runnables[id]; exists {
			// Generate a new ID if there's a conflict
			newID := uid.HumanUid()
			n.SetID(newID)
			id = newID
		}

		d.runnables[id] = n
		if !contains(d.runnableSequence, id) {
			d.runnableSequence = append(d.runnableSequence, id)
		}
	}
}

// RunnableRemove removes a node from the DAG.
func (d *Dag) RunnableRemove(node RunnableInterface) bool {
	id := node.GetID()
	if id == "" {
		return false
	}

	if _, exists := d.runnables[id]; !exists {
		return false
	}

	// Remove from runnables
	delete(d.runnables, id)

	// Remove from runnableSequence
	for i, seqID := range d.runnableSequence {
		if seqID == id {
			d.runnableSequence = append(d.runnableSequence[:i], d.runnableSequence[i+1:]...)
			break
		}
	}

	// Remove dependencies
	delete(d.dependencies, id)

	// Remove this node from other nodes' dependencies
	for depID, depList := range d.dependencies {
		for i, dep := range depList {
			if dep == id {
				d.dependencies[depID] = append(depList[:i], depList[i+1:]...)
				break
			}
		}
	}

	return true
}

// RunnableList returns all runnable nodes in the DAG.
func (d *Dag) RunnableList() []RunnableInterface {
	result := make([]RunnableInterface, 0, len(d.runnables))
	for _, node := range d.runnables {
		result = append(result, node)
	}
	return result
}

// Run executes all nodes in the DAG in the correct order.
func (d *Dag) Run(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
	ordered, err := d.topologicalSort(d.buildDependencyGraph(ctx, data))
	if err != nil {
		return ctx, data, fmt.Errorf("failed to sort DAG: %w", err)
	}

	for _, runner := range ordered {
		ctx, data, err = runner.Run(ctx, data)
		if err != nil {
			return ctx, data, fmt.Errorf("failed to run step %s: %w", runner.GetID(), err)
		}
	}

	return ctx, data, nil
}

// DependencyAdd adds a dependency between two nodes.
func (d *Dag) DependencyAdd(dependent RunnableInterface, dependency ...RunnableInterface) {
	depID := dependent.GetID()
	if depID == "" {
		return
	}

	if _, exists := d.dependencies[depID]; !exists {
		d.dependencies[depID] = make([]string, 0)
	}

	for _, dep := range dependency {
		depNodeID := dep.GetID()
		if depNodeID == "" {
			continue
		}

		if !contains(d.dependencies[depID], depNodeID) {
			d.dependencies[depID] = append(d.dependencies[depID], depNodeID)
		}
	}
}

func (d *Dag) getCtxAndData() (context.Context, map[string]any) {
	return context.Background(), make(map[string]any)
}

func (d *Dag) DependencyAddIf(dependent RunnableInterface, dependency RunnableInterface, condition func(context.Context, map[string]any) bool) {
	depID := dependent.GetID()
	depNodeID := dependency.GetID()
	if depID == "" || depNodeID == "" {
		return
	}

	ctx, data := d.getCtxAndData()

	if condition != nil && condition(ctx, data) {
		if _, exists := d.dependencies[depID]; !exists {
			d.dependencies[depID] = make([]string, 0)
		}

		if !contains(d.dependencies[depID], depNodeID) {
			d.dependencies[depID] = append(d.dependencies[depID], depNodeID)
		}
	}
}

// DependencyList returns all dependencies for a given node.
func (d *Dag) DependencyList(ctx context.Context, node RunnableInterface, data map[string]any) []RunnableInterface {
	id := node.GetID()
	if id == "" {
		return nil
	}

	if deps, ok := d.dependencies[id]; ok {
		result := make([]RunnableInterface, 0, len(deps))
		for _, depID := range deps {
			if depNode, ok := d.runnables[depID]; ok {
				result = append(result, depNode)
			}
		}
		return result
	}
	return nil
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// buildDependencyGraph builds a graph of runner dependencies
func (d *Dag) buildDependencyGraph(ctx context.Context, data map[string]any) map[RunnableInterface][]RunnableInterface {
	graph := make(map[RunnableInterface][]RunnableInterface)
	for _, runner := range d.runnables {
		graph[runner] = make([]RunnableInterface, 0)
		if deps, ok := d.dependencies[runner.GetID()]; ok {
			for _, depID := range deps {
				if depNode, ok := d.runnables[depID]; ok {
					graph[runner] = append(graph[runner], depNode)
				}
			}
		}
	}
	return graph
}

// topologicalSort performs a topological sort on the dependency graph
func (d *Dag) topologicalSort(graph map[RunnableInterface][]RunnableInterface) ([]RunnableInterface, error) {
	var result []RunnableInterface
	visited := make(map[RunnableInterface]bool)
	tempMark := make(map[RunnableInterface]bool)

	var visit func(node RunnableInterface) error
	visit = func(node RunnableInterface) error {
		if tempMark[node] {
			return fmt.Errorf("cycle detected")
		}
		if visited[node] {
			return nil
		}

		tempMark[node] = true
		for _, neighbor := range graph[node] {
			if err := visit(neighbor); err != nil {
				return err
			}
		}

		tempMark[node] = false
		visited[node] = true
		result = append(result, node)
		return nil
	}

	// Sort nodes by their dependencies
	nodes := make([]RunnableInterface, 0, len(graph))
	for node := range graph {
		nodes = append(nodes, node)
	}

	// Sort nodes by number of dependencies
	sort.Slice(nodes, func(i, j int) bool {
		return len(graph[nodes[i]]) < len(graph[nodes[j]])
	})

	// Visit nodes with fewer dependencies first
	for _, node := range nodes {
		if err := visit(node); err != nil {
			return nil, err
		}
	}

	return result, nil
}
