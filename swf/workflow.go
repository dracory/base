package swf

import (
	"encoding/json"
	"fmt"
	"time"
)

// StepDetails contains metadata about a step
type StepDetails struct {
	Completed string
	Meta      map[string]interface{}
}

// WorkflowState represents the current state of a workflow
type WorkflowState struct {
	CurrentStep string
	History     []string
	StepDetails map[string]*StepDetails
}

// Progress represents workflow progress
type Progress struct {
	Total     int
	Completed int
	Current   int
	Pending   int
	Percents  float64
}

// Workflow represents a workflow
type Workflow struct {
	steps map[string]*Step
	state *WorkflowState
}

// NewWorkflow creates a new Workflow
func NewWorkflow() *Workflow {
	return &Workflow{
		steps: make(map[string]*Step),
		state: &WorkflowState{
			History:     make([]string, 0),
			StepDetails: make(map[string]*StepDetails),
		},
	}
}

// AddStep adds a step to the workflow
func (w *Workflow) AddStep(step *Step) {
	w.steps[step.Name] = step
	if w.state.CurrentStep == "" {
		w.state.CurrentStep = step.Name
		// w.state.History = append(w.state.History, step.Name)
	}
	if w.state.StepDetails[step.Name] == nil {
		w.state.StepDetails[step.Name] = &StepDetails{
			Meta: make(map[string]interface{}),
		}
	}
}

// GetCurrentStep returns the current step
func (w *Workflow) GetCurrentStep() *Step {
	if w.state.CurrentStep == "" {
		return nil
	}
	return w.steps[w.state.CurrentStep]
}

// SetCurrentStep sets the current step
func (w *Workflow) SetCurrentStep(step interface{}) error {
	var stepName string
	switch s := step.(type) {
	case string:
		stepName = s
	case *Step:
		stepName = s.Name
	default:
		return fmt.Errorf("invalid step type: %T", step)
	}

	if _, exists := w.steps[stepName]; !exists {
		return fmt.Errorf("step not found: %s", stepName)
	}

	// Mark the previous step as completed
	if w.state.CurrentStep != "" && w.state.CurrentStep != stepName {
		details, exists := w.state.StepDetails[w.state.CurrentStep]
		if !exists {
			details = &StepDetails{
				Meta: make(map[string]interface{}),
			}
			w.state.StepDetails[w.state.CurrentStep] = details
		}
		details.Completed = time.Now().Format(time.RFC3339)
	}

	w.state.CurrentStep = stepName
	w.state.History = append(w.state.History, stepName)
	return nil
}

// IsStepCurrent checks if a step is the current step
func (w *Workflow) IsStepCurrent(step interface{}) bool {
	var stepName string
	switch s := step.(type) {
	case string:
		stepName = s
	case *Step:
		stepName = s.Name
	default:
		return false
	}

	return w.state.CurrentStep == stepName
}

// IsStepComplete checks if a step is completed
func (w *Workflow) IsStepComplete(step interface{}) bool {
	var stepName string
	switch s := step.(type) {
	case string:
		stepName = s
	case *Step:
		stepName = s.Name
	default:
		return false
	}

	// Get step positions
	stepKeys := make([]string, 0, len(w.steps))
	for k := range w.steps {
		stepKeys = append(stepKeys, k)
	}

	currentStepPosition := -1
	stepPosition := -1

	for i, key := range stepKeys {
		if key == w.state.CurrentStep {
			currentStepPosition = i
		}
		if key == stepName {
			stepPosition = i
		}
	}

	// If the step is before the current step, it's complete
	if stepPosition < currentStepPosition {
		return true
	}

	// Check if the step is explicitly marked as completed
	details, exists := w.state.StepDetails[stepName]
	if !exists {
		return false
	}
	return details.Completed != ""
}

// GetProgress returns the workflow progress
func (w *Workflow) GetProgress() *Progress {
	total := len(w.steps)
	completed := 0
	current := 0

	// Get step positions
	stepKeys := make([]string, 0, len(w.steps))
	for k := range w.steps {
		stepKeys = append(stepKeys, k)
	}

	currentStepPosition := -1
	for i, key := range stepKeys {
		if key == w.state.CurrentStep {
			currentStepPosition = i
			break
		}
	}

	// Count completed steps
	for i, key := range stepKeys {
		if i < currentStepPosition || w.IsStepComplete(key) {
			completed++
		}
		if i == currentStepPosition {
			current = completed
		}
	}

	pending := total - completed
	percents := float64(completed) / float64(total) * 100

	return &Progress{
		Total:     total,
		Completed: completed,
		Current:   current,
		Pending:   pending,
		Percents:  percents,
	}
}

// GetSteps returns all steps
func (w *Workflow) GetSteps() map[string]*Step {
	return w.steps
}

// GetStep returns a step by name
func (w *Workflow) GetStep(name string) *Step {
	return w.steps[name]
}

// GetStepMeta returns step metadata
func (w *Workflow) GetStepMeta(step interface{}, key string) interface{} {
	var stepName string
	switch s := step.(type) {
	case string:
		stepName = s
	case *Step:
		stepName = s.Name
	default:
		return nil
	}

	details, exists := w.state.StepDetails[stepName]
	if !exists || details.Meta == nil {
		return nil
	}
	return details.Meta[key]
}

// SetStepMeta sets step metadata
func (w *Workflow) SetStepMeta(step interface{}, key string, value interface{}) {
	var stepName string
	switch s := step.(type) {
	case string:
		stepName = s
	case *Step:
		stepName = s.Name
	default:
		return
	}

	details, exists := w.state.StepDetails[stepName]
	if !exists {
		details = &StepDetails{
			Meta: make(map[string]interface{}),
		}
		w.state.StepDetails[stepName] = details
	}
	details.Meta[key] = value
}

// MarkStepAsCompleted marks a step as completed
func (w *Workflow) MarkStepAsCompleted(step interface{}) bool {
	var stepName string
	switch s := step.(type) {
	case string:
		stepName = s
	case *Step:
		stepName = s.Name
	default:
		return false
	}

	details, exists := w.state.StepDetails[stepName]
	if !exists {
		return false
	}
	details.Completed = time.Now().Format(time.RFC3339)
	return true
}

// GetState returns the current workflow state
func (w *Workflow) GetState() *WorkflowState {
	return w.state
}

// ToString serializes the workflow state to a string
func (w *Workflow) ToString() (string, error) {
	data, err := json.Marshal(w.state)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FromString deserializes the workflow state from a string
func (w *Workflow) FromString(str string) error {
	state := &WorkflowState{}
	err := json.Unmarshal([]byte(str), state)
	if err != nil {
		return err
	}
	w.state = state
	return nil
}
