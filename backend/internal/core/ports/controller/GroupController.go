package controller

import (
	"net/http"
)

type GroupController interface {
	Create(w http.ResponseWriter, r *http.Request)
}
