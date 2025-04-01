package wf

import (
	"context"

	"github.com/dracory/base/arr"
)

type pipelineImplementation struct {
	id    string
	name  string
	nodes []RunnableInterface
}

func NewPipeline() PipelineInterface {
	return &pipelineImplementation{
		nodes: make([]RunnableInterface, 0),
	}
}

var _ PipelineInterface = (*pipelineImplementation)(nil)

func (p *pipelineImplementation) GetID() string {
	return p.id
}

func (p *pipelineImplementation) SetID(id string) {
	p.id = id
}

func (p *pipelineImplementation) GetName() string {
	return p.name
}

func (p *pipelineImplementation) SetName(name string) {
	p.name = name
}

func (p *pipelineImplementation) Run(ctx context.Context, data map[string]any) (context.Context, map[string]any, error) {
	for _, node := range p.nodes {
		ctx, data, err := node.Run(ctx, data)
		if err != nil {
			return ctx, data, err
		}
	}
	return ctx, data, nil
}

func (p *pipelineImplementation) RunnableAdd(node ...RunnableInterface) {
	p.nodes = append(p.nodes, node...)
}

func (p *pipelineImplementation) RunnableRemove(node RunnableInterface) bool {
	id := node.GetID()

	if id == "" {
		return false
	}

	for i, n := range p.nodes {
		if n.GetID() == id {
			p.nodes = arr.IndexRemove(p.nodes, i)
			return true
		}
	}

	return false
}

func (p *pipelineImplementation) RunnableList() []RunnableInterface {
	return p.nodes
}
