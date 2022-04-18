package dto

// swagger:model PermissionDto
type PermissionDto struct {
	Email         string `json:"email"`
	CanWriteGroup bool   `json:"canWriteGroup"`
}
