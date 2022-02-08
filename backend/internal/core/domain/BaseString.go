package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:model BaseString
type BaseString struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	SourceLanguage Language           `bson:"sourceLanguage"`
	Description    string             `bson:"description"`
	Group          Group              `bson:"group"`
	Author         User               `bson:"author"`
}
