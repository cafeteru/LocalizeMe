package impl

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/controller"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/controller/impl"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/utils"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"log"
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
	pattern := "/groups"
	tokenAuth := utils.ConfigJWTRoutes()
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Route(pattern, func(r chi.Router) {
			r.Post("/", g.controller.Create)
			r.Get("/", g.controller.FindAll)
			r.Put("/", g.controller.Update)
			r.Patch("/{id}", g.controller.Disable)
		})
	})
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
