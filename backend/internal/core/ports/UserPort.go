package ports

import "github.com/go-chi/chi"

type UserPort interface {
	CreateRoutes(r *chi.Mux)
	InitRoutes(r *chi.Mux)
}
