package utils

import (
	"github.com/go-chi/jwtauth/v5"
	"os"
	"uniovi-localizeme/tools"
)

func ConfigJWTRoutes() *jwtauth.JWTAuth {
	tools.LoadEnv()
	alg := os.Getenv("ALG")
	secret := os.Getenv("SECRET")
	tokenAuth := jwtauth.New(alg, []byte(secret), nil)
	return tokenAuth
}
