package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/internal/core/domain/dto"
)

type GroupService interface {
	Create(request dto.GroupDto) (domain.Group, error)
	Delete(id primitive.ObjectID, user *domain.User) (bool, error)
	Disable(id primitive.ObjectID, user *domain.User) (*domain.Group, error)
	FindAll() (*[]domain.Group, error)
	FindByPermissions(id primitive.ObjectID) (*[]domain.Group, error)
	FindCanWrite(id primitive.ObjectID) (*[]domain.Group, error)
	Update(group domain.Group, user *domain.User) (*domain.Group, error)
}
