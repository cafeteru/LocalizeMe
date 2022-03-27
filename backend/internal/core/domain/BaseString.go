package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:model BaseString
type BaseString struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	SourceLanguage Language           `bson:"SourceLanguage"`
	Description    string             `bson:"Description"`
	Group          Group              `bson:"Group"`
	Author         User               `bson:"Author"`
}
