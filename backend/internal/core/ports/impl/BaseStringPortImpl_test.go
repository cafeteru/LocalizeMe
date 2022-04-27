package impl

import (
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBaseStringPortImpl_CreateUserPort(t *testing.T) {
	port := CreateBaseStringPort()
	assert.NotNil(t, port)
	assert.NotNil(t, port.controller)
}

func TestBaseStringPortImpl_InitRoutes(t *testing.T) {
	r := chi.NewRouter()
	port := CreateBaseStringPort()
	port.InitRoutes(r)
	assert.NotNil(t, port)
	assert.NotNil(t, port.controller)
}
