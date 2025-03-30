package main

import (
	"testing"
)

func TestDependencies(t *testing.T) {
	// Create and run the context
	dag := NewDag()
	ctx := NewExampleContext()
	if err := dag.Run(ctx); err != nil {
		t.Fatal(err)
	}

	// Verify final price calculation
	// Base price: 100
	// After 20% discount: 80
	// Add shipping: 90
	// Add 20% tax: 108
	if ctx.finalPrice != 108 {
		t.Fatalf("expected final price 108, got: %d", ctx.finalPrice)
	}

	// Verify steps completed in correct order
	if len(ctx.stepsCompleted) != 4 ||
		ctx.stepsCompleted[0] != "SetBasePrice" ||
		ctx.stepsCompleted[1] != "ApplyDiscount" ||
		ctx.stepsCompleted[2] != "AddShipping" ||
		ctx.stepsCompleted[3] != "CalculateTax" {
		t.Fatalf("steps completed in wrong order: %v", ctx.stepsCompleted)
	}

	// Verify base price is set correctly
	if ctx.basePrice != 100 {
		t.Fatalf("expected base price 100, got: %d", ctx.basePrice)
	}
}
