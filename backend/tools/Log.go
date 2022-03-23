package tools

import (
	"errors"
	"log"
)

func ErrorLog(errorMessage string, functionName string) error {
	err := errors.New(errorMessage)
	log.Printf("%s: %s", functionName, err)
	return err
}

func ErrorLogDetails(original error, errorMessage string, functionName string) error {
	err := errors.New(errorMessage)
	log.Printf("%s: %s -> %s", functionName, err, original)
	return err
}
