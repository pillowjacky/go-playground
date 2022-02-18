package route

import "github.com/go-chi/chi/v5"

func ProvideRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/red")
	})

	return r
}
