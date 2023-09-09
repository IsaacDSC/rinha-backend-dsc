package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HealthcheckController(router *chi.Mux) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("A live"))
	})
}
