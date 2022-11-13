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

type UserRepositoryImpl struct {
	generic.Repository[domain.User]
}

func CreateUserRepository() *UserRepositoryImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	repository := &UserRepositoryImpl{}
	repository.Repository.Config = generic.ConfigRepository{
		Name:                 constants.Users,
		CreateErrorMessage:   constants.InsertUser,
		FindByIdErrorMessage: constants.FindUserById,
		DeleteErrorMessage:   constants.DeleteUser,
	}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return repository
}

func (u *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"email": bson.M{"$eq": email}}
	result := collection.FindOne(context.TODO(), filter)
	var user domain.User
	if err = result.Decode(&user); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindUserByEmail, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &user, nil
}

func (u *UserRepositoryImpl) Update(user domain.User) (*mongo.UpdateResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection()
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": user.ID}}
	update := bson.M{
		"$set": bson.M{
			"email":    user.Email,
			"password": user.Password,
			"active":   user.Active,
			"admin":    user.Admin,
		},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.UpdateUser, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}
