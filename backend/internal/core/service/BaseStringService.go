package service

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/xliff"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseStringService interface {
	Create(request domain.BaseString, user *domain.User) (domain.BaseString, error)
	Delete(id primitive.ObjectID, user *domain.User) (bool, error)
	Disable(id primitive.ObjectID, user *domain.User) (*domain.BaseString, error)
	FindAll() (*[]domain.BaseString, error)
	FindByGroup(id primitive.ObjectID, user *domain.User) (*[]domain.BaseString, error)
	FindByPermissions(id primitive.ObjectID) (*[]domain.BaseString, error)
	Read(xliff xliff.Xliff, user *domain.User, stageId primitive.ObjectID, groupId primitive.ObjectID) (*[]domain.BaseString, error)
	Update(baseString domain.BaseString, user *domain.User) (*domain.BaseString, error)
}
