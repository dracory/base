package main

import (
	"fmt"
)

func main() {
	// Create and run the context
	ctx := NewExampleContext()

	resultCtx, err := NewDag().Run(ctx)

	if err != nil {
		fmt.Printf("Error running DAG: %v\n", err)
		return
	}

	// Access values from the context
	finalPrice := resultCtx.(*ExampleContext).finalPrice
	stepsCompleted := resultCtx.(*ExampleContext).stepsCompleted

	fmt.Printf("Final price: %d\n", finalPrice)
	fmt.Printf("Steps completed: %v\n", stepsCompleted)
}
