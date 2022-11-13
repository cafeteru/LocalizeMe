package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"uniovi-localizeme/constants"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/tools"
)

type GroupRepositoryImpl struct {
	GenericRepository[domain.Group]
}

func CreateGroupRepository() *GroupRepositoryImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	repository := &GroupRepositoryImpl{}
	repository.GenericRepository.Config = ConfigRepository{
		name:                 constants.Groups,
		createErrorMessage:   constants.InsertGroup,
		findByIdErrorMessage: constants.FindGroupById,
		deleteErrorMessage:   constants.DeleteGroup,
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return repository
}

func (g *GroupRepositoryImpl) FindByPermissions(id primitive.ObjectID) (*[]domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := g.getCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{
		"$or": []bson.M{
			{"public": true},
			{"owner._id": id},
			{"permissions.user._id": id},
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

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &groups, nil
}

func (g *GroupRepositoryImpl) FindByName(name string) (*domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := g.getCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"name": bson.M{"$eq": name}}
	result := collection.FindOne(context.TODO(), filter)
	var group domain.Group
	if err = result.Decode(&group); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindGroupByName, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &group, nil
}

func (g *GroupRepositoryImpl) FindCanWrite(id primitive.ObjectID) (*[]domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := g.getCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{
		"$or": []bson.M{
			{"public": true},
			{"owner._id": id},
			{"$and": []bson.M{
				{"permissions.user._id": id},
				{"permissions.canWrite": true},
			}},
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

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &groups, nil
}

func (g *GroupRepositoryImpl) Update(group domain.Group) (*mongo.UpdateResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := g.getCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": group.ID}}
	update := bson.M{
		"$set": bson.M{
			"owner":       group.Owner,
			"name":        group.Name,
			"active":      group.Active,
			"permissions": group.Permissions,
			"Public":      group.Public,
		},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.UpdateLanguage, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}
