package list_project

import (
	"gravity/internal/pkg/http_helper"
	"net/http"
)

type Controller struct {
	svc *service
}

func NewController(svc *service) *Controller {
	return &Controller{svc}
}

func (ctrl *Controller) Handle(w http.ResponseWriter, r *http.Request) {

	token := r.Header.Get("X-Session-Token")

	output := ctrl.svc.Execute(token)

	http_helper.JsonResponse(http.StatusOK, w, map[string]interface{}{"projects": output.Projects})
}
