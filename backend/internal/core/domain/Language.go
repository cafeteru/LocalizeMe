package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:model Language
type Language struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	IsoCode     string             `bson:"isoCode" json:"isoCode"`
	Description string             `bson:"description" json:"description"`
	Active      bool               `bson:"active" json:"active"`
}
