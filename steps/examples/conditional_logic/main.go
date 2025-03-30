package main

import (
	"fmt"
)

func main() {
	// Create context
	ctx := NewOrderContext("standard")

	// Create and run DAG
	dag := NewConditionalDag("standard")
	if err := dag.Run(ctx); err != nil {
		fmt.Printf("Error running DAG: %v\n", err)
		return
	}

	// Print results
	fmt.Printf("Total amount: %.2f\n", ctx.Get("totalAmount").(float64))
	fmt.Println("Steps executed:", ctx.stepsExecuted)
}
