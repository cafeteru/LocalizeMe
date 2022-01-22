package tools

import (
	"errors"
	slog "github.com/go-eden/slf4go"
)

func ErrorLog(errorMessage string, functionName string) error {
	err := errors.New(errorMessage)
	slog.Errorf("%s: %s", functionName, err)
	return err
}

func ErrorLogDetails(original error, errorMessage string, functionName string) error {
	err := errors.New(errorMessage)
	slog.Errorf("%s: %s -> %s", functionName, err, original)
	return err
}
