package impl

import (
	"testing"
)

func TestEncryptPassword(t *testing.T) {
	password := "password"
	encrypt := CreateEncryptPasswordImpl()
	result, err := encrypt.EncryptPassword(password)
	if err != nil {
		t.Error("Expected", nil, "Got", err)
	}
	if password == result {
		t.Error("Expected", password, "Got", result)
	}
}

func TestCheckPassword(t *testing.T) {
	p1 := "password"
	p2 := "password"
	encrypt := CreateEncryptPasswordImpl()
	result, _ := encrypt.EncryptPassword(p1)
	isSamePassword := encrypt.CheckPassword(result, p2)
	if !isSamePassword {
		t.Error("Expected the password are the same")
	}
}
