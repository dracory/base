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

	// conditional dependencies (DependentID, DependencyID, Condition func)
	conditionalDependencies map[string]map[string]func(context.Context, map[string]any) bool
}

// NewDag creates a new DAG
func NewDag() DagInterface {
	dag := &Dag{
		runnableSequence: make([]string, 0),
		runnables:        make(map[string]RunnableInterface),
		dependencies:     make(map[string][]string),
		conditionalDependencies: make(map[string]map[string]func(context.Context, map[string]any) bool),
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

	// Remove conditional dependencies
	delete(d.conditionalDependencies, id)

	// Remove this node from other nodes' conditional dependencies
	for depID, conditions := range d.conditionalDependencies {
		delete(conditions, id)
		if len(conditions) == 0 {
			delete(d.conditionalDependencies, depID)
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
	graph := buildDependencyGraph(d.runnables, d.dependencies, d.conditionalDependencies, ctx, data)

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
	depID := dependent.GetID()
	if depID == "" {
		return
	}

	for _, dep := range dependency {
		depNodeID := dep.GetID()
		if depNodeID == "" {
			continue
		}

		if _, exists := d.dependencies[depID]; !exists {
			d.dependencies[depID] = make([]string, 0)
		}

		if !slices.Contains(d.dependencies[depID], depNodeID) {
			d.dependencies[depID] = append(d.dependencies[depID], depNodeID)
		}
	}
}

// DependencyAddIf adds a conditional dependency between two nodes.
func (d *Dag) DependencyAddIf(dependent RunnableInterface, dependency RunnableInterface, condition func(context.Context, map[string]any) bool) {
	depID := dependent.GetID()
	if depID == "" {
		return
	}

	depNodeID := dependency.GetID()
	if depNodeID == "" {
		return
	}

	if _, exists := d.conditionalDependencies[depID]; !exists {
		d.conditionalDependencies[depID] = make(map[string]func(context.Context, map[string]any) bool)
	}

	if _, exists := d.conditionalDependencies[depID][depNodeID]; !exists {
		d.conditionalDependencies[depID][depNodeID] = condition
	}
}

// DependencyList returns all dependencies for a given node.
func (d *Dag) DependencyList(ctx context.Context, node RunnableInterface, data map[string]any) []RunnableInterface {
	result := make([]RunnableInterface, 0)
	id := node.GetID()
	if id == "" {
		return result
	}

	// Add regular dependencies
	if deps, ok := d.dependencies[id]; ok {
		for _, depID := range deps {
			if n, ok := d.runnables[depID]; ok {
				result = append(result, n)
			}
		}
	}

	// Add conditional dependencies
	if conds, ok := d.conditionalDependencies[id]; ok {
		for depID, cond := range conds {
			if n, ok := d.runnables[depID]; ok {
				if cond(ctx, data) {
					result = append(result, n)
				}
			}
		}
	}

	return result
}
