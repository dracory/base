package steps

type pipelineImplementation struct {
	id    string
	name  string
	nodes []RunnableInterface
}
