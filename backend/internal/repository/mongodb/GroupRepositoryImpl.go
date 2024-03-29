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
	name string
	AbstractRepository
}

func CreateGroupRepository() *GroupRepositoryImpl {
	return &GroupRepositoryImpl{name: constants.Groups}
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

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (g *GroupRepositoryImpl) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := g.GetCollection(g.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.DeleteGroup, tools.GetCurrentFuncName())
	}

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

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &groups, nil
}

func (g *GroupRepositoryImpl) FindByPermissions(id primitive.ObjectID) (*[]domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := g.GetCollection(g.name)
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

func (g *GroupRepositoryImpl) FindById(id primitive.ObjectID) (*domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := g.GetCollection(g.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result := collection.FindOne(context.TODO(), filter)
	var group domain.Group
	if err = result.Decode(&group); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindGroupById, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &group, nil
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

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &group, nil
}

func (g *GroupRepositoryImpl) FindCanWrite(id primitive.ObjectID) (*[]domain.Group, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := g.GetCollection(g.name)
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
	collection, err := g.GetCollection(g.name)
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
