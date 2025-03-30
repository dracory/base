package basic_usage

import (
	"fmt"
)

func main() {
	ctx := NewExampleContext()
	ctx.Set("value", 0)

	dag := NewDag()

	if err := dag.Run(ctx); err != nil {
		fmt.Printf("Error running DAG: %v\n", err)
		return
	}

	fmt.Printf("Final value: %d\n", ctx.Get("value"))
}
