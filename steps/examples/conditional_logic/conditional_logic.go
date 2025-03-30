package main

import (
	"github.com/dracory/base/object"
	"github.com/dracory/base/steps"
)

// OrderContext implements StepContextInterface
type OrderContext struct {
	*object.SerializablePropertyObject
	stepsExecuted []string
}

// NewOrderContext creates a new OrderContext
func NewOrderContext(orderType string) *OrderContext {
	ctx := &OrderContext{
		SerializablePropertyObject: object.NewSerializablePropertyObject().(*object.SerializablePropertyObject),
	}
	ctx.Set("orderType", orderType)
	ctx.Set("totalAmount", 100.0)
	ctx.Set("stepsToSkip", map[string]bool{})
	return ctx
}

// Name returns the context name
func (c *OrderContext) Name() string {
	return c.Get("name").(string)
}

// SetName sets the context name
func (c *OrderContext) SetName(name string) steps.StepContextInterface {
	c.Set("name", name)
	return c
}

// NewStepProcessOrder creates a step that processes the order
func NewStepProcessOrder() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) error {
		ctx.(*OrderContext).stepsExecuted = append(ctx.(*OrderContext).stepsExecuted, "ProcessOrder")
		return nil
	})
}

// NewStepApplyDiscount creates a step that applies a discount
func NewStepApplyDiscount() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) error {
		orderCtx := ctx.(*OrderContext)
		totalAmount := orderCtx.Get("totalAmount").(float64)
		orderCtx.Set("totalAmount", totalAmount*0.9) // 10% discount
		orderCtx.stepsExecuted = append(orderCtx.stepsExecuted, "ApplyDiscount")
		return nil
	})
}

// NewStepAddShipping creates a step that adds shipping cost
func NewStepAddShipping() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) error {
		orderCtx := ctx.(*OrderContext)
		stepsToSkip := orderCtx.Get("stepsToSkip")
		skipShipping := false
		if stepsToSkip != nil {
			if skip, ok := stepsToSkip.(map[string]bool)["addShipping"]; ok {
				skipShipping = skip
			}
		}
		if !skipShipping {
			totalAmount := orderCtx.Get("totalAmount").(float64)
			orderCtx.Set("totalAmount", totalAmount+5.0) // Fixed shipping cost
			orderCtx.stepsExecuted = append(orderCtx.stepsExecuted, "AddShipping")
		}
		return nil
	})
}

// NewStepCalculateTax creates a step that calculates tax
func NewStepCalculateTax() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) error {
		orderCtx := ctx.(*OrderContext)
		stepsToSkip := orderCtx.Get("stepsToSkip")
		skipTax := false
		if stepsToSkip != nil {
			if skip, ok := stepsToSkip.(map[string]bool)["calculateTax"]; ok {
				skipTax = skip
			}
		}
		if !skipTax {
			totalAmount := orderCtx.Get("totalAmount").(float64)
			orderCtx.Set("totalAmount", totalAmount*1.2) // 20% tax
			orderCtx.stepsExecuted = append(orderCtx.stepsExecuted, "CalculateTax")
		}
		return nil
	})
}

// NewStepDetermineSkip creates a step that determines which steps to skip
func NewStepDetermineSkip() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) error {
		orderCtx := ctx.(*OrderContext)
		orderType := orderCtx.Get("orderType").(string)
		stepsToSkip := orderCtx.Get("stepsToSkip")
		if stepsToSkip == nil {
			stepsToSkip = map[string]bool{}
			orderCtx.Set("stepsToSkip", stepsToSkip)
		}
		
		// Skip shipping for digital and subscription orders
		if orderType == "digital" || orderType == "subscription" {
			stepsToSkip.(map[string]bool)["addShipping"] = true
		}
		
		// Skip tax for subscription orders
		if orderType == "subscription" {
			stepsToSkip.(map[string]bool)["calculateTax"] = true
		}
		
		orderCtx.stepsExecuted = append(orderCtx.stepsExecuted, "DetermineSkip")
		return nil
	})
}

// NewDag creates a DAG with conditional logic
func NewDag(orderType string) (steps.DagInterface, error) {
	dag := steps.NewDag()

	// Create steps
	processOrder := NewStepProcessOrder()
	applyDiscount := NewStepApplyDiscount()
	addShipping := NewStepAddShipping()
	calculateTax := NewStepCalculateTax()
	determineSkip := NewStepDetermineSkip()

	// Set up dependencies
	applyDiscount.AddDependency(processOrder)
	determineSkip.AddDependency(applyDiscount)
	addShipping.AddDependency(determineSkip)
	calculateTax.AddDependency(determineSkip)

	// Add steps to DAG
	dag.AddStep(processOrder)
	dag.AddStep(applyDiscount)
	dag.AddStep(determineSkip)
	dag.AddStep(addShipping)
	dag.AddStep(calculateTax)

	return dag, nil
}

// NewConditionalDag creates a DAG with conditional logic
func NewConditionalDag(orderType string) steps.DagInterface {
	dag, err := NewDag(orderType)
	if err != nil {
		panic(err)
	}

	// Set step names for identification
	steps := dag.GetSteps()
	for _, step := range steps {
		switch step.GetID() {
		case steps[0].GetID():
			step.SetName("processOrder")
		case steps[1].GetID():
			step.SetName("applyDiscount")
		case steps[2].GetID():
			step.SetName("determineSkip")
		case steps[3].GetID():
			step.SetName("addShipping")
		case steps[4].GetID():
			step.SetName("calculateTax")
		}
	}
	return dag
}
