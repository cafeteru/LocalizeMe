package service

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
