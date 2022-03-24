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

type StageRepositoryImpl struct {
	name string
	AbstractRepository
}

func CreateStageRepository() *StageRepositoryImpl {
	return &StageRepositoryImpl{name: "stages"}
}

func (u *StageRepositoryImpl) Create(stage domain.Stage) (*mongo.InsertOneResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection(u.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	result, err := collection.InsertOne(context.TODO(), stage)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.InsertStage, tools.GetCurrentFuncName())
	}
	u.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (u *StageRepositoryImpl) FindAll() (*[]domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection(u.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	var stages []domain.Stage
	cursor, _ := collection.Find(context.TODO(), bson.D{})
	for cursor.Next(context.TODO()) {
		var stage domain.Stage
		if err := cursor.Decode(&stage); err != nil {
			return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
		}
		stages = append(stages, stage)
	}
	if err := cursor.Err(); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	if err := cursor.Close(context.TODO()); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	u.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &stages, nil
}

func (u *StageRepositoryImpl) FindByName(name string) (*domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection(u.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"name": bson.M{"$eq": name}}
	result := collection.FindOne(context.TODO(), filter)
	var stage domain.Stage
	if err = result.Decode(&stage); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindStageByName, tools.GetCurrentFuncName())
	}
	u.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &stage, nil
}
