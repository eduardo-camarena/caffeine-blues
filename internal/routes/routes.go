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

	mux.Get("/", useBaseHtml(services.IndexHandler))
	mux.Post("/brews/new", services.NewBrew)

	return mux
}

func useBaseHtml(
	f func(w http.ResponseWriter, r *http.Request, htmlTemplates []string),
) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var htmlTemplates []string
		_, ok := r.Header["Hx-Request"]

		if ok == false {
			htmlTemplates = append(htmlTemplates, "web/index.html", "web/newBrew.html")
		}

		f(w, r, htmlTemplates)
	}
}
