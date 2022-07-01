package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:model BaseString
type BaseString struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SourceLanguage *Language          `bson:"sourceLanguage" json:"sourceLanguage"`
	Identifier     string             `bson:"identifier" json:"identifier"`
	Group          *Group             `bson:"group" json:"group" `
	Author         *User              `bson:"author" json:"author"`
	Active         bool               `bson:"active" json:"active"`
	Translations   []Translation      `bson:"translations" json:"translations"`
}

func (b BaseString) FindTranslationLastVersionByLanguage(language Language) string {
	if !language.Active {
		return ""
	}
	var translation Translation
	lastVersion := 0
	for _, value := range b.Translations {
		if value.Language.ID == language.ID && value.Version >= lastVersion && value.Active {
			translation = value
			lastVersion = value.Version
		}
	}
	return translation.Content
}

func (b BaseString) FindTranslationLastVersionByLanguageAndState(language Language, state Stage) string {
	if !language.Active || !state.Active {
		return ""
	}
	var translation Translation
	lastVersion := 0
	for _, value := range b.Translations {
		if value.Language != nil && value.Language.ID == language.ID &&
			value.Stage != nil && value.Stage.ID == state.ID &&
			value.Version >= lastVersion &&
			value.Active {
			translation = value
			lastVersion = value.Version
		}
	}
	return translation.Content
}
