package main

import (
	"context"
	"fmt"
)

func main() {
	// Create and run DAG
	dag := NewConditionalDag("standard", 100.0)
	ctx := context.Background()
	data := map[string]any{
		"orderType":      "standard",
		"totalAmount":    100.0,
		"stepsExecuted": []string{},
	}
	_, data, err := dag.Run(ctx, data)
	if err != nil {
		fmt.Printf("Error running DAG: %v\n", err)
		return
	}

	// Print results
	fmt.Printf("Total amount: %.2f\n", data["totalAmount"].(float64))
	fmt.Println("Steps executed:", data["stepsExecuted"].([]string))
}
