package steps

import (
	"context"

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
