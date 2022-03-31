package impl

import (
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLanguagePortImpl_CreateUserPort(t *testing.T) {
	port := CreateLanguagePort()
	assert.NotNil(t, port)
	assert.NotNil(t, port.controller)
}

func TestLanguagePortImpl_InitRoutes(t *testing.T) {
	r := chi.NewRouter()
	port := CreateLanguagePort()
	port.InitRoutes(r)
	assert.NotNil(t, port)
	assert.NotNil(t, port.controller)
}
