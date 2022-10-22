package impl

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"log"
	"uniovi-localizeme/constants"
	"uniovi-localizeme/internal/core/ports/controller"
	"uniovi-localizeme/internal/core/ports/controller/impl"
	"uniovi-localizeme/internal/core/ports/utils"
	"uniovi-localizeme/tools"
)

type UserPortImpl struct {
	controller controller.UserController
}

func CreateUserPort() *UserPortImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	userController := impl.CreateUserController()
	port := &UserPortImpl{userController}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return port
}

func (u UserPortImpl) InitRoutes(r *chi.Mux) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	pattern := "/" + constants.Users
	tokenAuth := utils.ConfigJWTRoutes()
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Route(pattern, func(r chi.Router) {
			r.Get("/", u.controller.FindAll)
			r.Get("/me", u.controller.FindMe)
			r.Get("/id/{id}", u.controller.FindById)
			r.Put("/", u.controller.Update)
			r.Put("/me", u.controller.UpdateMe)
			r.Patch("/{id}", u.controller.Disable)
			r.Delete("/{id}", u.controller.Delete)
		})
	})
	r.Group(func(r chi.Router) {
		r.Post("/login", u.controller.Login)
		r.Post(pattern, u.controller.Create)
	})
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
