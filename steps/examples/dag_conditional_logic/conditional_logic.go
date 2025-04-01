package main

import (
	"context"

	"github.com/dracory/base/steps"
)

// NewStepProcessOrder creates a step that processes the order
func NewStepProcessOrder() steps.StepInterface {
	step := steps.NewStep()
	step.SetName("ProcessOrder")
	step.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		stepsExecuted := data["stepsExecuted"].([]string)
		data["stepsExecuted"] = append(stepsExecuted, "ProcessOrder")
		return ctx, data, nil
	})
	return step
}

// NewStepApplyDiscount creates a step that applies a discount
func NewStepApplyDiscount() steps.StepInterface {
	step := steps.NewStep()
	step.SetName("ApplyDiscount")
	step.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		totalAmount := data["totalAmount"].(float64)
		data["totalAmount"] = totalAmount * 0.9 // 10% discount
		stepsExecuted := data["stepsExecuted"].([]string)
		data["stepsExecuted"] = append(stepsExecuted, "ApplyDiscount")
		return ctx, data, nil
	})
	return step
}

// NewStepAddShipping creates a step that adds shipping cost
func NewStepAddShipping() steps.StepInterface {
	step := steps.NewStep()
	step.SetName("AddShipping")
	step.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		totalAmount := data["totalAmount"].(float64)
		data["totalAmount"] = totalAmount + 5.0 // Fixed shipping cost
		stepsExecuted := data["stepsExecuted"].([]string)
		data["stepsExecuted"] = append(stepsExecuted, "AddShipping")
		return ctx, data, nil
	})
	return step
}

// NewStepCalculateTax creates a step that calculates tax
func NewStepCalculateTax() steps.StepInterface {
	step := steps.NewStep()
	step.SetName("CalculateTax")
	step.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		totalAmount := data["totalAmount"].(float64)
		data["totalAmount"] = totalAmount * 1.2 // 20% tax
		stepsExecuted := data["stepsExecuted"].([]string)
		data["stepsExecuted"] = append(stepsExecuted, "CalculateTax")
		return ctx, data, nil
	})
	return step
}

// NewConditionalDag creates a DAG with conditional logic
//
// # Depending on the order type, a different set of steps is added to the DAG
//
// On digital orders, only ProcessOrder and ApplyDiscount are added
// On physical orders, ProcessOrder, ApplyDiscount, and AddShipping are added
// On subscription orders, only ProcessOrder and ApplyDiscount are added
//
// Parameters:
// - orderType: The type of order (digital, physical, subscription)
// - totalAmount: The total amount of the order
// Returns:
// - dag: The DAG with conditional logic
// - error: Error if any
func NewConditionalDag(orderType string, totalAmount float64) (steps.DagInterface, error) {
	dag := steps.NewDag()
	dag.SetName("Conditional Logic Example DAG")

	// Create common steps
	processOrder := NewStepProcessOrder()
	applyDiscount := NewStepApplyDiscount()
	calculateTax := NewStepCalculateTax()

	// Add common steps to DAG
	dag.RunnableAdd(processOrder, applyDiscount, calculateTax)

	// Set up common dependencies
	dag.DependencyAdd(applyDiscount, processOrder)

	// Add shipping step and dependencies only for physical orders
	if orderType == "physical" {
		addShipping := NewStepAddShipping()
		dag.RunnableAdd(addShipping)
		dag.DependencyAdd(addShipping, applyDiscount)
		dag.DependencyAdd(calculateTax, addShipping)
	} else {
		dag.DependencyAdd(calculateTax, applyDiscount)
	}

	return dag, nil
}

// RunConditionalExample runs the conditional logic example
func RunConditionalExample(orderType string, totalAmount float64) (map[string]any, error) {
	dag, err := NewConditionalDag(orderType, totalAmount)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	data := map[string]any{
		"orderType":     orderType,
		"totalAmount":   totalAmount,
		"stepsExecuted": []string{},
	}
	_, data, err = dag.Run(ctx, data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
