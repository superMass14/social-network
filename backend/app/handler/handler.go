package handler

import (
	"backend/app/cors"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	cors.SetCors(&w)
	// ...
}

func RequestValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cors.SetCors(&w)
		next.ServeHTTP(w, r)
	}
}
