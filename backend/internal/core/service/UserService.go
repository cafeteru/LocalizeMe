package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/internal/core/domain/dto"
)

type UserService interface {
	Create(user dto.UserDto) (domain.User, error)
	Delete(id primitive.ObjectID) (bool, error)
	Disable(id primitive.ObjectID) (*domain.User, error)
	FindAll() (*[]domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindById(id primitive.ObjectID) (*domain.User, error)
	Login(user dto.UserDto) (*dto.TokenDto, error)
	Update(request domain.User) (*domain.User, error)
}
