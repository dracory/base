package main

import (
	"fmt"
)

func main() {
	// Create context
	ctx := NewOrderContext("standard", 100.0)

	// Create and run DAG
	dag := NewConditionalDag("standard", 100.0)
	ctx, err := dag.Run(ctx)
	if err != nil {
		fmt.Printf("Error running DAG: %v\n", err)
		return
	}

	// Print results
	fmt.Printf("Total amount: %.2f\n", ctx.Get("totalAmount").(float64))
	fmt.Println("Steps executed:", ctx.Get("stepsExecuted").([]string))
}
