package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserControllerImpl_Login_NoRegister(t *testing.T) {
	server := CreateServer("8087")
	assert.NotNil(t, server)
	defer server.Start()
	err := server.Shutdown()
	if err != nil {
		assert.Failf(t, "Fail Shutdown", "Fail Shutdown")
	}
}
