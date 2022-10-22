package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/internal/core/domain/dto"
)

type LanguageService interface {
	Create(request dto.LanguageDto) (domain.Language, error)
	Delete(id primitive.ObjectID) (bool, error)
	Disable(id primitive.ObjectID) (*domain.Language, error)
	FindAll() (*[]domain.Language, error)
	Update(request domain.Language) (*domain.Language, error)
}
