package mongodb

import (
	"context"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type LanguageRepositoryImpl struct {
	name string
	AbstractRepository
}

func CreateLanguageRepository() *LanguageRepositoryImpl {
	return &LanguageRepositoryImpl{name: "languages"}
}

func (l *LanguageRepositoryImpl) Create(language domain.Language) (*mongo.InsertOneResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := l.GetCollection(l.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	result, err := collection.InsertOne(context.TODO(), language)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.InsertLanguage, tools.GetCurrentFuncName())
	}
	l.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (l *LanguageRepositoryImpl) FindByIsoCode(isoCode string) (*domain.Language, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := l.GetCollection(l.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"IsoCode": bson.M{"$eq": isoCode}}
	result := collection.FindOne(context.TODO(), filter)
	var stage domain.Language
	if err = result.Decode(&stage); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindLanguageByIsoCode, tools.GetCurrentFuncName())
	}
	l.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &stage, nil
}
