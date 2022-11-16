package create_project

import (
	"encoding/json"
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
	var createRepositoryRequestBody RequestBody

	token := r.Header.Get("X-Session-Token")

	if err := json.NewDecoder(r.Body).Decode(&createRepositoryRequestBody); err != nil {
		http_helper.JsonResponse(http.StatusBadRequest, w, map[string]interface{}{"error": err.Error()})
		return
	}

	if createRepositoryRequestBody.Name == "" || createRepositoryRequestBody.Owner == "" || createRepositoryRequestBody.Boilerplate == "" {
		http_helper.JsonResponse(http.StatusBadRequest, w, map[string]interface{}{"error": "invalid data"})
		return
	}

	input := ServiceInput{
		Name:        createRepositoryRequestBody.Name,
		Description: createRepositoryRequestBody.Description,
		Owner:       createRepositoryRequestBody.Owner,
		Boilerplate: createRepositoryRequestBody.Boilerplate,
		PipelineUrl: "http://localhost:8080/pipeline",
	}

	output := ctrl.svc.Execute(token, input)
	if output.Error != nil {
		http_helper.JsonResponse(http.StatusBadRequest, w, map[string]interface{}{"error": output.Error.Error()})
		return
	}

	http_helper.JsonResponse(http.StatusCreated, w, map[string]interface{}{"message": output.Url})
}
