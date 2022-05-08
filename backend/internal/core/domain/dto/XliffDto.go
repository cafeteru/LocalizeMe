package dto

// swagger:model XliffDto
type XliffDto struct {
	SourceLanguageId string   `json:"sourceLanguageId"`
	TargetLanguageId string   `json:"targetLanguageId"`
	BaseStringIds    []string `json:"baseStringIds"`
}
