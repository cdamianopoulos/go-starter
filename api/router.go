package api

import (
	"go-starter/api/zoo"
	"go-starter/separateRepos/healthcheck"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Router returns a chi Mux server that implements http.Handler.
func Router() (r *chi.Mux) {
	r = chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get(healthcheck.URLPattern, healthcheck.Status)
	r.Get(zoo.URLPattern, zoo.Status)

	return
}
