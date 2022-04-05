package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:model BaseString
type BaseString struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SourceLanguage Language           `bson:"source_language" json:"sourceLanguage"`
	Description    string             `bson:"description" json:"description"`
	Group          Group              `bson:"group" json:"group"`
	Author         User               `bson:"author" json:"author"`
	Active         bool               `bson:"active" json:"active"`
}
