package impl

import (
	"encoding/xml"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/xliff"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/ports/utils"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service/impl"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type XliffControllerImpl struct {
	userService       service.UserService
	baseStringService service.BaseStringService
}

func CreateXliffController() *XliffControllerImpl {
	return &XliffControllerImpl{impl.CreateUserService(), impl.CreateBaseStringService()}
}

// swagger:route POST /xliffs Xliffs CreateXliff
// Read a .xliff file to create or update baseStrings.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: []BaseString
// - 400: ErrorDto
// - 401: ErrorDto
// - 403: ErrorDto
// - 422: ErrorDto
func (x XliffControllerImpl) Read(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, x.userService)
	if user == nil {
		return
	}
	var request xliff.Xliff
	if err := xml.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	stageId := r.URL.Query().Get("stage")
	groupId := r.URL.Query().Get("group")
	objectID, err := primitive.ObjectIDFromHex(stageId)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	groupObjectID, err := primitive.ObjectIDFromHex(groupId)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	baseStrings, err := x.baseStringService.Read(request, user, objectID, groupObjectID)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, baseStrings)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
