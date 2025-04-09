package wf

import (
	"context"
	"fmt"
	"slices"

	"github.com/dracory/base/arr"
	"github.com/google/uuid"
)

type pipelineImplementation struct {
	id    string
	name  string
	nodes []RunnableInterface
	state StateInterface
}

// NewPipeline creates a new pipeline
func NewPipeline() PipelineInterface {
	return &pipelineImplementation{
		id:    uuid.New().String(),
		state: NewState(),
	}
}

var _ PipelineInterface = (*pipelineImplementation)(nil)

func (p *pipelineImplementation) GetID() string {
	return p.id
}

func (p *pipelineImplementation) SetID(id string) {
	p.id = id
}

func (p *pipelineImplementation) GetName() string {
	return p.name
}

func (p *pipelineImplementation) SetName(name string) {
	p.name = name
}

func (p *pipelineImplementation) Run(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
	// If we have a saved state, use it
	if p.state.GetStatus() == StateStatusPaused {
		return p.resumeFromState(ctx, data)
	}

	// Initialize new state
	p.state = NewState()
	p.state.SetStatus(StateStatusRunning)
	p.state.SetWorkflowData(data)

	// Execute steps in order
	for _, node := range p.nodes {
		// Skip completed steps
		if slices.Contains(p.state.GetCompletedSteps(), node.GetID()) {
			continue
		}

		// Update current step
		p.state.SetCurrentStepID(node.GetID())

		// Execute step
		ctx, data, err := node.Run(ctx, data)
		if err != nil {
			p.state.SetStatus(StateStatusFailed)
			return ctx, data, err
		}

		// Mark step as completed
		p.state.AddCompletedStep(node.GetID())
		p.state.SetWorkflowData(data)
	}

	p.state.SetStatus(StateStatusComplete)
	return ctx, data, nil
}

// Pause pauses the workflow execution
func (p *pipelineImplementation) Pause() error {
	if p.state.GetStatus() != StateStatusRunning {
		return fmt.Errorf("workflow is not running")
	}
	p.state.SetStatus(StateStatusPaused)
	return nil
}

// Resume resumes the workflow execution from the last saved state
func (p *pipelineImplementation) Resume(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
	if p.state.GetStatus() != StateStatusPaused {
		return ctx, data, fmt.Errorf("workflow is not paused")
	}
	return p.resumeFromState(ctx, data)
}

// resumeFromState resumes the workflow from the saved state
func (p *pipelineImplementation) resumeFromState(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
	// Update data with saved state
	savedData := p.state.GetWorkflowData()
	for k, v := range savedData {
		data[k] = v
	}

	// Find the current step
	currentStepID := p.state.GetCurrentStepID()
	var currentStepIndex int
	for i, node := range p.nodes {
		if node.GetID() == currentStepID {
			currentStepIndex = i
			break
		}
	}

	// Execute remaining steps
	p.state.SetStatus(StateStatusRunning)
	for i := currentStepIndex; i < len(p.nodes); i++ {
		node := p.nodes[i]

		// Skip completed steps
		if slices.Contains(p.state.GetCompletedSteps(), node.GetID()) {
			continue
		}

		// Update current step
		p.state.SetCurrentStepID(node.GetID())

		// Execute step
		ctx, data, err := node.Run(ctx, data)
		if err != nil {
			p.state.SetStatus(StateStatusFailed)
			return ctx, data, err
		}

		// Mark step as completed
		p.state.AddCompletedStep(node.GetID())
		p.state.SetWorkflowData(data)
	}

	p.state.SetStatus(StateStatusComplete)
	return ctx, data, nil
}

// GetState returns the current workflow state
func (p *pipelineImplementation) GetState() StateInterface {
	return p.state
}

// SetState sets the workflow state
func (p *pipelineImplementation) SetState(state StateInterface) {
	p.state = state
}

func (p *pipelineImplementation) RunnableAdd(node ...RunnableInterface) {
	p.nodes = append(p.nodes, node...)
}

func (p *pipelineImplementation) RunnableRemove(node RunnableInterface) bool {
	id := node.GetID()

	if id == "" {
		return false
	}

	for i, n := range p.nodes {
		if n.GetID() == id {
			p.nodes = arr.IndexRemove(p.nodes, i)
			return true
		}
	}

	return false
}

func (p *pipelineImplementation) RunnableList() []RunnableInterface {
	return p.nodes
}

// State helper methods
func (p *pipelineImplementation) IsRunning() bool {
	return p.state.GetStatus() == StateStatusRunning
}

func (p *pipelineImplementation) IsPaused() bool {
	return p.state.GetStatus() == StateStatusPaused
}

func (p *pipelineImplementation) IsCompleted() bool {
	return p.state.GetStatus() == StateStatusComplete
}

func (p *pipelineImplementation) IsFailed() bool {
	return p.state.GetStatus() == StateStatusFailed
}

func (p *pipelineImplementation) IsWaiting() bool {
	return p.state.GetStatus() == "" // Initial state before running
}
