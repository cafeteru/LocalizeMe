package ports

import (
	"encoding/json"
	slog "github.com/go-eden/slf4go"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"net/http"
)

func CreateResponse(w http.ResponseWriter, code int, result interface{}) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
}

func CreateErrorResponse(w http.ResponseWriter, err error, code int) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(code)
		errorDto := dto.ErrorDto{Message: err.Error()}
		errEncoder := json.NewEncoder(w).Encode(errorDto)
		if errEncoder != nil {
			return
		}
	}
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
}
