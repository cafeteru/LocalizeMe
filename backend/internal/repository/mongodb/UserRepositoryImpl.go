package mongodb

import (
	"context"
	slog "github.com/go-eden/slf4go"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
	"go.mongodb.org/mongo-driver/bson"
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
	return result, err
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
	return &user, err
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
	if err = cursor.Close(context.TODO()); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	u.CloseConnection()
	slog.Debugf("%s: end", tools.GetCurrentFuncName())
	return &users, err
}
