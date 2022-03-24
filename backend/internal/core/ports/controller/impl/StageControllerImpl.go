package impl

import (
	"encoding/json"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/utils"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"log"
	"net/http"
)

type StageControllerImpl struct {
	stageService service.StageService
	userService  service.UserService
}

func CreateStageController(s service.StageService, u service.UserService) *StageControllerImpl {
	return &StageControllerImpl{s, u}
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
func (s StageControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if user := utils.CheckUserIsAdmin(w, r, s.userService); user == nil {
		return
	}
	var request dto.StageRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	user, err := s.stageService.Create(request)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, user)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /stages Stages FindAll
// Return all stages.
//
// Responses:
// - 200: []Stage
func (s StageControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if user := utils.CheckUserIsAdmin(w, r, s.userService); user == nil {
		return
	}
	stages, err := s.stageService.FindAll()
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	utils.CreateResponse(w, http.StatusOK, stages)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
