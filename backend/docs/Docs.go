// Package docs LocalizeMe.
//
// Real-time management system for localization strings.
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//	Host: localhost:8080
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Security:
//	- Bearer: []
//
//	securityDefinitions:
//	  Bearer:
//	    type: apiKey
//	    name: Authorization
//	    in: header
//	    scheme: bearer
//	    bearerFormat: JWT
//
// swagger:meta
package docs

import (
	"uniovi-localizeme/internal/core/domain"
	"uniovi-localizeme/internal/core/domain/dto"
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

// swagger:parameters FindStageByName
type _ struct {
	// The stage´s name
	// in: path
	// required: true
	// type: string
	Name string `json:"name"`
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

// swagger:parameters DisableLanguage DeleteLanguage FindByLanguageBaseStrings
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

// swagger:parameters DisableGroup DeleteGroup FindByGroupBaseStrings
type _ struct {
	// The group´s id
	// in: path
	// required: true
	// type: string
	Id string `json:"id"`
}

// swagger:parameters CreateBaseString UpdateBaseString
type _ struct {
	// in: body
	// required: true
	BaseString domain.BaseString
}

// swagger:parameters DisableBaseString DeleteBaseString
type _ struct {
	// The baseString´s id
	// in: path
	// required: true
	// type: string
	Id string `json:"id"`
}

// swagger:parameters ReadXliff
type _ struct {
	// The stage´s id
	// in: query
	// required: true
	// type: string
	Stage string `json:"stage"`
	// The group´s id
	// in: query
	// required: true
	// type: string
	Group string `json:"group"`
}

// swagger:parameters CreateXliff
type _ struct {
	// in: body
	// required: true
	XliffDto dto.XliffDto
}

// swagger:model Xliff
type _ struct {
	// in: body
	File FileXml `json:"file"`
}

type FileXml struct {
	Units   []Unit  `json:"unit" xml:"unit"`
	Version float32 `xml:"version,attr" json:"version"`
	SrcLang string  `xml:"srcLang,attr" json:"srcLang"`
	TrgLang string  `xml:"trgLang,attr" json:"trgLang"`
}

type Unit struct {
	Id      string  `xml:"id,attr" json:"id"`
	Segment Segment `json:"segment" xml:"segment"`
}

type Segment struct {
	Source string `json:"source" xml:"source"`
	Target string `json:"target" xml:"target"`
}

// swagger:parameters FindByIdentifierBaseStrings
type _ struct {
	// The baseString´s isoCode
	// in: path
	// required: true
	// type: string
	IsoCode string `json:"isoCode"`
}
