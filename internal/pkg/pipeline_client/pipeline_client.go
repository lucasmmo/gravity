package pipeline_client

import "net/http"

type CreatePipeline interface {
	CreatePipeline() error
}

type jenkinsPipeline struct {
	client *http.Client
}

func NewJenkinsPipeline(client *http.Client) *jenkinsPipeline {
	return &jenkinsPipeline{client}
}

func (cli *jenkinsPipeline) CreatePipeline() error {

	return nil
}
