package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"uniovi-localizeme/internal/core/domain"
)

type BaseStringRepository interface {
	Create(baseString domain.BaseString) (*mongo.InsertOneResult, error)
	Delete(id primitive.ObjectID) (*mongo.DeleteResult, error)
	FindAll() (*[]domain.BaseString, error)
	FindByGroup(id primitive.ObjectID) (*[]domain.BaseString, error)
	FindByLanguage(id primitive.ObjectID) (*[]domain.BaseString, error)
	FindById(id primitive.ObjectID) (*domain.BaseString, error)
	FindByIdentifier(identifier string) (*domain.BaseString, error)
	FindByIdentifierAndLanguageAndStage(identifier string, isoCode string, stageName string) (*domain.BaseString, error)
	FindByPermissions(id primitive.ObjectID) (*[]domain.BaseString, error)
	Update(baseString domain.BaseString) (*mongo.UpdateResult, error)
}
