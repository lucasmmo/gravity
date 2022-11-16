package main

import (
	"gravity/internal/app/config_server"
	"gravity/internal/app/create_project"
	"gravity/internal/app/list_project"
	"gravity/internal/app/store_project"
	"gravity/internal/pkg/git_client"
	"gravity/internal/pkg/pipeline_client"
	"gravity/internal/pkg/sql_client"
	"log"
	"net/http"

	"github.com/lucasmmo/gravity-sdk/event"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dependencies struct {
	configServerController     *config_server.Controller
	createRepositoryController *create_project.Controller
	listProjectsController     *list_project.Controller
}

func InitDependencies(dsn string) *dependencies {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	httpClient := &http.Client{}
	githubClient := git_client.NewGithubClient(httpClient)
	sqlProjectClient := sql_client.NewProjectRepository(db)
	jenkinsClient := pipeline_client.NewJenkinsPipeline(httpClient)

	configServerCtrl := config_server.NewController()

	listProjectsSvc := list_project.NewService(sqlProjectClient, githubClient)
	listProjectsCtrl := list_project.NewController(listProjectsSvc)

	createProjectDsp := event.NewDispatcher()
	createProjectSvc := create_project.NewService(jenkinsClient, githubClient, githubClient, createProjectDsp)
	createProjectCtrl := create_project.NewController(createProjectSvc)

	storeProjectSvc := store_project.NewService(sqlProjectClient, githubClient)
	storeProjectLst := store_project.NewListener(storeProjectSvc)

	createProjectDsp.AddListener("create_project", storeProjectLst)

	return &dependencies{
		configServerController:     configServerCtrl,
		createRepositoryController: createProjectCtrl,
		listProjectsController:     listProjectsCtrl,
	}
}
