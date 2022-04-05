package repository

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type GroupRepository interface {
	Create(group domain.Group) (*mongo.InsertOneResult, error)
}
