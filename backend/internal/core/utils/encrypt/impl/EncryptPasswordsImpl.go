package impl

import (
	slog "github.com/go-eden/slf4go"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"golang.org/x/crypto/bcrypt"
)

type EncryptPasswordImpl struct{}

func CreateEncryptPasswordImpl() *EncryptPasswordImpl {
	return &EncryptPasswordImpl{}
}

func (e EncryptPasswordImpl) EncryptPassword(password string) (string, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	result := string(bytes)
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return result, err
}

func (e EncryptPasswordImpl) CheckPassword(encryptPassword, password string) bool {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))
	result := err == nil
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return result
}
