package main

import (
	"fmt"
)

func main() {
	// Create and run the context
	ctx := NewExampleContext()
	dag := NewDag()
	
	if err := dag.Run(ctx); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	
	fmt.Printf("Final value: %d\n", ctx.value)
	fmt.Printf("Steps completed: %v\n", ctx.stepsCompleted)
	fmt.Printf("Error count: %d\n", ctx.errorCount)
}
