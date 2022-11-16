package middleware

import (
	"gravity/internal/pkg/http_helper"
	"net/http"
)

func Token(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if r.URL.Path != "/health" && r.URL.Path != "/deploy" {
			if token == "" {
				http_helper.JsonResponse(http.StatusUnauthorized, w, nil)
				return
			}
		}

		h.ServeHTTP(w, r)
	})
}
