package main

import (
	"fmt"
)

func main() {
	// Create and run the context
	ctx := NewExampleContext()
	dag := NewDag()
	if err := dag.Run(ctx); err != nil {
		fmt.Printf("Error running DAG: %v\n", err)
		return
	}

	fmt.Printf("Final price: %d\n", ctx.finalPrice)
	fmt.Printf("Steps completed: %v\n", ctx.stepsCompleted)
}
