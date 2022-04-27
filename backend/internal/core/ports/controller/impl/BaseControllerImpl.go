package impl

import (
	"encoding/json"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/utils"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service/impl"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"log"
	"net/http"
)

type BaseStringControllerImpl struct {
	baseStringService service.BaseStringService
	userService       service.UserService
}

func CreateBaseStringController() *BaseStringControllerImpl {
	return &BaseStringControllerImpl{impl.CreateBaseStringService(), impl.CreateUserService()}
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
func (g BaseStringControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if user := utils.CheckUserIsActive(w, r, g.userService); user == nil {
		return
	}
	var baseString domain.BaseString
	if err := json.NewDecoder(r.Body).Decode(&baseString); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	user, err := g.baseStringService.Create(baseString)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, user)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
