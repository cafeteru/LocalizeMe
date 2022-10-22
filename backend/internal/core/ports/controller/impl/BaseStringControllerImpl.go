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
	"uniovi-localizeme/internal/core/ports/utils"
	"uniovi-localizeme/internal/core/service"
	"uniovi-localizeme/internal/core/service/impl"
	"uniovi-localizeme/tools"
)

type BaseStringControllerImpl struct {
	baseStringService service.BaseStringService
	userService       service.UserService
}

func CreateBaseStringController() *BaseStringControllerImpl {
	return &BaseStringControllerImpl{
		impl.CreateBaseStringService(),
		impl.CreateUserService(),
	}
}

// swagger:route POST /baseStrings BaseStrings CreateBaseString
// Create a new baseString.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: BaseString
// - 400: ErrorDto
// - 401: ErrorDto
// - 422: ErrorDto
func (b BaseStringControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, b.userService)
	if user == nil {
		return
	}
	var baseString domain.BaseString
	if err := json.NewDecoder(r.Body).Decode(&baseString); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	baseString, err := b.baseStringService.Create(baseString, user)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, baseString)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route DELETE /baseStrings/{id} BaseStrings DeleteBaseString
// Delete a baseString by id.
//
// Responses:
// - 200: description:bool
// - 400: ErrorDto
// - 401: ErrorDto
// - 403: ErrorDto
func (b BaseStringControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, b.userService)
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
	result, err := b.baseStringService.Delete(objectID, user)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusNotFound)
		return
	}
	utils.CreateResponse(w, http.StatusOK, result)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route PATCH /baseStrings/{id} BaseStrings DisableBaseString
// Disable of a baseString.
//
// Responses:
// - 200: BaseString
// - 400: ErrorDto
// - 401: ErrorDto
// - 403: ErrorDto
func (b BaseStringControllerImpl) Disable(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, b.userService)
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
	stage, err := b.baseStringService.Disable(objectID, user)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, stage)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /baseStrings BaseStrings FindAllBaseStrings
// Return all baseStrings.
//
// Responses:
// - 200: []BaseString
// - 400: ErrorDto
// - 401: ErrorDto
// - 500: ErrorDto
func (b BaseStringControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, b.userService)
	if user == nil {
		return
	}
	var baseStrings *[]domain.BaseString
	var err error
	if user.Admin {
		baseStrings, err = b.baseStringService.FindAll()
	} else {
		baseStrings, err = b.baseStringService.FindByPermissions(user.ID)
	}
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, baseStrings)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /baseStrings/group/{id} BaseStrings FindByGroupBaseStrings
// Return all baseStrings from a group.
//
//	Parameters:
//	  + name: id
//	    in: path
//	    description: BaseString's id
//	    type: string
//	    required: true
//
// Responses:
// - 200: []BaseString
// - 400: ErrorDto
// - 401: ErrorDto
// - 500: ErrorDto
func (b BaseStringControllerImpl) FindByGroup(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, b.userService)
	if user == nil {
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
	baseStrings, err := b.baseStringService.FindByGroup(objectID, user)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, baseStrings)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /baseStrings/language/{id} BaseStrings FindByLanguageBaseStrings
// Return all baseStrings from a language.
//
// Responses:
// - 200: []BaseString
// - 400: ErrorDto
// - 401: ErrorDto
// - 500: ErrorDto
func (b BaseStringControllerImpl) FindByLanguage(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, b.userService)
	if user == nil {
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
	baseStrings, err := b.baseStringService.FindByLanguage(objectID, user)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, baseStrings)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// Return a baseString from an identifier.
//
// Responses:
// - 200: BaseString
// - 400: ErrorDto
// - 401: ErrorDto
// - 500: ErrorDto
//
//swagger:route GET /baseStrings/identifier/{identifier} BaseStrings FindByIdentifierBaseStrings
func (b BaseStringControllerImpl) FindByIdentifier(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, b.userService)
	if user == nil {
		return
	}
	identifier := chi.URLParam(r, "identifier")
	if identifier == "" {
		err := errors.New(constants.IdentifierBaseStringInvalid)
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	baseStrings, err := b.baseStringService.FindByIdentifier(identifier, user)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, baseStrings)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /content/env BaseStrings FindByIdentifierAndLanguageAndStage
// Return a baseString from an identifier, an isoCode and stage.
//
//	Parameters:
//	  + name: identifier
//	    description: BaseString's identifier.
//	    in: query
//	    type: string
//	  + name: isoCode
//	    description: Language's isoCode.
//	    in: query
//	    type: string
//	  + name: stage
//	    description: Stage's id.
//	    in: query
//	    type: string
//
// Responses:
// - 200: BaseString
// - 400: ErrorDto
// - 401: ErrorDto
// - 500: ErrorDto
func (b BaseStringControllerImpl) FindByIdentifierAndLanguageAndStage(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, b.userService)
	if user == nil {
		return
	}
	identifier := r.URL.Query().Get("identifier")
	isoCode := r.URL.Query().Get("isoCode")
	stageName := r.URL.Query().Get("stage")
	baseStrings, err := b.baseStringService.FindByIdentifierAndLanguageAndState(identifier, isoCode, stageName, user)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, baseStrings)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route PUT /baseStrings BaseStrings UpdateBaseString
// Update the information of a baseString.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: BaseString
// - 400: ErrorDto
// - 401: ErrorDto
// - 422: ErrorDto
func (b BaseStringControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, b.userService)
	if user == nil {
		return
	}
	var baseString domain.BaseString
	if err := json.NewDecoder(r.Body).Decode(&baseString); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	update, err := b.baseStringService.Update(baseString, user)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, update)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
