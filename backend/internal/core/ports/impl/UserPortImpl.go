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

type UserPortImpl struct {
	controller controller.UserController
}

func CreateUserPort() *UserPortImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	userRepository := mongodb.CreateUserRepository()
	userService := service.CreateUserService(userRepository, encrypt.CreateEncryptPasswordImpl())
	userController := impl.CreateUserController(userService)
	port := &UserPortImpl{userController}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return port
}

func (u UserPortImpl) InitUserRoutes(r *chi.Mux) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	u.CreateUserRoutes(r)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

func (u UserPortImpl) CreateUserRoutes(r *chi.Mux) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	pattern := "/users"
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
