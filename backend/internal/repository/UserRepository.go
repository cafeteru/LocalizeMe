package repository

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(user domain.User) (*mongo.InsertOneResult, error)
	Delete(id primitive.ObjectID) (*mongo.DeleteResult, error)
	FindAll() (*[]domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindById(id primitive.ObjectID) (*domain.User, error)
	Update(id primitive.ObjectID, user domain.User) (*mongo.UpdateResult, error)
}
