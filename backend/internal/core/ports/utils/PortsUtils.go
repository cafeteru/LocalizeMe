package utils

import (
	"github.com/go-chi/jwtauth/v5"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"os"
)

func ConfigJWTRoutes() *jwtauth.JWTAuth {
	tools.LoadEnv()
	alg := os.Getenv("ALG")
	secret := os.Getenv("SECRET")
	tokenAuth := jwtauth.New(alg, []byte(secret), nil)
	return tokenAuth
}
