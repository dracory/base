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

	// First reverse the result to maintain topological order
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	// Create a map to track the number of dependencies for each node
	dependencyCount := make(map[RunnableInterface]int)
	for node := range graph {
		dependencyCount[node] = 0
	}
	for _, deps := range graph {
		for _, dep := range deps {
			dependencyCount[dep]++
		}
	}

	// Sort the result to make it deterministic while preserving dependency order
	// First sort by dependency count (ascending)
	// Then sort by name to maintain a consistent order
	// Then sort by ID to maintain a consistent order
	sort.SliceStable(result, func(i, j int) bool {
		// Get the current nodes
		current := result[i]
		compare := result[j]

		// Check if current node depends on compare node
		dependsOnCompare := false
		for _, dep := range graph[current] {
			if dep == compare {
				dependsOnCompare = true
				break
			}
		}

		// Check if compare node depends on current node
		dependsOnCurrent := false
		for _, dep := range graph[compare] {
			if dep == current {
				dependsOnCurrent = true
				break
			}
		}

		// If there's a dependency relationship, respect it
		if dependsOnCompare {
			return false // current should come after compare
		}
		if dependsOnCurrent {
			return true // current should come before compare
		}

		// If both nodes have the same name, sort by ID
		if current.GetName() == compare.GetName() {
			return current.GetID() < compare.GetID()
		}

		// Otherwise, sort by name
		return current.GetName() < compare.GetName()
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
