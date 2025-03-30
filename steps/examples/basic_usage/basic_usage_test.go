package basic_usage

import (
	"testing"
)

func TestBasicUsage(t *testing.T) {
	dag := NewDag()
	ctx := NewExampleContext()
	ctx.Set("value", 0)
	
	if err := dag.Run(ctx); err != nil {
		t.Fatal(err)
	}

	if value := ctx.Get("value").(int); value != 5 {
		t.Fatalf("expected value 5, got: %d", value)
	}
}
