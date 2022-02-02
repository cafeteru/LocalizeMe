package server

import (
	"testing"
)

func TestUserControllerImpl_Login_NoRegister(t *testing.T) {
	server := CreateServer("8087")
	if server == nil {
		t.Error("Expected create server but got", nil)
	}
}
