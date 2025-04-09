package wf

// StateStatus represents the current status of a workflow
type StateStatus string

const (
	// Visit states for graph traversal
	visitStateUnvisited = "unvisited"
	visitStateVisiting  = "visiting"
	visitStateVisited   = "visited"

	// State status constants
	StateStatusRunning  = "running"
	StateStatusPaused   = "paused"
	StateStatusComplete = "complete"
	StateStatusFailed   = "failed"
)
