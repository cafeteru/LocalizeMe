package mongodb

import (
	"context"
	slog "github.com/go-eden/slf4go"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var name = "users"

type UserRepositoryImpl struct {
	AbstractRepository
}

func CreateUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Create(user domain.User) (*mongo.InsertOneResult, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection(name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.InsertUser, tools.GetCurrentFuncName())
	}
	u.CloseConnection()
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (u *UserRepositoryImpl) Delete(id string) (*mongo.DeleteResult, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection(name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bson.M{"$eq": objectID}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.DeleteUserByEmail, tools.GetCurrentFuncName())
	}
	u.CloseConnection()
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (u *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection(name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"email": bson.M{"$eq": email}}
	result := collection.FindOne(context.TODO(), filter)
	var user domain.User
	if err = result.Decode(&user); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindUserByEmail, tools.GetCurrentFuncName())
	}
	u.CloseConnection()
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return &user, nil
}

func (u *UserRepositoryImpl) FindById(id string) (*domain.User, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection(name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bson.M{"$eq": objectID}}
	result := collection.FindOne(context.TODO(), filter)
	var user domain.User
	if err = result.Decode(&user); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindUserById, tools.GetCurrentFuncName())
	}
	u.CloseConnection()
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return &user, nil
}

func (u *UserRepositoryImpl) FindAll() (*[]domain.User, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection(name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	var users []domain.User
	cursor, _ := collection.Find(context.TODO(), bson.D{})
	for cursor.Next(context.TODO()) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
		}
		user.Password = ""
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	if err := cursor.Close(context.TODO()); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	u.CloseConnection()
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return &users, nil
}

func (u *UserRepositoryImpl) Update(id string, user domain.User) (*mongo.UpdateResult, error) {
	slog.Debugf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection(name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bson.M{"$eq": objectID}}
	update := bson.M{
		"$set": bson.M{
			"email":    user.Email,
			"password": user.Password,
			"isActive": user.IsActive,
			"isAdmin":  user.IsAdmin,
		},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.UpdateUser, tools.GetCurrentFuncName())
	}
	u.CloseConnection()
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}
