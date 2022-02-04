package impl

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth/v5"
	slog "github.com/go-eden/slf4go"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/utils"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"net/http"
)

type UserControllerImpl struct {
	service service.UserService
}

func CreateUserController(u service.UserService) *UserControllerImpl {
	return &UserControllerImpl{u}
}

func (u UserControllerImpl) CreateUserRoutes(r *chi.Mux) {
	pattern := "/users"
	tokenAuth := utils.ConfigJWTRoutes()
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Route(pattern, func(r chi.Router) {
			r.Get("/", u.FindAll)
			r.Get("/me", u.FindMe)
			r.Get("/email/{email}", u.FindByEmail)
		})
	})
	r.Group(func(r chi.Router) {
		r.Post("/login", u.Login)
		r.Post(pattern, u.Create)
	})
}

// swagger:route POST /login Users Login
// Get token to user the application.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: TokenDto
// - 400: ErrorDto
func (u UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	var request dto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	tokenDto, err := u.service.Login(request)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, tokenDto)
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route POST /users Users CreateUser
// Create a new user.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: User
// - 400: ErrorDto
func (u UserControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	var request dto.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	user, err := u.service.Create(request)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, user)
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /users Users FindAll
// Return all users.
//
// Responses:
// - 200: []User
func (u UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	if user := utils.CheckUserIsAdmin(w, r, u.service); user == nil {
		return
	}
	users, err := u.service.FindAll()
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	utils.CreateResponse(w, http.StatusOK, users)
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /users/me Users GetMe
// Return the information of the identified user.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: User
// - 400: ErrorDto
func (u UserControllerImpl) FindMe(w http.ResponseWriter, r *http.Request) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, u.service)
	if user == nil {
		return
	}
	user.Password = ""
	utils.CreateResponse(w, http.StatusOK, user)
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /users/email/{email} Users FindByEmail
// Return the information of the user by email.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: User
// - 400: ErrorDto
// - 401: ErrorDto
// - 404: ErrorDto
func (u UserControllerImpl) FindByEmail(w http.ResponseWriter, r *http.Request) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsAdmin(w, r, u.service)
	if user == nil {
		return
	}
	email := chi.URLParam(r, "email")
	if email == "" {
		err := errors.New(constants.EmailAlreadyRegister)
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	user, err := u.service.FindByEmail(email)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusNotFound)
		return
	}
	user.Password = ""
	utils.CreateResponse(w, http.StatusOK, user)
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
}
