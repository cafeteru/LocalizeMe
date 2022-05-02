package controller

import (
	"net/http"
)

type StageController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Disable(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindByName(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}
