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

type BaseStringPortImpl struct {
	controller controller.BaseStringController
}

func CreateBaseStringPort() *BaseStringPortImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	baseStringController := impl.CreateBaseStringController()
	port := &BaseStringPortImpl{baseStringController}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return port
}

func (b BaseStringPortImpl) InitRoutes(r *chi.Mux) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	pattern := "/" + constants.BaseStrings
	tokenAuth := utils.ConfigJWTRoutes()
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Route(pattern, func(r chi.Router) {
			r.Post("/", b.controller.Create)
			r.Get("/", b.controller.FindAll)
			r.Get("/content/env", b.controller.FindByIdentifierAndLanguageAndStage)
			r.Get("/group/{id}", b.controller.FindByGroup)
			r.Get("/identifier/{identifier}", b.controller.FindByIdentifier)
			r.Get("/language/{id}", b.controller.FindByLanguage)
			r.Put("/", b.controller.Update)
			r.Patch("/{id}", b.controller.Disable)
			r.Delete("/{id}", b.controller.Delete)
		})
	})
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
