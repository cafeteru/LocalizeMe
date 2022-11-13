package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"uniovi-localizeme/constants"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/internal/repository/mongodb/generic"
	"uniovi-localizeme/tools"
)

type LanguageRepositoryImpl struct {
	generic.Repository[domain.Language]
}

func CreateLanguageRepository() *LanguageRepositoryImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	repository := &LanguageRepositoryImpl{}
	repository.Repository.Config = generic.ConfigRepository{
		Name:                 constants.Languages,
		CreateErrorMessage:   constants.InsertLanguage,
		FindByIdErrorMessage: constants.FindLanguageById,
		DeleteErrorMessage:   constants.DeleteLanguage,
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return repository
}

func (l *LanguageRepositoryImpl) FindByIsoCode(isoCode string) (*domain.Language, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := l.GetCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"isoCode": bson.M{"$eq": isoCode}}
	var language domain.Language
	if err = collection.FindOne(context.TODO(), filter).Decode(&language); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindLanguageByIsoCode, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &language, nil
}

func (l *LanguageRepositoryImpl) Update(language domain.Language) (*mongo.UpdateResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := l.GetCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": language.ID}}
	update := bson.M{
		"$set": bson.M{
			"description": language.Description,
			"isoCode":     language.IsoCode,
			"active":      language.Active,
		},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.UpdateLanguage, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}
