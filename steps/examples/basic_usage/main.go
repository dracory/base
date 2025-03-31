package basic_usage

import (
	"fmt"
	"github.com/dracory/base/steps"
)

func main() {
	// Create and run the DAG
	ctx := steps.NewStepContext()
	ctx, err := steps.NewDag().Run(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Access the value from the context
	value := ctx.Get("value")
	fmt.Printf("Value: %v\n", value)
}
