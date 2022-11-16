package store_project_test

import (
	"gravity/internal/app/store_project"
	"testing"
)

const TEST_TOKEN = "IUY!@#UG!@IOYU#G!"

type TestSqlClient struct {
	data map[string]interface{}
}

func NewTestSqlClient() *TestSqlClient {
	return &TestSqlClient{
		data: make(map[string]interface{}),
	}
}

func (t *TestSqlClient) Store(name, owner string) error {
	t.data["name"] = name
	return nil
}

type TestGitClient struct{}

func NewTestGitClient() *TestGitClient { return &TestGitClient{} }

func (t *TestGitClient) GetUserdata(token string) (map[string]string, error) {
	return map[string]string{"owner": "test-user"}, nil
}

func TestCreateScopeTest(t *testing.T) {
	t.Run("should create a new scope", func(t *testing.T) {
		input := store_project.ServiceInput{
			ProjectName: "TESTING PROJECT",
		}

		sqlClient := NewTestSqlClient()
		gitClient := NewTestGitClient()

		sut := store_project.NewService(sqlClient, gitClient)

		if err := sut.Execute(TEST_TOKEN, input.ProjectName); err != nil {
			t.Error(err.Error())
		}
	})
}
