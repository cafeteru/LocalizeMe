package repository

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LanguageRepository interface {
	Create(language domain.Language) (*mongo.InsertOneResult, error)
	Delete(id primitive.ObjectID) (*mongo.DeleteResult, error)
	FindAll() (*[]domain.Language, error)
	FindById(id primitive.ObjectID) (*domain.Language, error)
	FindByIsoCode(isoCode string) (*domain.Language, error)
	Update(language domain.Language) (*mongo.UpdateResult, error)
}
