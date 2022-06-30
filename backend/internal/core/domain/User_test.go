package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_ClearPassword(t *testing.T) {
	user := User{Password: "1"}
	assert.NotEmpty(t, user.Password)
	user.ClearPassword()
	assert.Empty(t, user.Password)
}
