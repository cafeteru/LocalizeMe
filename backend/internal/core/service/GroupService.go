package service

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
)

type GroupService interface {
	Create(request dto.GroupDto) (domain.Group, error)
	FindAll() (*[]domain.Group, error)
	FindByPermissions(email string) (*[]domain.Group, error)
}
