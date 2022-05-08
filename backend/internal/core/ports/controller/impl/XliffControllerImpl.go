package impl

import (
	"encoding/json"
	"encoding/xml"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/xmlDto"
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

// swagger:route POST /xliffs Xliffs ReadXliff
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
	var request xmlDto.Xliff
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

// swagger:route POST /xliffs/create Xliffs CreateXliff
// Create a .xliff file.
//
// Consumes:
// - application/json
// Produces:
// - application/xml
// - application/json
//
// Responses:
// - 200: Xliff
// - 400: ErrorDto
// - 401: ErrorDto
// - 403: ErrorDto
// - 422: ErrorDto
func (x XliffControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	user := utils.CheckUserIsActive(w, r, x.userService)
	if user == nil {
		return
	}
	var request dto.XliffDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	write, err := x.baseStringService.Write(request)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"))
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	err = xml.NewEncoder(w).Encode(write)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
