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

type StageRepositoryImpl struct {
	name string
	AbstractRepository
}

func CreateStageRepository() *StageRepositoryImpl {
	return &StageRepositoryImpl{name: constants.Stages}
}

func (s *StageRepositoryImpl) Create(stage domain.Stage) (*mongo.InsertOneResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := s.GetCollection(s.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	result, err := collection.InsertOne(context.TODO(), stage)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.InsertStage, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (s *StageRepositoryImpl) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := s.GetCollection(s.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.DeleteStage, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (s *StageRepositoryImpl) FindAll() (*[]domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := s.GetCollection(s.name)
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

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &stages, nil
}

func (s *StageRepositoryImpl) FindById(id primitive.ObjectID) (*domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := s.GetCollection(s.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result := collection.FindOne(context.TODO(), filter)
	var stage domain.Stage
	if err = result.Decode(&stage); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindStageById, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &stage, nil
}

func (s *StageRepositoryImpl) FindByName(name string) (*domain.Stage, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := s.GetCollection(s.name)
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
	collection, err := s.GetCollection(s.name)
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
