package controller

import (
	"net/http"
)

type UserController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Disable(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindByEmail(w http.ResponseWriter, r *http.Request)
	FindMe(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}
