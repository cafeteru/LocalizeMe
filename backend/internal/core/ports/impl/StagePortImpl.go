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

type StagePortImpl struct {
	controller controller.StageController
}

func CreateStagePort() *StagePortImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	stageController := impl.CreateStageController()
	port := &StagePortImpl{stageController}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return port
}

func (s StagePortImpl) InitRoutes(r *chi.Mux) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	pattern := "/" + constants.Stages
	tokenAuth := utils.ConfigJWTRoutes()
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Route(pattern, func(r chi.Router) {
			r.Post("/", s.controller.Create)
			r.Get("/", s.controller.FindAll)
			r.Get("/name/{name}", s.controller.FindByName)
			r.Put("/", s.controller.Update)
			r.Patch("/{id}", s.controller.Disable)
			r.Delete("/{id}", s.controller.Delete)
		})
	})
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
