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

type StageControllerImpl struct {
	stageService service.StageService
	userService  service.UserService
}

func CreateStageController() *StageControllerImpl {
	return &StageControllerImpl{impl.CreateStageService(), impl.CreateUserService()}
}

// swagger:route POST /stages Stages CreateStage
// Create a new stage.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: Stage
// - 400: ErrorDto
// - 401: ErrorDto
// - 422: ErrorDto
func (s StageControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if user := utils.CheckUserIsAdmin(w, r, s.userService); user == nil {
		return
	}
	var request dto.StageDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	stage, err := s.stageService.Create(request)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, stage)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route DELETE /stages/{id} Stages DeleteStage
// Return a stage by id.
//
// Responses:
// - 200: description:bool
// - 400: ErrorDto
// - 401: ErrorDto
func (s StageControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsAdmin(w, r, s.userService)
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
	result, err := s.stageService.Delete(objectID)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusNotFound)
		return
	}
	utils.CreateResponse(w, http.StatusOK, result)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route PATCH /stages/{id} Stages DisableStage
// Disable of a stage.
//
// Responses:
// - 200: Stage
// - 400: ErrorDto
// - 401: ErrorDto
func (s StageControllerImpl) Disable(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	isAdmin := utils.CheckUserIsAdmin(w, r, s.userService)
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
	stage, err := s.stageService.Disable(objectID)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, stage)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /stages Stages FindAllStages
// Return all stages.
//
// Responses:
// - 200: []Stage
// - 400: ErrorDto
// - 401: ErrorDto
// - 500: ErrorDto
func (s StageControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if user := utils.CheckUserIsActive(w, r, s.userService); user == nil {
		return
	}
	stages, err := s.stageService.FindAll()
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusOK, stages)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /stages/name/{name} Stages FindStageByName
// Return the information of the stage by name.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: Stage
// - 400: ErrorDto
// - 401: ErrorDto
// - 404: ErrorDto
func (s StageControllerImpl) FindByName(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if utils.CheckUserIsAdmin(w, r, s.userService) == nil {
		return
	}
	name := chi.URLParam(r, "name")
	if name == "" {
		err := errors.New(constants.NameNoValid)
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	stage, err := s.stageService.FindByName(name)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusNotFound)
		return
	}
	utils.CreateResponse(w, http.StatusOK, stage)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route PUT /stages Stages UpdateStage
// Update the information of a stage.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: Stage
// - 400: ErrorDto
// - 401: ErrorDto
// - 422: ErrorDto
func (s StageControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	isAdmin := utils.CheckUserIsAdmin(w, r, s.userService)
	if isAdmin == nil {
		return
	}
	var request domain.Stage
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	user, err := s.stageService.Update(request)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, user)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
