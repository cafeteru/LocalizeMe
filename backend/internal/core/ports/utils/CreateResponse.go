package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"uniovi-localizeme/internal/core/domain/dto"
	"uniovi-localizeme/tools"
)

func CreateResponse(w http.ResponseWriter, code int, result interface{}) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
}

func CreateErrorResponse(w http.ResponseWriter, err error, code int) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(code)
		errorDto := dto.ErrorDto{Message: err.Error()}
		errEncoder := json.NewEncoder(w).Encode(errorDto)
		if errEncoder != nil {
			return
		}
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
