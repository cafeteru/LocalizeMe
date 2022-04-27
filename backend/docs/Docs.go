// Package docs LocalizeMe.
//
// Real-time management system for localization strings.
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: localhost:8080
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - Bearer: []
//
//     securityDefinitions:
//       Bearer:
//         type: apiKey
//         name: Authorization
//         in: header
//         scheme: bearer
//         bearerFormat: JWT
//
// swagger:meta
package docs

import (
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"
	"gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain/dto"
)

// swagger:parameters CreateUser Login
type _ struct {
	// in: body
	// required: true
	User dto.UserDto
}

// swagger:parameters UpdateMe UpdateUser
type _ struct {
	// in: body
	// required: true
	User domain.User
}

// swagger:parameters DeleteUser DisableUser FindUserById
type _ struct {
	// The user´s id
	// in: path
	// required: true
	// type: string
	Id string `json:"id"`
}

// swagger:parameters DisableStage DeleteStage
type _ struct {
	// The stage´s id
	// in: path
	// required: true
	// type: string
	Id string `json:"id"`
}

// swagger:parameters CreateStage
type _ struct {
	// in: body
	// required: true
	StageDto dto.StageDto
}

// swagger:parameters UpdateStage
type _ struct {
	// in: body
	// required: true
	Stage domain.Stage
}

// swagger:parameters CreateLanguage
type _ struct {
	// in: body
	// required: true
	LanguageDto dto.LanguageDto
}

// swagger:parameters UpdateLanguage
type _ struct {
	// in: body
	// required: true
	Language domain.Language
}

// swagger:parameters DisableLanguage DeleteLanguage
type _ struct {
	// The language´s id
	// in: path
	// required: true
	// type: string
	Id string `json:"id"`
}

// swagger:parameters CreateGroup
type _ struct {
	// in: body
	// required: true
	GroupDto dto.GroupDto
}

// swagger:parameters UpdateGroup
type _ struct {
	// in: body
	// required: true
	Group domain.Group
}

// swagger:parameters DisableGroup DeleteGroup
type _ struct {
	// The group´s id
	// in: path
	// required: true
	// type: string
	Id string `json:"id"`
}

// swagger:parameters CreateBaseString
type _ struct {
	// in: body
	// required: true
	BaseString domain.BaseString
}
