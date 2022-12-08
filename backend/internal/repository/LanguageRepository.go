package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"uniovi-localizeme/internal/core/domain"
)

type LanguageRepository interface {
	Create(language domain.Language) (*mongo.InsertOneResult, error)
	Delete(id primitive.ObjectID) (*mongo.DeleteResult, error)
	FindAll() (*[]domain.Language, error)
	FindById(id primitive.ObjectID) (*domain.Language, error)
	FindByIsoCode(isoCode string) (*domain.Language, error)
	Update(language domain.Language) (*mongo.UpdateResult, error)
}
