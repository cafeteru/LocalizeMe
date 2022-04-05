package mongodb

import (
	"context"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/constants"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/tools"
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

func (l *GroupRepositoryImpl) Create(group domain.Group) (*mongo.InsertOneResult, error) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	collection, err := l.GetCollection(l.name)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.CreateConnection, tools.GetCurrentFuncName())
	}
	result, err := collection.InsertOne(context.TODO(), group)
	if err != nil {
		return nil, tools.ErrorLogDetails(err, constants.InsertGroup, tools.GetCurrentFuncName())
	}
	l.CloseConnection()
	log.Printf("%s: end", tools.GetCurrentFuncName())
	return result, nil
}
