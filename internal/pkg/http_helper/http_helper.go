package http_helper

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(statusCode int, w http.ResponseWriter, body interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}
