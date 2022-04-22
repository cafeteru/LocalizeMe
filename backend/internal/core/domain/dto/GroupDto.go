package dto

import "gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"

// swagger:model GroupDto
type GroupDto struct {
	Name        string          `json:"name"`
	Owner       domain.User     `json:"owner"`
	Permissions []PermissionDto `json:"permissions"`
	Public      bool            `json:"public"`
}
