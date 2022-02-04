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

// swagger:parameters FindByEmail
type _ struct {
	// The user´s email
	// in: path
	// required: true
	// type: string
	Email string `json:"email"`
}
