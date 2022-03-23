package controller

import (
	"net/http"
)

type StageController interface {
	Create(w http.ResponseWriter, r *http.Request)
}
