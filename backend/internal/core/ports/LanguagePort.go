package ports

import "github.com/go-chi/chi"

type LanguagePort interface {
	CreateRoutes(r *chi.Mux)
	InitRoutes(r *chi.Mux)
}
