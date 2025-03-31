package main

import (
	"fmt"

	"github.com/dracory/base/object"
	"github.com/dracory/base/steps"
)

// ExampleContext implements StepContextInterface
type ExampleContext struct {
	*object.SerializablePropertyObject
	basePrice      int
	finalPrice     int
	stepsCompleted []string
}

func (c *ExampleContext) Data() map[string]string {
	return map[string]string{
		"base_price":  fmt.Sprintf("%d", c.basePrice),
		"final_price": fmt.Sprintf("%d", c.finalPrice),
		"steps":       fmt.Sprintf("%v", c.stepsCompleted),
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

// NewStepSetBasePrice creates a new step that sets the base price
func NewStepSetBasePrice() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) (steps.StepContextInterface, error) {
		ctx.(*ExampleContext).basePrice = 100
		ctx.(*ExampleContext).stepsCompleted = append(ctx.(*ExampleContext).stepsCompleted, "SetBasePrice")
		return ctx, nil
	})
}

// NewStepApplyDiscount creates a new step that applies a discount
func NewStepApplyDiscount() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) (steps.StepContextInterface, error) {
		ctx.(*ExampleContext).finalPrice = int(float64(ctx.(*ExampleContext).basePrice) * 0.8) // 20% discount
		ctx.(*ExampleContext).stepsCompleted = append(ctx.(*ExampleContext).stepsCompleted, "ApplyDiscount")
		return ctx, nil
	})
}

// NewStepAddShipping creates a new step that adds shipping cost
func NewStepAddShipping() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) (steps.StepContextInterface, error) {
		ctx.(*ExampleContext).finalPrice += 10
		ctx.(*ExampleContext).stepsCompleted = append(ctx.(*ExampleContext).stepsCompleted, "AddShipping")
		return ctx, nil
	})
}

// NewStepCalculateTax creates a new step that calculates tax
func NewStepCalculateTax() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) (steps.StepContextInterface, error) {
		tax := int(float64(ctx.(*ExampleContext).finalPrice) * 0.2) // 20% tax
		ctx.(*ExampleContext).finalPrice += tax
		ctx.(*ExampleContext).stepsCompleted = append(ctx.(*ExampleContext).stepsCompleted, "CalculateTax")
		return ctx, nil
	})
}

// NewDag creates a new DAG with dependent steps
func NewDag() steps.DagInterface {
	dag := steps.NewDag()

	stepSetBasePrice := NewStepSetBasePrice()
	stepApplyDiscount := NewStepApplyDiscount()
	stepAddShipping := NewStepAddShipping()
	stepCalculateTax := NewStepCalculateTax()

	dag.AddStep(stepSetBasePrice)
	dag.AddStep(stepApplyDiscount)
	dag.AddStep(stepAddShipping)
	dag.AddStep(stepCalculateTax)

	dag.AddDependency(stepApplyDiscount, stepSetBasePrice)
	dag.AddDependency(stepAddShipping, stepApplyDiscount)
	dag.AddDependency(stepCalculateTax, stepAddShipping)

	return dag
}
