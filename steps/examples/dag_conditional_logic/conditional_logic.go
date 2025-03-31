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
func NewConditionalDag(orderType string, totalAmount float64) steps.DagInterface {
	dag := steps.NewDag()
	dag.SetName("Conditional Logic Example DAG")

	// Create steps
	processOrder := NewStepProcessOrder()
	applyDiscount := NewStepApplyDiscount()
	addShipping := NewStepAddShipping()
	calculateTax := NewStepCalculateTax()

	// Add steps to DAG
	dag.RunnableAdd(processOrder, applyDiscount, addShipping, calculateTax)

	// Set up dependencies
	dag.DependencyAdd(applyDiscount, processOrder)
	dag.DependencyAddIf(addShipping, applyDiscount, func(ctx context.Context, data map[string]any) bool {
		return data["orderType"] != "digital" && data["orderType"] != "subscription"
	})
	dag.DependencyAddIf(calculateTax, addShipping, func(ctx context.Context, data map[string]any) bool {
		return data["orderType"] != "subscription"
	})
	dag.DependencyAddIf(calculateTax, applyDiscount, func(ctx context.Context, data map[string]any) bool {
		return data["orderType"] == "digital" && data["orderType"] != "subscription"
	})

	return dag
}

// RunConditionalExample runs the conditional logic example
func RunConditionalExample(orderType string, totalAmount float64) (map[string]any, error) {
	dag := NewConditionalDag(orderType, totalAmount)
	ctx := context.Background()
	data := map[string]any{
		"orderType":      orderType,
		"totalAmount":    totalAmount,
		"stepsExecuted": []string{},
	}
	_, data, err := dag.Run(ctx, data)
	return data, err
}
