package ports

import "github.com/go-chi/chi"

type GroupPort interface {
	CreateRoutes(r *chi.Mux)
	InitRoutes(r *chi.Mux)
}
