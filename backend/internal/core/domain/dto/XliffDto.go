package dto

import "uniovi-localizeme/internal/core/domain"

// swagger:model XliffDto
type XliffDto struct {
	SourceLanguageId string       `json:"sourceLanguageId"`
	TargetLanguageId string       `json:"targetLanguageId"`
	BaseStringIds    []string     `json:"baseStringIds"`
	Stage            domain.Stage `json:"stage"`
}
