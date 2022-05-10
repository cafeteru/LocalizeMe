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

type UserRepositoryImpl struct {
	name string
	AbstractRepository
}

func CreateUserRepository() *UserRepositoryImpl {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	repository := &UserRepositoryImpl{name: constants.Users}
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return repository
}

func (u *UserRepositoryImpl) Create(user domain.User) (*mongo.InsertOneResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection(u.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.InsertUser, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (u *UserRepositoryImpl) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection(u.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.DeleteUser, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}

func (u *UserRepositoryImpl) FindByEmail(email string) (*domain.User, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection(u.name)
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

func (u *UserRepositoryImpl) FindById(id primitive.ObjectID) (*domain.User, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection(u.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result := collection.FindOne(context.TODO(), filter)
	var user domain.User
	if err = result.Decode(&user); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.FindUserById, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &user, nil
}

func (u *UserRepositoryImpl) FindAll() (*[]domain.User, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection(u.name)
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
		user.ClearPassword()
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}
	if err := cursor.Close(context.TODO()); err != nil {
		return nil, tools.ErrorLogDetails(err, constants.ReadDatabase, tools.GetCurrentFuncName())
	}

	log.Printf("%s: end", tools.GetCurrentFuncName())
	return &users, nil
}

func (u *UserRepositoryImpl) Update(user domain.User) (*mongo.UpdateResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := u.GetCollection(u.name)
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
