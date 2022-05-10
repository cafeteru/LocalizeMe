package controller

import (
	"net/http"
)

type BaseStringController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Disable(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindByGroup(w http.ResponseWriter, r *http.Request)
	FindByIdentifier(w http.ResponseWriter, r *http.Request)
	FindByIdentifierAndLanguage(w http.ResponseWriter, r *http.Request)
	FindByLanguage(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}
