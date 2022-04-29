package service

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
)

type BaseStringService interface {
	Create(request domain.BaseString, user *domain.User) (domain.BaseString, error)
	FindAll() (*[]domain.BaseString, error)
	FindByPermissions(email string) (*[]domain.BaseString, error)
	Update(baseString domain.BaseString, user *domain.User) (*domain.BaseString, error)
}
