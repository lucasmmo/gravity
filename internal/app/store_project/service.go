package store_project

import (
	"gravity/internal/pkg/git_client"
	"gravity/internal/pkg/sql_client"
)

const EVENT_NAME = "create_repository"

type service struct {
	sqlClient sql_client.StoreProject
	gitClient git_client.GetUserdata
}

func NewService(sqlClient sql_client.StoreProject, gitClient git_client.GetUserdata) *service {
	return &service{sqlClient, gitClient}
}

func (svc *service) Execute(token, projectName string) error {
	userData, err := svc.gitClient.GetUserdata(token)
	if err != nil {
		return err
	}

	return svc.sqlClient.Store(projectName, userData["login"])
}
