package impl

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/utils"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type UserControllerImpl struct {
	service service.UserService
}

func CreateUserController(u service.UserService) *UserControllerImpl {
	return &UserControllerImpl{u}
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
	log.Printf("%s: start", tools.GetCurrentFuncName())
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
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route POST /userService Users CreateUser
// Create a new user.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: User
// - 400: ErrorDto
func (u UserControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
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
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route DELETE /userService/{id} Users Delete
// Return a user by email.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: bool
// - 400: ErrorDto
// - 401: ErrorDto
// - 404: ErrorDto
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

// swagger:route PATCH /userService/{id} Users Disable
// Disable of a user.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: User
// - 400: ErrorDto
// - 401: ErrorDto
// - 404: ErrorDto
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
	utils.CreateResponse(w, http.StatusCreated, user)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /userService Users FindAll
// Return all userService.
//
// Responses:
// - 200: []User
func (u UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if user := utils.CheckUserIsAdmin(w, r, u.service); user == nil {
		return
	}
	users, err := u.service.FindAll()
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	utils.CreateResponse(w, http.StatusOK, users)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /userService/me Users GetMe
// Return the information of the identified user.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: User
// - 400: ErrorDto
func (u UserControllerImpl) FindMe(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, u.service)
	if user == nil {
		return
	}
	user.Password = ""
	utils.CreateResponse(w, http.StatusOK, user)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /userService/{id} Users FindById
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
	user, err := u.service.FindById(objectID)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusNotFound)
		return
	}
	user.Password = ""
	utils.CreateResponse(w, http.StatusOK, user)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route PUT /userService/{id} Users Update
// Update the information of a user.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: User
// - 400: ErrorDto
// - 401: ErrorDto
// - 404: ErrorDto
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
	id := chi.URLParam(r, "id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	user, err := u.service.Update(objectID, request)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, user)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route PUT /userService/me Users UpdateMe
// Update the information of the identified user.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: User
// - 400: ErrorDto
// - 401: ErrorDto
// - 404: ErrorDto
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
	request.IsAdmin = user.IsAdmin
	userUpdate, err := u.service.Update(user.ID, request)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, userUpdate)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
