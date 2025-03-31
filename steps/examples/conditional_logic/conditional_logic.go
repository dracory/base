package main

import (
	"fmt"

	"github.com/dracory/base/steps"
)

// OrderContext implements StepContextInterface
type OrderContext struct {
	steps.StepContextInterface
}

// NewOrderContext creates a new OrderContext
func NewOrderContext(orderType string, totalAmount float64) steps.StepContextInterface {
	stepContext := steps.NewStepContext()
	ctx := &OrderContext{
		StepContextInterface: stepContext,
	}
	ctx.Set("orderType", orderType)
	ctx.Set("totalAmount", totalAmount)
	ctx.Set("stepsExecuted", []string{})
	return ctx
}

// NewStepProcessOrder creates a step that processes the order
func NewStepProcessOrder() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) (steps.StepContextInterface, error) {
		if ctx == nil {
			return ctx, fmt.Errorf("context is nil")
		}
		stepsExecuted := ctx.Get("stepsExecuted").([]string)
		ctx.Set("stepsExecuted", append(stepsExecuted, "ProcessOrder"))
		return ctx, nil
	}).SetName("ProcessOrder")
}

// NewStepApplyDiscount creates a step that applies a discount
func NewStepApplyDiscount() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) (steps.StepContextInterface, error) {
		if ctx == nil {
			return ctx, fmt.Errorf("context is nil")
		}
		totalAmount := ctx.Get("totalAmount").(float64)
		ctx.Set("totalAmount", totalAmount*0.9) // 10% discount
		stepsExecuted := ctx.Get("stepsExecuted").([]string)
		ctx.Set("stepsExecuted", append(stepsExecuted, "ApplyDiscount"))
		return ctx, nil
	}).SetName("ApplyDiscount")
}

// NewStepAddShipping creates a step that adds shipping cost
func NewStepAddShipping() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) (steps.StepContextInterface, error) {
		if ctx == nil {
			return ctx, fmt.Errorf("context is nil")
		}
		totalAmount := ctx.Get("totalAmount").(float64)
		ctx.Set("totalAmount", totalAmount+5.0) // Fixed shipping cost
		stepsExecuted := ctx.Get("stepsExecuted").([]string)
		ctx.Set("stepsExecuted", append(stepsExecuted, "AddShipping"))
		return ctx, nil
	}).SetName("AddShipping")
}

// NewStepCalculateTax creates a step that calculates tax
func NewStepCalculateTax() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) (steps.StepContextInterface, error) {
		if ctx == nil {
			return ctx, fmt.Errorf("context is nil")
		}
		totalAmount := ctx.Get("totalAmount").(float64)
		ctx.Set("totalAmount", totalAmount*1.2) // 20% tax
		stepsExecuted := ctx.Get("stepsExecuted").([]string)
		ctx.Set("stepsExecuted", append(stepsExecuted, "CalculateTax"))
		return ctx, nil
	}).SetName("CalculateTax")
}

// NewDag creates a DAG with conditional logic
func NewDag(orderType string, totalAmount float64) (steps.DagInterface, error) {
	dag := steps.NewDag()

	// Create steps
	processOrder := NewStepProcessOrder()
	applyDiscount := NewStepApplyDiscount()
	addShipping := NewStepAddShipping()
	calculateTax := NewStepCalculateTax()

	// Add steps to DAG
	dag.AddStep(processOrder)
	dag.AddStep(applyDiscount)
	dag.AddStep(addShipping)
	dag.AddStep(calculateTax)

	// Set up dependencies
	dag.AddDependency(applyDiscount, processOrder)
	dag.AddDependencyIf(addShipping, applyDiscount, func(ctx steps.StepContextInterface) bool {
		if ctx == nil {
			return false
		}
		orderType, ok := ctx.Get("orderType").(string)
		if !ok {
			return false
		}
		return orderType != "digital" && orderType != "subscription"
	})

	dag.AddDependencyIf(calculateTax, addShipping, func(ctx steps.StepContextInterface) bool {
		if ctx == nil {
			return false
		}
		orderType, ok := ctx.Get("orderType").(string)
		if !ok {
			return false
		}
		return orderType != "subscription"
	})

	return dag, nil
}

// NewConditionalDag creates a DAG with conditional logic
func NewConditionalDag(orderType string, totalAmount float64) steps.DagInterface {
	dag, err := NewDag(orderType, totalAmount)
	if err != nil {
		panic(err)
	}
	return dag
}
