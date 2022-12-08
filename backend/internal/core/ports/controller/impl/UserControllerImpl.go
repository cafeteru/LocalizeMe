package impl

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"uniovi-localizeme/constants"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/internal/core/domain/dto"
	"uniovi-localizeme/internal/core/ports/utils"
	"uniovi-localizeme/internal/core/service"
	"uniovi-localizeme/internal/core/service/impl"
	"uniovi-localizeme/tools"
)

type UserControllerImpl struct {
	service service.UserService
}

func CreateUserController() *UserControllerImpl {
	return &UserControllerImpl{impl.CreateUserService()}
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
// - 401: ErrorDto
// - 422: ErrorDto
func (u UserControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	var request dto.UserDto
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
	log.Printf("%s: end", tools.GetCurrentFuncName())
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
// - 401: ErrorDto
// - 422: ErrorDto
func (u UserControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	var request dto.UserDto
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
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route DELETE /users/{id} Users DeleteUser
// Return a user by ID.
//
// Responses:
// - 200: description:bool
// - 400: ErrorDto
// - 401: ErrorDto
func (u UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsAdmin(w, r, u.service)
	if user == nil {
		return
	}
	id := chi.URLParam(r, "id")
	if id == "" {
		err := errors.New(constants.IdNoValid)
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	objectID, _ := primitive.ObjectIDFromHex(id)
	result, err := u.service.Delete(objectID)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusNotFound)
		return
	}
	utils.CreateResponse(w, http.StatusOK, result)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route PATCH /users/{id} Users DisableUser
// Disable of a user.
//
// Responses:
// - 200: User
// - 400: ErrorDto
// - 401: ErrorDto
func (u UserControllerImpl) Disable(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	isAdmin := utils.CheckUserIsAdmin(w, r, u.service)
	if isAdmin == nil {
		return
	}
	id := chi.URLParam(r, "id")
	if id == "" {
		err := errors.New(constants.IdNoValid)
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	objectID, _ := primitive.ObjectIDFromHex(id)
	user, err := u.service.Disable(objectID)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, user)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /users Users FindUsersAll
// Return all users.
//
// Responses:
// - 200: []User
// - 400: ErrorDto
// - 401: ErrorDto
// - 500: ErrorDto
func (u UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if user := utils.CheckUserIsActive(w, r, u.service); user == nil {
		return
	}
	users, err := u.service.FindAll()
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, users)
	log.Printf("%s: end", tools.GetCurrentFuncName())
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
// - 401: ErrorDto
func (u UserControllerImpl) FindMe(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, u.service)
	if user == nil {
		return
	}
	user.ClearPassword()
	utils.CreateResponse(w, http.StatusOK, user)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /users/{id} Users FindUserById
// Return the information of the user by id.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: User
// - 400: ErrorDto
// - 401: ErrorDto
// - 404: ErrorDto
func (u UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if utils.CheckUserIsAdmin(w, r, u.service) == nil {
		return
	}
	id := chi.URLParam(r, "id")
	if id == "" {
		err := errors.New(constants.IdNoValid)
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	user, err := u.service.FindById(objectID)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusNotFound)
		return
	}
	user.ClearPassword()
	utils.CreateResponse(w, http.StatusOK, user)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route PUT /users Users UpdateUser
// Update the information of a user.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: User
// - 400: ErrorDto
// - 401: ErrorDto
// - 422: ErrorDto
func (u UserControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	isAdmin := utils.CheckUserIsAdmin(w, r, u.service)
	if isAdmin == nil {
		return
	}
	var request domain.User
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	user, err := u.service.Update(request)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, user)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route PUT /users/me Users UpdateMe
// Update the information of the identified user.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: User
// - 400: ErrorDto
// - 401: ErrorDto
// - 422: ErrorDto
func (u UserControllerImpl) UpdateMe(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, u.service)
	if user == nil {
		return
	}
	var request domain.User
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	request.Admin = user.Admin
	userUpdate, err := u.service.Update(request)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, userUpdate)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
