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

type LanguageControllerImpl struct {
	languageService service.LanguageService
	userService     service.UserService
}

func CreateLanguageController(l service.LanguageService, u service.UserService) *LanguageControllerImpl {
	return &LanguageControllerImpl{l, u}
}

// swagger:route POST /languages Languages CreateLanguage
// Create a new language.
//
// Consumes:
// - application/json
//
// Responses:
// - 200: Language
// - 400: ErrorDto
// - 401: ErrorDto
// - 422: ErrorDto
func (l LanguageControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if user := utils.CheckUserIsAdmin(w, r, l.userService); user == nil {
		return
	}
	var request dto.LanguageDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	user, err := l.languageService.Create(request)
	if err != nil {
		utils.CreateErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.CreateResponse(w, http.StatusCreated, user)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
