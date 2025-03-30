package basic_usage

import (
	"github.com/dracory/base/object"
	"github.com/dracory/base/steps"
)

// ExampleContext implements StepContextInterface
type ExampleContext struct {
	*object.SerializablePropertyObject
}

// NewExampleContext creates a new ExampleContext
func NewExampleContext() *ExampleContext {
	return &ExampleContext{
		SerializablePropertyObject: object.NewSerializablePropertyObject().(*object.SerializablePropertyObject),
	}
}

func (c *ExampleContext) Name() string {
	return c.Get("name").(string)
}

func (c *ExampleContext) SetName(name string) steps.StepContextInterface {
	c.Set("name", name)
	return c
}

// NewIncrementStep creates a new step that increments a value
func NewIncrementStep() steps.StepInterface {
	return steps.NewStep(func(ctx steps.StepContextInterface) error {
		value := ctx.Get("value").(int)
		value++
		ctx.Set("value", value)
		return nil
	})
}

// NewDag creates a new DAG with 5 increment steps
func NewDag() steps.DagInterface {
	dag := steps.NewDag()
	dag.AddStep(NewIncrementStep())
	dag.AddStep(NewIncrementStep())
	dag.AddStep(NewIncrementStep())
	dag.AddStep(NewIncrementStep())
	dag.AddStep(NewIncrementStep())
	return dag
}
