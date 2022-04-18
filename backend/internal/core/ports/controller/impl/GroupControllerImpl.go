package impl

import (
	"encoding/json"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/utils"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service/impl"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"log"
	"net/http"
)

type GroupControllerImpl struct {
	groupService service.GroupService
	userService  service.UserService
}

func CreateGroupController() *GroupControllerImpl {
	return &GroupControllerImpl{impl.CreateGroupService(), impl.CreateUserService()}
}

// swagger:route POST /groups Groups CreateGroup
// Create a new group.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: Group
// - 400: ErrorDto
// - 401: ErrorDto
// - 422: ErrorDto
func (g GroupControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if user := utils.CheckUserIsActive(w, r, g.userService); user == nil {
		return
	}
	var groupDto dto.GroupDto
	if err := json.NewDecoder(r.Body).Decode(&groupDto); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	user, err := g.groupService.Create(groupDto)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, user)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

// swagger:route GET /groups Groups FindAllGroups
// Return all groups.
//
// Responses:
// - 200: []Group
// - 400: ErrorDto
// - 401: ErrorDto
// - 500: ErrorDto
func (l GroupControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if user := utils.CheckUserIsAdmin(w, r, l.userService); user == nil {
		return
	}
	stages, err := l.groupService.FindAll()
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusInternalServerError)
		return
	}
	utils.CreateResponse(w, http.StatusOK, stages)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
