package ports

import "github.com/go-chi/chi"

type XliffPort interface {
	CreateRoutes(r *chi.Mux)
	InitRoutes(r *chi.Mux)
}
