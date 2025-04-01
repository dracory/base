package main

import (
	"context"
	"fmt"
)

func main() {
	// Create and run DAG
	dag, err := NewConditionalDag("standard", 100.0)
	if err != nil {
		fmt.Printf("Error creating DAG: %v\n", err)
		return
	}

	_, data, err := dag.Run(context.Background(), map[string]any{
		"orderType":     "standard",
		"totalAmount":   100.0,
		"stepsExecuted": []string{},
	})
	if err != nil {
		fmt.Printf("Error running DAG: %v\n", err)
		return
	}

	// Print results
	fmt.Printf("Total amount: %.2f\n", data["totalAmount"].(float64))
	fmt.Println("Steps executed:", data["stepsExecuted"].([]string))
}
