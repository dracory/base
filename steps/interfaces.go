package steps

import "github.com/gouniverse/dataobject"

type StepContextInterface interface {
	dataobject.DataObjectInterface
}

type StepInterface interface {
	GetExecute() func(ctx StepContextInterface) error
	SetExecute(fn func(ctx StepContextInterface) error) StepInterface
	AddDependency(step StepInterface) StepInterface
	GetDependencies() []StepInterface
	Run(ctx StepContextInterface) (err error)
}

type StepsInterface interface {
	AddStep(step StepInterface)
	Run(ctx StepContextInterface) error
}

func Step(fn func(ctx StepContextInterface) error) StepInterface {
	return &stepImplementation{
		execute: fn,
	}
}

func Steps() StepsInterface {
	return &stepsImplementation{
		steps: []StepInterface{},
	}
}

type stepsImplementation struct {
	steps []StepInterface
}

func (s *stepsImplementation) AddStep(step StepInterface) {
	s.steps = append(s.steps, step)
}

func (s *stepsImplementation) Run(ctx StepContextInterface) error {
	for _, step := range s.steps {
		err := step.Run(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

type stepImplementation struct {
	execute      func(ctx StepContextInterface) error
	dependencies []StepInterface
}

func (s *stepImplementation) GetExecute() func(ctx StepContextInterface) error {
	return s.execute
}

func (s *stepImplementation) SetExecute(fn func(ctx StepContextInterface) error) StepInterface {
	s.execute = fn
	return s
}

func (s *stepImplementation) AddDependency(step StepInterface) StepInterface {
	s.dependencies = append(s.dependencies, step)
	return s
}

func (s *stepImplementation) GetDependencies() []StepInterface {
	return s.dependencies
}

func (s *stepImplementation) Run(ctx StepContextInterface) error {
	return s.execute(ctx)
}
