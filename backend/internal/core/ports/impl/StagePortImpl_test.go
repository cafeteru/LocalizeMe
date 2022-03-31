package impl

import (
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStagePortImpl_CreateUserPort(t *testing.T) {
	port := CreateStagePort()
	assert.NotNil(t, port)
	assert.NotNil(t, port.controller)
}

func TestStagePortImpl_InitRoutes(t *testing.T) {
	r := chi.NewRouter()
	port := CreateStagePort()
	port.InitRoutes(r)
	assert.NotNil(t, port)
	assert.NotNil(t, port.controller)
}
