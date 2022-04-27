package controller

import (
	"net/http"
)

type BaseStringController interface {
	Create(w http.ResponseWriter, r *http.Request)
}
