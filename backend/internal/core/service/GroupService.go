package service

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupService interface {
	Create(request dto.GroupDto) (domain.Group, error)
	Delete(id primitive.ObjectID, user *domain.User) (bool, error)
	Disable(id primitive.ObjectID, user *domain.User) (*domain.Group, error)
	FindAll() (*[]domain.Group, error)
	FindByPermissions(email string) (*[]domain.Group, error)
	FindCanWrite(email string) (*[]domain.Group, error)
	Update(group domain.Group, user *domain.User) (*domain.Group, error)
}
