package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dracory/base/wf"
)

// NewSendEmailStep creates a step that sends a verification email
func NewSendEmailStep() wf.StepInterface {
	step := wf.NewStep()
	step.SetName("Send Verification Email")
	step.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		email := data["email"].(string)
		code := generateVerificationCode()
		data["verificationCode"] = code

		// In a real application, this would send an actual email
		fmt.Printf("Sending verification code %s to %s\n", code, email)

		return ctx, data, nil
	})
	return step
}

// NewWaitForVerificationStep creates a step that waits for user input
func NewWaitForVerificationStep() wf.StepInterface {
	step := wf.NewStep()
	step.SetName("Wait for Verification")
	step.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		// In a real application, this would wait for user input
		// For this example, we'll simulate waiting by sleeping
		time.Sleep(2 * time.Second)

		// Simulate user entering the code
		enteredCode := data["verificationCode"].(string)
		data["enteredCode"] = enteredCode

		return ctx, data, nil
	})
	return step
}

// NewVerifyCodeStep creates a step that verifies the entered code
func NewVerifyCodeStep() wf.StepInterface {
	step := wf.NewStep()
	step.SetName("Verify Code")
	step.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		expectedCode := data["verificationCode"].(string)
		enteredCode := data["enteredCode"].(string)

		if expectedCode != enteredCode {
			return ctx, data, fmt.Errorf("invalid verification code")
		}

		data["verified"] = true
		return ctx, data, nil
	})
	return step
}

// NewCompleteStep creates a step that completes the workflow
func NewCompleteStep() wf.StepInterface {
	step := wf.NewStep()
	step.SetName("Complete")
	step.SetHandler(func(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
		fmt.Println("Email verification completed successfully!")
		return ctx, data, nil
	})
	return step
}

// generateVerificationCode generates a random 6-digit code
func generateVerificationCode() string {
	return fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)
}

// NewEmailVerificationWorkflow creates a workflow for email verification
func NewEmailVerificationWorkflow() wf.DagInterface {
	dag := wf.NewDag()
	dag.SetName("Email Verification Workflow")

	// Create steps
	sendEmail := NewSendEmailStep()
	waitForVerification := NewWaitForVerificationStep()
	verifyCode := NewVerifyCodeStep()
	complete := NewCompleteStep()

	// Add steps to DAG
	dag.RunnableAdd(sendEmail, waitForVerification, verifyCode, complete)

	// Set up dependencies
	dag.DependencyAdd(waitForVerification, sendEmail)
	dag.DependencyAdd(verifyCode, waitForVerification)
	dag.DependencyAdd(complete, verifyCode)

	return dag
}

// RunEmailVerificationExample demonstrates the email verification workflow
func RunEmailVerificationExample() error {
	// Create workflow
	dag := NewEmailVerificationWorkflow()

	// Initialize data
	ctx := context.Background()
	data := map[string]any{
		"email": "user@example.com",
	}

	// Start workflow
	ctx, data, err := dag.Run(ctx, data)
	if err != nil {
		return fmt.Errorf("workflow failed: %v", err)
	}

	// Pause the workflow after sending email
	if dag.GetState().GetStatus() == wf.StateStatus(wf.StateStatusRunning) {
		err = dag.Pause()
		if err != nil {
			return fmt.Errorf("failed to pause workflow: %v", err)
		}
		fmt.Println("Workflow paused after sending email")
	}

	// Save workflow state
	state := dag.GetState()
	stateJSON, err := state.ToJSON()
	if err != nil {
		return fmt.Errorf("failed to save state: %v", err)
	}
	fmt.Printf("Saved workflow state: %s\n", string(stateJSON))

	// Create a new workflow instance
	newDag := NewEmailVerificationWorkflow()

	// Load saved state
	newState := wf.NewState()
	if err := newState.FromJSON(stateJSON); err != nil {
		return fmt.Errorf("failed to load state: %v", err)
	}
	newDag.SetState(newState)

	// Resume workflow
	ctx, data, err = newDag.Resume(ctx, data)
	if err != nil {
		return fmt.Errorf("workflow resume failed: %v", err)
	}

	return nil
}
