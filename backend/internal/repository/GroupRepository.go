package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"uniovi-localizeme/internal/core/domain"
)

type GroupRepository interface {
	Create(group domain.Group) (*mongo.InsertOneResult, error)
	Delete(id primitive.ObjectID) (*mongo.DeleteResult, error)
	FindAll() (*[]domain.Group, error)
	FindById(id primitive.ObjectID) (*domain.Group, error)
	FindByPermissions(id primitive.ObjectID) (*[]domain.Group, error)
	FindByName(name string) (*domain.Group, error)
	FindCanWrite(id primitive.ObjectID) (*[]domain.Group, error)
	Update(group domain.Group) (*mongo.UpdateResult, error)
}
