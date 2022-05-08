package impl

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/controller"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/controller/impl"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/utils"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"log"
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
