package steps

import (
	"context"

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
	dag := &Dag{}
	dag.SetName("New DAG")
	dag.id = uid.HumanUid()
	dag.runnables = make(map[string]RunnableInterface)
	dag.dependencies = make(map[string][]string)
	dag.conditionalDependencies = make(map[string]map[string]func(context.Context, map[string]any) bool)
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
	graph := d.buildDependencyGraph(ctx, data)

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

// buildDependencyGraph builds a graph of runner dependencies
func (d *Dag) buildDependencyGraph(ctx context.Context, data map[string]any) map[RunnableInterface][]RunnableInterface {
	graph := make(map[RunnableInterface][]RunnableInterface)

	// Add all nodes
	for _, node := range d.runnables {
		graph[node] = make([]RunnableInterface, 0)
	}

	// Add regular dependencies
	for depID, depList := range d.dependencies {
		if node, ok := d.runnables[depID]; ok {
			for _, dep := range depList {
				if depNode, ok := d.runnables[dep]; ok {
					graph[node] = append(graph[node], depNode)
				}
			}
		}
	}

	// Add conditional dependencies
	for depID, conditions := range d.conditionalDependencies {
		if node, ok := d.runnables[depID]; ok {
			for dep, condition := range conditions {
				if depNode, ok := d.runnables[dep]; ok {
					if condition(ctx, data) {
						graph[node] = append(graph[node], depNode)
					}
				}
			}
		}
	}

	return graph
}

func (d *Dag) getCtxAndData() (context.Context, map[string]any) {
	return context.Background(), make(map[string]any)
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

// DependencyAddIf adds a conditional dependency between two nodes.
func (d *Dag) DependencyAddIf(dependent RunnableInterface, dependency RunnableInterface, condition func(context.Context, map[string]any) bool) {
	depID := dependent.GetID()
	depNodeID := dependency.GetID()
	if depID == "" || depNodeID == "" {
		return
	}

	if _, exists := d.conditionalDependencies[depID]; !exists {
		d.conditionalDependencies[depID] = make(map[string]func(context.Context, map[string]any) bool)
	}

	d.conditionalDependencies[depID][depNodeID] = condition
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
