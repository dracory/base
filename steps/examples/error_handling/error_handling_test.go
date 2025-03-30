package main

import (
	"testing"
)

func TestErrorHandling(t *testing.T) {
	dag := NewDag()
	ctx := NewExampleContext()
	
	if err := dag.Run(ctx); err == nil {
		t.Error("Expected error, got nil")
	}
	
	if ctx.value != 200 {
		t.Errorf("Expected value 200, got: %d", ctx.value)
	}
	
	// Verify steps completed in correct order
	if len(ctx.stepsCompleted) != 3 || 
		ctx.stepsCompleted[0] != "SetInitialValue" || 
		ctx.stepsCompleted[1] != "ProcessData" || 
		ctx.stepsCompleted[2] != "VerifyData" {
		t.Errorf("Steps completed in wrong order: %v", ctx.stepsCompleted)
	}
	
	if ctx.errorCount != 1 {
		t.Errorf("Expected 1 error, got: %d", ctx.errorCount)
	}
}
