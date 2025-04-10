# Simple Workflow Package (SWF)

This package provides a simple workflow management system in Go.

## Overview

The workflow package allows you to create and manage multi-step workflows. It provides functionality to:

- Create and manage steps in a workflow
- Track the current step
- Determine if steps are complete
- Calculate workflow progress
- Store and retrieve metadata for steps
- Serialize and deserialize workflow state

## Components

### Step

A `Step` represents a single step in a workflow. Each step has:

- `Name`: Unique identifier for the step
- `Type`: Type of the step (default: "normal")
- `Title`: Display title for the step
- `Description`: Description of what the step does
- `Responsible`: Person or role responsible for the step

### Workflow

A `Workflow` manages multiple steps and tracks the workflow state. It provides methods to:

- Add steps to the workflow
- Get and set the current step
- Check if a step is current or complete
- Calculate workflow progress
- Store and retrieve metadata for steps
- Serialize and deserialize workflow state

## Usage

```go
package main

import (
    "fmt"
    "github.com/yourusername/workflow"
)

func main() {
    // Create a new workflow
    wf := workflow.NewWorkflow()

    // Create steps
    step1 := workflow.NewStep("step1")
    step1.Title = "First Step"
    step1.Description = "This is the first step of the workflow"

    step2 := workflow.NewStep("step2")
    step2.Title = "Second Step"
    step2.Description = "This is the second step of the workflow"

    // Add steps to the workflow
    wf.AddStep(step1)
    wf.AddStep(step2)

    // Get the current step
    currentStep := wf.GetCurrentStep()
    fmt.Printf("Current step: %s\n", currentStep.Name)

    // Move to the next step
    wf.SetCurrentStep(step2)

    // Check if a step is complete
    isComplete := wf.IsStepComplete(step1)
    fmt.Printf("Is step1 complete? %v\n", isComplete)

    // Get progress
    progress := wf.GetProgress()
    fmt.Printf("Progress: %d/%d steps completed (%.2f%%)\n",
        progress.Completed, progress.Total, progress.Percents)
}
```

## Differences from PHP Implementation

This Go implementation has some differences from the original PHP implementation:

1. The `GetActionLink()` method in the `Step` struct is a simplified implementation, as the PHP version uses a framework-specific function.
2. The `GetState()` method now returns the actual state struct instead of accessing a non-existent `memory` property.
3. Error handling is more explicit in the Go version, with methods returning errors where appropriate.
4. The Go version uses strong typing and interfaces to handle different parameter types.

## License

This code is a conversion of the Sinevia Workflow library. Please refer to the original license for usage terms.
