package service

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LanguageService interface {
	Create(request dto.LanguageDto) (domain.Language, error)
	Disable(id primitive.ObjectID) (*domain.Language, error)
	FindAll() (*[]domain.Language, error)
	Update(request domain.Language) (*domain.Language, error)
}
