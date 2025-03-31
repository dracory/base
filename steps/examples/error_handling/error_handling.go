package main

import (
	"fmt"

	"github.com/dracory/base/steps"
)

// ExampleContext implements StepContextInterface
type ExampleContext struct {
	steps.StepContextInterface
	value          int
	stepsCompleted []string
	errorCount     int
}

func (c *ExampleContext) Data() map[string]string {
	return map[string]string{
		"value":  fmt.Sprintf("%d", c.value),
		"steps":  fmt.Sprintf("%v", c.stepsCompleted),
		"errors": fmt.Sprintf("%d", c.errorCount),
	}
}

func (c *ExampleContext) Name() string {
	return c.Get("name").(string)
}

func (c *ExampleContext) SetName(name string) steps.StepContextInterface {
	c.Set("name", name)
	return c
}

// NewExampleContext creates a new ExampleContext
func NewExampleContext() *ExampleContext {
	return &ExampleContext{
		StepContextInterface: steps.NewStepContext(),
	}
}

// NewStepSetInitialValue creates a step that sets an initial value
func NewStepSetInitialValue() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) (steps.StepContextInterface, error) {
		ctx.Set("value", 1)
		return ctx, nil
	})
}

// NewStepProcessData creates a step that processes data
func NewStepProcessData() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) (steps.StepContextInterface, error) {
		if !ctx.Has("value") {
			return ctx, fmt.Errorf("value not found")
		}
		value := ctx.Get("value").(int)
		ctx.Set("value", value*2)
		return ctx, nil
	})
}

// NewStepVerifyData creates a step that verifies data
func NewStepVerifyData() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) (steps.StepContextInterface, error) {
		return ctx, fmt.Errorf("intentional error")
	})
}

// NewDag creates a DAG with error handling
func NewDag() steps.DagInterface {
	dag := steps.NewDag()

	step1 := NewStepSetInitialValue()
	step2 := NewStepProcessData()
	stepWithError := NewStepVerifyData()

	dag.AddStep(step1)
	dag.AddStep(step2)
	dag.AddStep(stepWithError)

	dag.AddDependency(step2, step1)
	dag.AddDependency(stepWithError, step2)

	return dag
}
