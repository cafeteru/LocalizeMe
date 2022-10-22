package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"uniovi-localizeme/internal/core/domain"
)

type UserRepository interface {
	Create(user domain.User) (*mongo.InsertOneResult, error)
	Delete(id primitive.ObjectID) (*mongo.DeleteResult, error)
	FindAll() (*[]domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindById(id primitive.ObjectID) (*domain.User, error)
	Update(user domain.User) (*mongo.UpdateResult, error)
}
