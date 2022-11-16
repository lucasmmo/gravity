package list_project

import (
	"gravity/internal/pkg/git_client"
	"gravity/internal/pkg/sql_client"
)

type service struct {
	sqlClient sql_client.ListProject
	gitClient git_client.GetUserdata
}

func NewService(sqlClient sql_client.ListProject, gitClient git_client.GetUserdata) *service {
	return &service{sqlClient, gitClient}
}

func (svc *service) Execute(token string) ServiceOutput {
	userData, err := svc.gitClient.GetUserdata(token)
	if err != nil {
		return ServiceOutput{[]string{}, err}
	}

	projects := svc.sqlClient.List(userData["login"])
	return ServiceOutput{projects, nil}
}
