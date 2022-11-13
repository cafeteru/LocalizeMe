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

type StageRepositoryImpl struct {
	generic.Repository[domain.Stage]
}

func CreateStageRepository() *StageRepositoryImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	repository := &StageRepositoryImpl{}
	repository.Repository.Config = generic.ConfigRepository{
		Name:                 constants.Stages,
		CreateErrorMessage:   constants.InsertStage,
		FindByIdErrorMessage: constants.FindStageById,
		DeleteErrorMessage:   constants.DeleteStage,
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return repository
}

func (s *StageRepositoryImpl) FindByName(name string) (*domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := s.GetCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"Name": bson.M{"$eq": name}}
	result := collection.FindOne(context.TODO(), filter)
	var stage domain.Stage
	if err = result.Decode(&stage); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindStageByName, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &stage, nil
}

func (s *StageRepositoryImpl) Update(stage domain.Stage) (*mongo.UpdateResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := s.GetCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": stage.ID}}
	update := bson.M{
		"$set": bson.M{
			"Name":   stage.Name,
			"active": stage.Active,
		},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.UpdateStage, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}
