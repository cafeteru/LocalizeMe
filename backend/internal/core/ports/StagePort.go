package ports

import "github.com/go-chi/chi"

type StagePort interface {
	CreateRoutes(r *chi.Mux)
	InitRoutes(r *chi.Mux)
}
