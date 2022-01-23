package dto

// swagger:parameters CreateUser Login
type _ struct {
	// in: body
	// required: true
	User UserRequest
}

type UserRequest struct {
	Email    string
	Password string
}
