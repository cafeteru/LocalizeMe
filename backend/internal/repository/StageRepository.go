package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"uniovi-localizeme/internal/core/domain"
)

type StageRepository interface {
	Create(stage domain.Stage) (*mongo.InsertOneResult, error)
	Delete(id primitive.ObjectID) (*mongo.DeleteResult, error)
	FindAll() (*[]domain.Stage, error)
	FindById(id primitive.ObjectID) (*domain.Stage, error)
	FindByName(name string) (*domain.Stage, error)
	Update(stage domain.Stage) (*mongo.UpdateResult, error)
}
