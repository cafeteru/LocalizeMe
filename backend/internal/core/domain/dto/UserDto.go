package dto

// swagger:model UserDto
type UserDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
