package repository

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StageRepository interface {
	Create(stage domain.Stage) (*mongo.InsertOneResult, error)
	FindAll() (*[]domain.Stage, error)
	FindById(id primitive.ObjectID) (*domain.Stage, error)
	FindByName(name string) (*domain.Stage, error)
	Update(stage domain.Stage) (*mongo.UpdateResult, error)
}
