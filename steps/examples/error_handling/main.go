package main

import (
	"fmt"
)

func main() {
	// Create and run the context
	ctx := NewExampleContext()
	dag := NewDag()

	result, err := dag.Run(ctx)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("Final value: %d\n", result.(*ExampleContext).value)
	fmt.Printf("Steps completed: %v\n", result.(*ExampleContext).stepsCompleted)
	fmt.Printf("Error count: %d\n", result.(*ExampleContext).errorCount)
}
