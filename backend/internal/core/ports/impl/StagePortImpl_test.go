package impl

import (
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStagePortImpl_CreateUserPort(t *testing.T) {
	userPort := CreateStagePort()
	assert.NotNil(t, userPort)
	assert.NotNil(t, userPort.controller)
}

func TestStagePortImpl_InitRoutes(t *testing.T) {
	r := chi.NewRouter()
	userPort := CreateStagePort()
	userPort.InitStageRoutes(r)
	assert.NotNil(t, userPort)
	assert.NotNil(t, userPort.controller)
}
