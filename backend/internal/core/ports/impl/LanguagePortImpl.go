package impl

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/controller"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/controller/impl"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/utils"
	service "gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service/impl"
	encrypt "gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/utils/encrypt/impl"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/repository/mongodb"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"log"
)

type LanguagePortImpl struct {
	controller controller.LanguageController
}

func CreateLanguagePort() *LanguagePortImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	languageRepository := mongodb.CreateLanguageRepository()
	languageService := service.CreateLanguageService(languageRepository)
	userRepository := mongodb.CreateUserRepository()
	userService := service.CreateUserService(userRepository, encrypt.CreateEncryptPasswordImpl())
	languageController := impl.CreateLanguageController(languageService, userService)
	port := &LanguagePortImpl{languageController}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return port
}

func (l LanguagePortImpl) InitRoutes(r *chi.Mux) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	pattern := "/languages"
	tokenAuth := utils.ConfigJWTRoutes()
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Route(pattern, func(r chi.Router) {
			r.Post("/", l.controller.Create)
			r.Get("/", l.controller.FindAll)
		})
	})
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
