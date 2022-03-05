package service

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
)

type UserService interface {
	Create(request dto.UserRequest) (domain.User, error)
	Delete(id string) (bool, error)
	Disable(id string) (*domain.User, error)
	FindAll() (*[]domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	Login(request dto.UserRequest) (*dto.TokenDto, error)
	Update(id string, request domain.User) (*domain.User, error)
}
