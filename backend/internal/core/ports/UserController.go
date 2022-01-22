package ports

import "github.com/go-chi/chi"

type UserController interface {
	CreateUserRoutes(r *chi.Mux)
}
