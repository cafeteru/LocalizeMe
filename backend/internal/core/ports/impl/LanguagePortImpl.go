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

type LanguagePortImpl struct {
	controller controller.LanguageController
}

func CreateLanguagePort() *LanguagePortImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	languageController := impl.CreateLanguageController()
	port := &LanguagePortImpl{languageController}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return port
}

func (l LanguagePortImpl) InitRoutes(r *chi.Mux) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	pattern := "/" + constants.Languages
	tokenAuth := utils.ConfigJWTRoutes()
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Route(pattern, func(r chi.Router) {
			r.Post("/", l.controller.Create)
			r.Get("/", l.controller.FindAll)
			r.Put("/", l.controller.Update)
			r.Patch("/{id}", l.controller.Disable)
			r.Delete("/{id}", l.controller.Delete)
		})
	})
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
