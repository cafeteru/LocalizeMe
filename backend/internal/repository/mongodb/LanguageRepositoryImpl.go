package mongodb

import (
	"context"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type LanguageRepositoryImpl struct {
	name string
	AbstractRepository
}

func CreateLanguageRepository() *LanguageRepositoryImpl {
	return &LanguageRepositoryImpl{name: constants.Languages}
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
	// l.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (l *LanguageRepositoryImpl) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := l.GetCollection(l.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.DeleteLanguage, tools.GetCurrentFuncName())
	}
	// l.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (l *LanguageRepositoryImpl) FindAll() (*[]domain.Language, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := l.GetCollection(l.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	var languages []domain.Language
	cursor, _ := collection.Find(context.TODO(), bson.D{})
	for cursor.Next(context.TODO()) {
		var language domain.Language
		if err := cursor.Decode(&language); err != nil {
			return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
		}
		languages = append(languages, language)
	}
	if err := cursor.Err(); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	if err := cursor.Close(context.TODO()); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	// l.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &languages, nil
}

func (l *LanguageRepositoryImpl) FindById(id primitive.ObjectID) (*domain.Language, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := l.GetCollection(l.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result := collection.FindOne(context.TODO(), filter)
	var language domain.Language
	if err = result.Decode(&language); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindLanguageById, tools.GetCurrentFuncName())
	}
	// l.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &language, nil
}

func (l *LanguageRepositoryImpl) FindByIsoCode(isoCode string) (*domain.Language, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := l.GetCollection(l.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"isoCode": bson.M{"$eq": isoCode}}
	var language domain.Language
	if err = collection.FindOne(context.TODO(), filter).Decode(&language); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindLanguageByIsoCode, tools.GetCurrentFuncName())
	}
	// l.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &language, nil
}

func (l *LanguageRepositoryImpl) Update(language domain.Language) (*mongo.UpdateResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := l.GetCollection(l.name)
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
	// l.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}
