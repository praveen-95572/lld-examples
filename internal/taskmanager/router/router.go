package router

import (
	"net/http"

	"lld-examples/internal/taskmanager/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(h *handler.TaskHandler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", h.CreateTask)
		r.Get("/", h.ListTasks)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", h.GetTask)
			r.Put("/status", h.UpdateStatus)
			r.Post("/comments", h.AddComment)
		})
	})

	return r
}
