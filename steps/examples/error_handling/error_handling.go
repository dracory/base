package main

import (
	"fmt"

	"github.com/dracory/base/object"
	"github.com/dracory/base/steps"
)

// ExampleContext implements StepContextInterface
type ExampleContext struct {
	*object.SerializablePropertyObject
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
		SerializablePropertyObject: object.NewSerializablePropertyObject().(*object.SerializablePropertyObject),
	}
}

// NewStepSetInitialValue creates a step that sets an initial value
func NewStepSetInitialValue() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) error {
		ctx.(*ExampleContext).value = 100
		ctx.(*ExampleContext).stepsCompleted = append(ctx.(*ExampleContext).stepsCompleted, "SetInitialValue")
		return nil
	})
}

// NewStepProcessData creates a step that processes data
func NewStepProcessData() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) error {
		ctx.(*ExampleContext).value *= 2
		ctx.(*ExampleContext).stepsCompleted = append(ctx.(*ExampleContext).stepsCompleted, "ProcessData")
		return nil
	})
}

// NewStepVerifyData creates a step that verifies data
func NewStepVerifyData() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) error {
		ctx.(*ExampleContext).stepsCompleted = append(ctx.(*ExampleContext).stepsCompleted, "VerifyData")

		// Simulate an error
		ctx.(*ExampleContext).errorCount++
		return fmt.Errorf("data verification failed")
	})
}

// NewDag creates a DAG with error handling
func NewDag() steps.DagInterface {
	dag := steps.NewDag()

	stepSetInitialValue := NewStepSetInitialValue()
	stepProcessData := NewStepProcessData().AddDependency(stepSetInitialValue)
	stepVerifyData := NewStepVerifyData().AddDependency(stepProcessData)

	dag.AddStep(stepSetInitialValue)
	dag.AddStep(stepProcessData)
	dag.AddStep(stepVerifyData)

	return dag
}
