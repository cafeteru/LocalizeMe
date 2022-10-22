package mock

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"uniovi-localizeme/internal/core/domain"
)

func createMockUser() domain.User {
	return domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "user@email.com",
		Password: "password",
		Admin:    false,
		Active:   true,
	}
}
