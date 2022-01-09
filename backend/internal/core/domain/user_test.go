package domain

import (
	"testing"
)

var user = User{
	Name:     "name",
	Password: "password",
	Role:     "role",
	Active:   true,
}

func TestEncryptPassword(t *testing.T) {
	result, err := User.EncryptPassword(user)
	if err != nil {
		t.Error("Expected", nil, "Got", err)
	}
	if result.Password == user.Password {
		t.Error("Expected", user.Password, "Got", result.Password)
	}
}

func TestCheckPassword(t *testing.T) {
	result, _ := User.EncryptPassword(user)
	isSamePassword := User.CheckPassword(result, user.Password)
	if !isSamePassword {
		t.Error("Expected the password are the same")
	}
}
