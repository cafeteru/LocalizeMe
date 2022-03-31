package controller

import (
	"net/http"
)

type LanguageController interface {
	Create(w http.ResponseWriter, r *http.Request)
}
