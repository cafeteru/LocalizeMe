package impl

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/controller"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/controller/impl"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/utils"
	service "gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service/impl"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository/mongodb"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"log"
)

type StagePortImpl struct {
	controller controller.StageController
}

func CreateStagePort() *StagePortImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	stageRepository := mongodb.CreateStageRepository()
	stageService := service.CreateStageService(stageRepository)
	stageController := impl.CreateStageController(stageService)
	port := &StagePortImpl{stageController}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return port
}

func (u StagePortImpl) InitStageRoutes(r *chi.Mux) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	u.CreateStageRoutes(r)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

func (u StagePortImpl) CreateStageRoutes(r *chi.Mux) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	pattern := "/stages"
	tokenAuth := utils.ConfigJWTRoutes()
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Route(pattern, func(r chi.Router) {
			r.Post("/", u.controller.Create)
		})
	})
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
