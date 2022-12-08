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

type XliffPortImpl struct {
	controller controller.XliffController
}

func CreateXliffPort() *XliffPortImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	XliffController := impl.CreateXliffController()
	port := &XliffPortImpl{XliffController}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return port
}

func (l XliffPortImpl) InitRoutes(r *chi.Mux) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	pattern := "/" + constants.Xliffs
	tokenAuth := utils.ConfigJWTRoutes()
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Route(pattern, func(r chi.Router) {
			r.Post("/", l.controller.Read)
			r.Post("/create", l.controller.Create)
		})
	})
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
