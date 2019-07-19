package routing

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// New - create a new instance of a Chi Router
func New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Heartbeat("/health"))
	r.Use(middleware.URLFormat)
	r.Use(middleware.Recoverer)

	return r
}
