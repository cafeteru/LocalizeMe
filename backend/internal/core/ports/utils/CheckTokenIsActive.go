package utils

import (
	"errors"
	"github.com/go-chi/jwtauth/v5"
	slog "github.com/go-eden/slf4go"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/service"
	"net/http"
)

func CheckUserIsActive(w http.ResponseWriter, r *http.Request, u service.UserService) *domain.User {
	slog.Debugf("%s: start", r.Context())
	_, tokenParts, _ := jwtauth.FromContext(r.Context())
	value, exists := tokenParts["email"]
	if !exists {
		createErrorResponse(w)
		return nil
	}
	user, err := u.FindByEmail(value.(string))
	if err != nil || user == nil || !user.IsActive {
		createErrorResponse(w)
		return nil
	}
	return user
}

func CheckUserIsAdmin(w http.ResponseWriter, r *http.Request, u service.UserService) *domain.User {
	user := CheckUserIsActive(w, r, u)
	if user != nil && user.IsAdmin {
		return user
	}
	err := errors.New(constants.UserNoAdmin)
	CreateErrorResponse(w, err, http.StatusUnauthorized)
	return nil
}

func createErrorResponse(w http.ResponseWriter) {
	err := errors.New(constants.UserNoActive)
	CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
}
