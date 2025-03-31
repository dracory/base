package steps

import (
	"context"
	"errors"
	"sort"
)

// visitNode performs a depth-first search to detect cycles and build the topological order
func visitNode(node RunnableInterface, graph map[RunnableInterface][]RunnableInterface, visited map[RunnableInterface]bool, tempMark map[RunnableInterface]bool, result *[]RunnableInterface) error {
	if tempMark[node] {
		return errors.New("cycle detected")
	}
	if visited[node] {
		return nil
	}

	tempMark[node] = true
	for _, dep := range graph[node] {
		if err := visitNode(dep, graph, visited, tempMark, result); err != nil {
			return err
		}
	}

	tempMark[node] = false
	visited[node] = true
	*result = append([]RunnableInterface{node}, *result...)
	return nil
}

// topologicalSort performs a topological sort on the dependency graph
func topologicalSort(graph map[RunnableInterface][]RunnableInterface) ([]RunnableInterface, error) {
	visited := make(map[RunnableInterface]bool)
	tempMark := make(map[RunnableInterface]bool)
	result := []RunnableInterface{}

	// Start with any node (since the graph is connected)
	for node := range graph {
		if err := visitNode(node, graph, visited, tempMark, &result); err != nil {
			return nil, err
		}
	}

	// Sort the result to make it deterministic
	sort.Slice(result, func(i, j int) bool {
		return result[i].GetName() < result[j].GetName()
	})

	return result, nil
}

// buildDependencyGraph builds a graph of runner dependencies
func buildDependencyGraph(runnables map[string]RunnableInterface, dependencies map[string][]string, conditionalDependencies map[string]map[string]func(context.Context, map[string]any) bool, ctx context.Context, data map[string]any) map[RunnableInterface][]RunnableInterface {
	graph := make(map[RunnableInterface][]RunnableInterface)

	// Add all nodes
	for _, node := range runnables {
		graph[node] = make([]RunnableInterface, 0)
	}

	// Add regular dependencies
	for depID, depList := range dependencies {
		if depNode, ok := runnables[depID]; ok {
			for _, nodeID := range depList {
				if node, ok := runnables[nodeID]; ok {
					graph[depNode] = append(graph[depNode], node)
				}
			}
		}
	}

	// Add conditional dependencies
	for depID, conditions := range conditionalDependencies {
		if depNode, ok := runnables[depID]; ok {
			for nodeID, condition := range conditions {
				if node, ok := runnables[nodeID]; ok {
					if condition(ctx, data) {
						graph[depNode] = append(graph[depNode], node)
					}
				}
			}
		}
	}

	return graph
}
