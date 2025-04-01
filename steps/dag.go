package steps

import (
	"context"

	"slices"

	"github.com/gouniverse/uid"
)

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
	dag := &Dag{
		runnableSequence: make([]string, 0),
		runnables:        make(map[string]RunnableInterface),
		dependencies:     make(map[string][]string),
	}
	dag.SetName("New DAG")
	dag.id = uid.HumanUid()
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
		if !slices.Contains(d.runnableSequence, id) {
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

// Run executes all nodes in the DAG in the correct order
func (d *Dag) Run(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
	// Build dependency graph
	graph := buildDependencyGraph(d.runnables, d.dependencies)

	// Get execution order
	order, err := topologicalSort(graph)
	if err != nil {
		return ctx, data, err
	}

	// Execute steps in order
	for _, node := range order {
		ctx, data, err = node.Run(ctx, data)
		if err != nil {
			return ctx, data, err
		}
	}

	return ctx, data, nil
}

// DependencyAdd adds a dependency between two nodes.
func (d *Dag) DependencyAdd(dependent RunnableInterface, dependency ...RunnableInterface) {
	dependentID := dependent.GetID()
	for _, dep := range dependency {
		depID := dep.GetID()
		d.dependencies[dependentID] = append(d.dependencies[dependentID], depID)
	}
}

// DependencyList returns all dependencies for a given node.
func (d *Dag) DependencyList(ctx context.Context, node RunnableInterface, data map[string]any) []RunnableInterface {
	dependencies := []RunnableInterface{}
	
	// Get all direct dependencies
	dependentID := node.GetID()
	if deps, ok := d.dependencies[dependentID]; ok {
		for _, depID := range deps {
			dep, ok := d.runnables[depID]
			if !ok {
				continue
			}

			// Add regular dependency
			dependencies = append(dependencies, dep)
		}
	}

	return dependencies
}
