package steps

import (
	"fmt"

	"github.com/dracory/base/object"
)

type StepContext struct {
	*object.SerializablePropertyObject
}

// NewStepContext creates a new step context
func NewStepContext() StepContextInterface {
	return &StepContext{
		SerializablePropertyObject: object.NewSerializablePropertyObject().(*object.SerializablePropertyObject),
	}
}

// Name returns the name of the context
func (s *StepContext) Name() string {
	return s.Get("name").(string)
}

// SetName sets the name of the context
func (s *StepContext) SetName(name string) StepContextInterface {
	s.Set("name", name)
	return s
}

// Step implements the StepInterface using a serializable property object
// for better identification and serialization capabilities.
type Step struct {
	*object.SerializablePropertyObject
	execute      func(ctx StepContextInterface) error
	dependencies []StepInterface
	conditionalDependencies []struct {
		step      StepInterface
		condition func(ctx StepContextInterface) bool
	}
	skipStep func(ctx StepContextInterface) bool
}

// NewStep creates a new step with the given execution function and optional ID.
func NewStep(fn func(ctx StepContextInterface) error, id ...string) StepInterface {
	step := &Step{
		SerializablePropertyObject: object.NewSerializablePropertyObject().(*object.SerializablePropertyObject),
		execute:                    fn,
	}
	if len(id) > 0 && id[0] != "" {
		step.SetID(id[0])
	}
	return step
}

func (s *Step) Name() string {
	return s.Get("name").(string)
}

func (s *Step) SetName(name string) StepInterface {
	s.Set("name", name)
	return s
}

// GetExecute returns the step's execution function
func (s *Step) GetExecute() func(ctx StepContextInterface) error {
	return s.execute
}

// SetExecute sets the step's execution function
func (s *Step) SetExecute(fn func(ctx StepContextInterface) error) StepInterface {
	s.execute = fn
	return s
}

// AddDependency adds a dependency to the step
func (s *Step) AddDependency(step StepInterface) StepInterface {
	s.dependencies = append(s.dependencies, step)
	return s
}

// AddDependencies adds multiple dependencies to the step
func (s *Step) AddDependencies(steps ...StepInterface) StepInterface {
	s.dependencies = append(s.dependencies, steps...)
	return s
}

// AddDependencyIf adds a dependency that only exists if the condition is true.
func (s *Step) AddDependencyIf(step StepInterface, condition func(ctx StepContextInterface) bool) StepInterface {
	s.conditionalDependencies = append(s.conditionalDependencies, struct {
		step      StepInterface
		condition func(ctx StepContextInterface) bool
	}{
		step:      step,
		condition: condition,
	})
	return s
}

// GetDependencies returns all active dependencies based on context.
func (s *Step) GetDependencies(ctx StepContextInterface) []StepInterface {
	var activeDependencies []StepInterface
	
	// Add regular dependencies
	activeDependencies = append(activeDependencies, s.dependencies...)
	
	// Add conditional dependencies that match the condition
	for _, dep := range s.conditionalDependencies {
		if dep.condition(ctx) {
			activeDependencies = append(activeDependencies, dep.step)
		}
	}
	
	return activeDependencies
}

// SetSkipStep sets the skip condition for the step
func (s *Step) SetSkipStep(fn func(ctx StepContextInterface) bool) StepInterface {
	s.skipStep = fn
	return s
}

// SkipStep returns true if the step should be skipped
func (s *Step) SkipStep(ctx StepContextInterface) bool {
	if s.skipStep != nil {
		return s.skipStep(ctx)
	}
	return false
}

// Run executes the step's function with the given context
func (s *Step) Run(ctx StepContextInterface) error {
	if s.SkipStep(ctx) {
		return nil
	}
	return s.execute(ctx)
}

// ToJSON serializes the step to JSON
func (s *Step) ToJSON() ([]byte, error) {
	return s.SerializablePropertyObject.ToJSON()
}

// FromJSON deserializes JSON data into the step
func (s *Step) FromJSON(data []byte) error {
	return s.SerializablePropertyObject.FromJSON(data)
}

type Dag struct {
	*object.SerializablePropertyObject
	steps   []StepInterface
	stepMap map[StepInterface]struct{}
}

// NewDag creates a new DAG
func NewDag() DagInterface {
	return &Dag{
		SerializablePropertyObject: object.NewSerializablePropertyObject().(*object.SerializablePropertyObject),
		steps:                      []StepInterface{},
		stepMap:                    make(map[StepInterface]struct{}),
	}
}

func (d *Dag) AddStep(step StepInterface) {
	if _, exists := d.stepMap[step]; !exists {
		d.steps = append(d.steps, step)
		d.stepMap[step] = struct{}{}
	}
}

func (d *Dag) AddSteps(steps ...StepInterface) {
	for _, step := range steps {
		if _, exists := d.stepMap[step]; !exists {
			d.steps = append(d.steps, step)
			d.stepMap[step] = struct{}{}
		}
	}
}

func (d *Dag) RemoveStep(step StepInterface) bool {
	if _, exists := d.stepMap[step]; !exists {
		return false
	}

	delete(d.stepMap, step)

	for i := 0; i < len(d.steps); i++ {
		if d.steps[i] == step {
			d.steps = append(d.steps[:i], d.steps[i+1:]...)
			break
		}
	}

	return true
}

func (d *Dag) GetSteps() []StepInterface {
	return d.steps
}

func (d *Dag) Run(ctx StepContextInterface) error {
	graph := d.buildDependencyGraph()
	sortedSteps, err := d.topologicalSort(graph)
	if err != nil {
		return err
	}

	return d.executeSteps(sortedSteps, ctx)
}

func (d *Dag) executeSteps(steps []StepInterface, ctx StepContextInterface) error {
	for _, step := range steps {
		if err := step.Run(ctx); err != nil {
			return err
		}
	}
	return nil
}

// buildDependencyGraph builds a graph of step dependencies
func (d *Dag) buildDependencyGraph() map[StepInterface][]StepInterface {
	graph := make(map[StepInterface][]StepInterface)
	ctx := NewStepContext()
	for _, step := range d.steps {
		graph[step] = step.GetDependencies(ctx)
	}
	return graph
}

func (d *Dag) topologicalSort(graph map[StepInterface][]StepInterface) ([]StepInterface, error) {
	var sortedSteps []StepInterface
	visited := make(map[StepInterface]string)

	var visit func(step StepInterface) error
	visit = func(step StepInterface) error {
		if visited[step] == "visiting" {
			return fmt.Errorf("cycle detected in DAG")
		}
		if visited[step] == "visited" {
			return nil
		}

		visited[step] = "visiting"
		for _, dep := range graph[step] {
			if err := visit(dep); err != nil {
				return err
			}
		}

		visited[step] = "visited"
		sortedSteps = append(sortedSteps, step)
		return nil
	}

	for _, step := range d.steps {
		if err := visit(step); err != nil {
			return nil, err
		}
	}

	return sortedSteps, nil
}
