package main

import (
	"gravity/cmd/api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(app *mux.Router, dep *dependencies) {
	app.Use(middleware.Token)
	app.HandleFunc("/health", dep.configServerController.Handle).Methods(http.MethodGet)
	app.HandleFunc("/project/create", dep.createRepositoryController.Handle).Methods(http.MethodPost)
	app.HandleFunc("/projects", dep.listProjectsController.Handle).Methods(http.MethodGet)
	// app.HandleFunc("/deploy", dep.makeDeployController.Handle).Methods(http.MethodPost)
}
