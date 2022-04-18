package impl

import (
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupPortImpl_CreateUserPort(t *testing.T) {
	port := CreateGroupPort()
	assert.NotNil(t, port)
	assert.NotNil(t, port.controller)
}

func TestGroupPortImpl_InitRoutes(t *testing.T) {
	r := chi.NewRouter()
	port := CreateGroupPort()
	port.InitRoutes(r)
	assert.NotNil(t, port)
	assert.NotNil(t, port.controller)
}
