package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"uniovi-localizeme/constants"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/tools"
)

type StageRepositoryImpl struct {
	GenericRepository[domain.Stage]
}

func CreateStageRepository() *StageRepositoryImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	repository := &StageRepositoryImpl{}
	repository.GenericRepository.Config = ConfigRepository{
		name:                 constants.Stages,
		createErrorMessage:   constants.InsertStage,
		findByIdErrorMessage: constants.FindStageById,
		deleteErrorMessage:   constants.DeleteStage,
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return repository
}

func (s *StageRepositoryImpl) FindByName(name string) (*domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := s.getCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"name": bson.M{"$eq": name}}
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
	collection, err := s.getCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": stage.ID}}
	update := bson.M{
		"$set": bson.M{
			"name":   stage.Name,
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
