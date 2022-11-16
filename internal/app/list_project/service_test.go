package list_project_test

import (
	"gravity/internal/app/list_project"
	"testing"
)

type TestSqlClient struct{}

func NewTestSqlClient() *TestSqlClient { return &TestSqlClient{} }

func (t *TestSqlClient) List(owner string) []string {
	return []string{"my-test-project"}
}

type TestGitClient struct{}

func NewTestGitClient() *TestGitClient { return &TestGitClient{} }

func (t *TestGitClient) GetUserdata(token string) (map[string]string, error) {
	return map[string]string{"owner": "test-user"}, nil
}
func TestListProject(t *testing.T) {
	t.Run("", func(t *testing.T) {
		sqlClient := NewTestSqlClient()
		gitClient := NewTestGitClient()

		sut := list_project.NewService(sqlClient, gitClient)

		res := sut.Execute("test-user")
		if len(res.Projects) == 0 {
			t.Error("length of result cannot be 0")
		}
	})
}
