package router

import (
	"github.com/Mark1708/go-pastebin/internal/health"
	"github.com/Mark1708/go-pastebin/internal/paste"
	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/health", health.Read)

	r.Route("/api/v1", func(r chi.Router) {
		pasteAPI := &paste.API{}
		r.Get("/pastes/{id}", pasteAPI.Get)
		r.Post("/pastes", pasteAPI.Create)
		r.Put("/pastes/{id}", pasteAPI.Update)
		r.Delete("/pastes/{id}", pasteAPI.Delete)
	})

	return r
}
