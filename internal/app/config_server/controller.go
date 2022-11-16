package config_server

import (
	"gravity/internal/pkg/http_helper"
	"net/http"
)

type Controller struct{}

func NewController() *Controller { return &Controller{} }

func (ctrl *Controller) Handle(w http.ResponseWriter, r *http.Request) {
	http_helper.JsonResponse(http.StatusOK, w, nil)
}
