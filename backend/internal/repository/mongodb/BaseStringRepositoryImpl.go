package mongodb

import (
	"context"
	"log"
	"uniovi-localizeme/constants"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/tools"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (b *BaseStringRepositoryImpl) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := b.GetCollection(b.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.DeleteBaseString, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (b *BaseStringRepositoryImpl) FindAll() (*[]domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := b.GetCollection(b.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	var baseStrings []domain.BaseString
	cursor, _ := collection.Find(context.TODO(), bson.D{})
	for cursor.Next(context.TODO()) {
		var baseString domain.BaseString
		if err := cursor.Decode(&baseString); err != nil {
			return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
		}
		baseStrings = append(baseStrings, baseString)
	}
	if err := cursor.Err(); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	if err := cursor.Close(context.TODO()); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &baseStrings, nil
}

func (b *BaseStringRepositoryImpl) FindByGroup(id primitive.ObjectID) (*[]domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := b.GetCollection(b.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"group._id": id}
	var baseStrings []domain.BaseString
	cursor, _ := collection.Find(context.TODO(), filter)
	for cursor.Next(context.TODO()) {
		var baseString domain.BaseString
		if err := cursor.Decode(&baseString); err != nil {
			return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
		}
		baseStrings = append(baseStrings, baseString)
	}
	if err = cursor.Err(); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	if err = cursor.Close(context.TODO()); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &baseStrings, nil
}

func (b *BaseStringRepositoryImpl) FindByLanguage(id primitive.ObjectID) (*[]domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := b.GetCollection(b.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"translations.language._id": id}
	var baseStrings []domain.BaseString
	cursor, _ := collection.Find(context.TODO(), filter)
	for cursor.Next(context.TODO()) {
		var baseString domain.BaseString
		if err := cursor.Decode(&baseString); err != nil {
			return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
		}
		baseStrings = append(baseStrings, baseString)
	}
	if err = cursor.Err(); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	if err = cursor.Close(context.TODO()); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &baseStrings, nil
}

func (b *BaseStringRepositoryImpl) FindByPermissions(id primitive.ObjectID) (*[]domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := b.GetCollection(b.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{
		"$or": []bson.M{
			{"author._id": id},
			{"group": nil},
			{"group.public": true},
			{"group.owner._id": id},
			{"group.permissions.user._id": id},
		},
	}
	var baseStrings []domain.BaseString
	cursor, _ := collection.Find(context.TODO(), filter)
	for cursor.Next(context.TODO()) {
		var baseString domain.BaseString
		if err = cursor.Decode(&baseString); err != nil {
			return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
		}
		baseStrings = append(baseStrings, baseString)
	}
	if err = cursor.Err(); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	if err = cursor.Close(context.TODO()); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &baseStrings, nil
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

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &baseString, nil
}

func (b *BaseStringRepositoryImpl) FindByIdentifier(identifier string) (*domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := b.GetCollection(b.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"identifier": bson.M{"$eq": identifier}}
	result := collection.FindOne(context.TODO(), filter)
	var baseString domain.BaseString
	if err = result.Decode(&baseString); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindBaseStringByIdentifier, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &baseString, nil
}

func (b *BaseStringRepositoryImpl) FindByIdentifierAndLanguageAndStage(identifier string, isoCode string, stageName string) (*domain.BaseString, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := b.GetCollection(b.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{
		"$and": []bson.M{
			{"identifier": identifier},
			{"translations.language.isoCode": isoCode},
			{"translations.stage.name": stageName},
		},
	}
	result := collection.FindOne(context.TODO(), filter)
	var baseString domain.BaseString
	if err = result.Decode(&baseString); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindBaseStringByIdentifierAndIsoCode, tools.GetCurrentFuncName())
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &baseString, nil
}

func (b *BaseStringRepositoryImpl) Update(baseString domain.BaseString) (*mongo.UpdateResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := b.GetCollection(b.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": baseString.ID}}
	update := bson.M{
		"$set": bson.M{
			"active":         baseString.Active,
			"author":         baseString.Author,
			"group":          baseString.Group,
			"page":           baseString.Page,
			"identifier":     baseString.Identifier,
			"sourceLanguage": baseString.SourceLanguage,
			"translations":   baseString.Translations,
		},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.UpdateBaseString, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}
