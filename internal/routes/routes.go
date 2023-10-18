package routes

import (
	"net/http"

	"github.com/eduardo-camarena/caffeineBlues/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() http.Handler {
	fileServer := http.FileServer(http.Dir("web/static"))
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	mux.Get("/", services.IndexHandler)
	mux.Post("/brews/new", services.NewBrew)

	return mux
}
