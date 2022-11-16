package create_project_test

import (
	"gravity/internal/app/create_project"
	"testing"

	"github.com/lucasmmo/gravity-sdk/event"
)

const EVENT_NAME = "create_projects_test"

type testGitRepoClient struct{}

func NewTestGitRepoClient() *testGitRepoClient { return &testGitRepoClient{} }

func (t *testGitRepoClient) CreateRepository(token, name, description, owner, boilerplate string) ([]string, error) {
	return []string{"https://git", "https://git.front", "ssh.git"}, nil
}

type testGitWebhookClient struct{}

func NewTestGitWebhookClient() *testGitWebhookClient { return &testGitWebhookClient{} }

func (t *testGitWebhookClient) CreateWebhook(token, name, description, owner, boilerplate string) error {
	return nil
}

type testCreatePipeline struct{}

func NewTestCreatePipeline() *testCreatePipeline {
	return &testCreatePipeline{}
}

func (cli *testCreatePipeline) CreatePipeline() error {
	return nil
}

func TestCreateRepositoryService(t *testing.T) {
	t.Run("should create a github repository with choosen boilerplate", func(t *testing.T) {
		input := create_project.ServiceInput{
			Name:        "testing",
			Owner:       "lucasmmo",
			Boilerplate: "go-web",
		}

		gitRepoClient := NewTestGitRepoClient()
		gitWebhookClient := NewTestGitWebhookClient()
		pipeClient := NewTestCreatePipeline()

		sut := create_project.NewService(pipeClient, gitRepoClient, gitWebhookClient, event.NewDispatcher())
		output := sut.Execute("hellomooto", input)
		if output.Error != nil {
			t.Error(output.Error.Error())
		}

		if output.Url == "" {
			t.Error("url cannot be empty")
		}

		t.Log(output)
	})
}
