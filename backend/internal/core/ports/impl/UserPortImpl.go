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
)

type UserPortImpl struct {
	controller controller.UserController
}

func CreateUserPort() *UserPortImpl {
	userRepository := mongodb.CreateUserRepository()
	userService := service.CreateUserService(userRepository, encrypt.CreateEncryptPasswordImpl())
	userController := impl.CreateUserController(userService)
	return &UserPortImpl{userController}
}

func (u UserPortImpl) InitRoutes(r *chi.Mux) {
	u.CreateUserRoutes(r)
}

func (u UserPortImpl) CreateUserRoutes(r *chi.Mux) {
	pattern := "/users"
	tokenAuth := utils.ConfigJWTRoutes()
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Route(pattern, func(r chi.Router) {
			r.Get("/", u.controller.FindAll)
			r.Get("/me", u.controller.FindMe)
			r.Get("/email/{email}", u.controller.FindByEmail)
			r.Put("/{id}", u.controller.Update)
			r.Patch("/{id}", u.controller.Disable)
			r.Delete("/{id}", u.controller.Delete)
		})
	})
	r.Group(func(r chi.Router) {
		r.Post("/login", u.controller.Login)
		r.Post(pattern, u.controller.Create)
	})
}
