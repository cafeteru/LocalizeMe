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

type GroupPortImpl struct {
	controller controller.GroupController
}

func CreateGroupPort() *GroupPortImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	groupController := impl.CreateGroupController()
	port := &GroupPortImpl{groupController}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return port
}

func (g GroupPortImpl) InitRoutes(r *chi.Mux) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	pattern := "/" + constants.Groups
	tokenAuth := utils.ConfigJWTRoutes()
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Route(pattern, func(r chi.Router) {
			r.Post("/", g.controller.Create)
			r.Get("/", g.controller.FindAll)
			r.Get("/canWrite", g.controller.FindCanWrite)
			r.Put("/", g.controller.Update)
			r.Patch("/{id}", g.controller.Disable)
			r.Delete("/{id}", g.controller.Delete)
		})
	})
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
