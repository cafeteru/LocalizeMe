package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:model Group
type Group struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}
