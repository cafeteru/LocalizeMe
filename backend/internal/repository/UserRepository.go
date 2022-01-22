package repository

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(user domain.User) (*mongo.InsertOneResult, error)
	FindByEmail(email string) (*domain.User, error)
}
