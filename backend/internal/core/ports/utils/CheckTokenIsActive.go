package utils

import (
	"errors"
	"github.com/go-chi/jwtauth/v5"
	slog "github.com/go-eden/slf4go"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service"
	"net/http"
)

func CheckUserIsActive(w http.ResponseWriter, r *http.Request, u service.UserService) bool {
	slog.Debugf("%s: start", r.Context())
	_, tokenParts, _ := jwtauth.FromContext(r.Context())
	value, exists := tokenParts["email"]
	if !exists {
		return createErrorResponse(w)
	}
	user, err := u.FindByEmail(value.(string))
	if err != nil || user == nil || !user.IsActive {
		return createErrorResponse(w)
	}
	return true
}

func createErrorResponse(w http.ResponseWriter) bool {
	err := errors.New(constants.UserNoActive)
	CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
	return false
}
