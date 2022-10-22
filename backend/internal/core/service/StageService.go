package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/internal/core/domain/dto"
)

type StageService interface {
	Create(request dto.StageDto) (*domain.Stage, error)
	Delete(id primitive.ObjectID) (bool, error)
	Disable(id primitive.ObjectID) (*domain.Stage, error)
	FindAll() (*[]domain.Stage, error)
	FindByName(name string) (*domain.Stage, error)
	Update(request domain.Stage) (*domain.Stage, error)
}
