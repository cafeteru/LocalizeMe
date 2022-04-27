package repository

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseStringRepository interface {
	Create(baseString domain.BaseString) (*mongo.InsertOneResult, error)
	FindById(id primitive.ObjectID) (*domain.BaseString, error)
	FindByIdentifier(name string) (*domain.BaseString, error)
}
