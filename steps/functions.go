package steps

import (
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
	*result = append(*result, node)
	return nil
}

// topologicalSort performs a topological sort on the dependency graph
func topologicalSort(graph map[RunnableInterface][]RunnableInterface) ([]RunnableInterface, error) {
	visited := make(map[RunnableInterface]bool)
	tempMark := make(map[RunnableInterface]bool)
	result := []RunnableInterface{}

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
