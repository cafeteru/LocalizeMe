package service

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	Create(request dto.UserRequest) (domain.User, error)
	Delete(id primitive.ObjectID) (bool, error)
	Disable(id primitive.ObjectID) (*domain.User, error)
	FindAll() (*[]domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Login(request dto.UserRequest) (*dto.TokenDto, error)
	Update(id primitive.ObjectID, request domain.User) (*domain.User, error)
}
