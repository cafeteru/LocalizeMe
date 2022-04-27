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

type BaseStringRepositoryImpl struct {
	name string
	AbstractRepository
}

func CreateBaseStringRepository() *BaseStringRepositoryImpl {
	return &BaseStringRepositoryImpl{name: constants.BaseStrings}
}

func (b *BaseStringRepositoryImpl) Create(group domain.BaseString) (*mongo.InsertOneResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := b.GetCollection(b.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	result, err := collection.InsertOne(context.TODO(), group)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.InsertBaseString, tools.GetCurrentFuncName())
	}
	b.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (b *BaseStringRepositoryImpl) FindById(id primitive.ObjectID) (*domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := b.GetCollection(b.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result := collection.FindOne(context.TODO(), filter)
	var baseString domain.BaseString
	if err = result.Decode(&baseString); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindBaseStringById, tools.GetCurrentFuncName())
	}
	b.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &baseString, nil
}

func (b *BaseStringRepositoryImpl) FindByIdentifier(name string) (*domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := b.GetCollection(b.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"identifier": bson.M{"$eq": name}}
	result := collection.FindOne(context.TODO(), filter)
	var baseString domain.BaseString
	if err = result.Decode(&baseString); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindBaseStringByIdentifier, tools.GetCurrentFuncName())
	}
	b.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &baseString, nil
}
