package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:model Language
type Language struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	IsoCode     string             `bson:"IsoCode"`
	Description string             `bson:"Description"`
}
