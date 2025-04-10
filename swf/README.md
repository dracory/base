# Simple Workflow Package (SWF)

This package provides a simple, linear workflow management system in Go. It is designed for straightforward, sequential workflows where steps are executed one after another in a predefined order.

## Overview

The workflow package allows you to create and manage linear, sequential workflows. It provides functionality to:

- Create and manage steps in a sequential workflow
- Track the current step
- Determine if steps are complete
- Calculate workflow progress
- Store and retrieve metadata for steps
- Serialize and deserialize workflow state

> **Note**: This is a simple linear workflow system. It does not support complex workflows with branching paths or DAGs (Directed Acyclic Graphs). Each step follows the previous one in a straightforward sequence.

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
    // Create a new linear workflow
    wf := workflow.NewWorkflow()

    // Create steps in sequence
    step1 := workflow.NewStep("step1")
    step1.Title = "First Step"
    step1.Description = "This is the first step of the workflow"

    step2 := workflow.NewStep("step2")
    step2.Title = "Second Step"
    step2.Description = "This is the second step of the workflow"

    // Add steps to the workflow in sequence
    wf.AddStep(step1)
    wf.AddStep(step2)

    // Steps will be executed in the order they were added
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

## When to Use This Package

This package is ideal for:

- Simple approval workflows
- Sequential document processing
- Step-by-step form completion
- Linear business processes

It is not suitable for:

- Complex workflows with branching paths
- DAG-based workflows
- Parallel processing workflows
- Workflows requiring conditional branching
