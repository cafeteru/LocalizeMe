package utils

import (
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"uniovi-localizeme/constants"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/internal/core/service"
)

func CheckUserIsActive(w http.ResponseWriter, r *http.Request, u service.UserService) *domain.User {
	log.Printf("%s: start", r.Context())
	_, tokenParts, _ := jwtauth.FromContext(r.Context())
	value, exists := tokenParts["id"]
	if !exists {
		err := errors.New(constants.InvalidToken)
		CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return nil
	}
	objectID, err := primitive.ObjectIDFromHex(value.(string))
	if err != nil {
		CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return nil
	}
	user, err := u.FindById(objectID)
	if err != nil {
		CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return nil
	}
	if !user.Active {
		err := errors.New(constants.UserNoActive)
		CreateErrorResponse(w, err, http.StatusUnprocessableEntity)
		return nil
	}
	return user
}

func CheckUserIsAdmin(w http.ResponseWriter, r *http.Request, u service.UserService) *domain.User {
	user := CheckUserIsActive(w, r, u)
	if user == nil {
		return nil
	}
	if user != nil && user.Admin {
		return user
	}
	err := errors.New(constants.UserNoAdmin)
	CreateErrorResponse(w, err, http.StatusUnauthorized)
	return nil
}
