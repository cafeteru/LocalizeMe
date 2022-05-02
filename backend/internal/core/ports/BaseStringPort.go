package ports

import "github.com/go-chi/chi"

type BaseStringPort interface {
	CreateRoutes(r *chi.Mux)
	InitRoutes(r *chi.Mux)
}
