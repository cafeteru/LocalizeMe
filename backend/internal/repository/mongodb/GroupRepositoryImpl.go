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

type GroupRepositoryImpl struct {
	name string
	AbstractRepository
}

func CreateGroupRepository() *GroupRepositoryImpl {
	return &GroupRepositoryImpl{name: "groups"}
}

func (g *GroupRepositoryImpl) Create(group domain.Group) (*mongo.InsertOneResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := g.GetCollection(g.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	result, err := collection.InsertOne(context.TODO(), group)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.InsertGroup, tools.GetCurrentFuncName())
	}
	g.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (g *GroupRepositoryImpl) FindAll() (*[]domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := g.GetCollection(g.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	var groups []domain.Group
	cursor, _ := collection.Find(context.TODO(), bson.D{})
	for cursor.Next(context.TODO()) {
		var group domain.Group
		if err := cursor.Decode(&group); err != nil {
			return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
		}
		groups = append(groups, group)
	}
	if err := cursor.Err(); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	if err := cursor.Close(context.TODO()); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	g.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &groups, nil
}

func (g *GroupRepositoryImpl) FindByPermissions(email string) (*[]domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := g.GetCollection(g.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{
		"$or": []bson.M{
			{"public": true},
			{"owner.email": email},
			{"permissions.user.email": email},
		},
	}
	var groups []domain.Group
	cursor, _ := collection.Find(context.TODO(), filter)
	for cursor.Next(context.TODO()) {
		var group domain.Group
		if err := cursor.Decode(&group); err != nil {
			return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
		}
		groups = append(groups, group)
	}
	if err := cursor.Err(); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	if err := cursor.Close(context.TODO()); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	g.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &groups, nil
}

func (g *GroupRepositoryImpl) FindByName(name string) (*domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := g.GetCollection(g.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"name": bson.M{"$eq": name}}
	result := collection.FindOne(context.TODO(), filter)
	var group domain.Group
	if err = result.Decode(&group); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindGroupByName, tools.GetCurrentFuncName())
	}
	g.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &group, nil
}
