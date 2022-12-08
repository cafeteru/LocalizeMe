package dto

import "uniovi-localizeme/internal/core/domain"

// swagger:model GroupDto
type GroupDto struct {
	Name        string              `json:"name"`
	Owner       domain.User         `json:"owner"`
	Permissions []domain.Permission `json:"permissions"`
	Public      bool                `json:"public"`
}
