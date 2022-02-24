package impl

import (
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserPortImpl_CreateUserPort(t *testing.T) {
	userPort := CreateUserPort()
	assert.NotNil(t, userPort)
	assert.NotNil(t, userPort.controller)
}

func TestUserPortImpl_InitRoutes(t *testing.T) {
	r := chi.NewRouter()
	userPort := CreateUserPort()
	userPort.InitRoutes(r)
	assert.NotNil(t, userPort)
	assert.NotNil(t, userPort.controller)
}
