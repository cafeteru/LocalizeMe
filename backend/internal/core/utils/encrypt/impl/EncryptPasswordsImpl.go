package impl

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type EncryptPasswordImpl struct{}

func CreateEncryptPasswordImpl() *EncryptPasswordImpl {
	return &EncryptPasswordImpl{}
}

func (e EncryptPasswordImpl) EncryptPassword(password string) (string, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	result := string(bytes)
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, err
}

func (e EncryptPasswordImpl) CheckPassword(encryptPassword, password string) bool {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	result := err == nil
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result
}
