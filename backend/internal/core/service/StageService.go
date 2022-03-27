package service

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
)

type StageService interface {
	Create(request dto.StageRequest) (domain.Stage, error)
	FindAll() (*[]domain.Stage, error)
	Update(request domain.Stage) (*domain.Stage, error)
}
