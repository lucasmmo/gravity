package create_project

import (
	"encoding/json"
	"gravity/internal/pkg/git_client"
	"gravity/internal/pkg/pipeline_client"

	"github.com/lucasmmo/gravity-sdk/event"
)

const EVENT_NAME = "create_project"
const WEBHOOK_NAME = "gravity_cd"

type service struct {
	pipelineClient   pipeline_client.CreatePipeline
	gitRepoClient    git_client.CreateRepository
	gitWebhookClient git_client.CreateWebhook
	dispatcher       event.Dispatcher
}

func NewService(pipelineClient pipeline_client.CreatePipeline, gitRepoClient git_client.CreateRepository, gitWebhookClient git_client.CreateWebhook, dispatcher event.Dispatcher) *service {
	return &service{pipelineClient, gitRepoClient, gitWebhookClient, dispatcher}
}

func (srv *service) Execute(token string, input ServiceInput) ServiceOutput {

	url, err := srv.gitRepoClient.CreateRepository(token, input.Name, input.Description, input.Owner, input.Boilerplate)
	if err != nil {
		return ServiceOutput{
			Error: err,
		}
	}

	if err := srv.pipelineClient.CreatePipeline(); err != nil {
		return ServiceOutput{
			Error: err,
		}
	}

	if err := srv.gitWebhookClient.CreateWebhook(token, input.Owner, input.Name, input.PipelineUrl, WEBHOOK_NAME); err != nil {
		return ServiceOutput{
			Error: err,
		}
	}

	data, err := json.Marshal(&EventData{
		Token:       token,
		ProjectName: input.Name,
	})
	if err != nil {
		return ServiceOutput{
			Error: err,
		}
	}

	srv.dispatcher.Dispatch(&Event{
		key:  EVENT_NAME,
		data: data,
	})

	return ServiceOutput{
		Url:   url[2],
		Error: nil,
	}
}
