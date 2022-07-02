package dto

import "gitlab.com/HP-SCDS/Observatorio/2021-2022/localizeme/uniovi-localizeme/internal/core/domain"

// swagger:model XliffDto
type XliffDto struct {
	SourceLanguageId string       `json:"sourceLanguageId"`
	TargetLanguageId string       `json:"targetLanguageId"`
	BaseStringIds    []string     `json:"baseStringIds"`
	Stage            domain.Stage `json:"stage"`
}
