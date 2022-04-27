package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:model BaseString
type BaseString struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	SourceLanguage Language           `bson:"sourceLanguage" json:"sourceLanguage"`
	Identifier     string             `bson:"identifier" json:"identifier"`
	Group          Group              `bson:"group" json:"group"`
	Author         User               `bson:"author" json:"author"`
	Active         bool               `bson:"active" json:"active"`
	Translations   []Translation      `bson:"translations" json:"translations"`
}
