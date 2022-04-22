package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:model Group
type Group struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Owner       User               `bson:"owner" json:"owner"`
	Permissions []Permission       `bson:"permissions" json:"permissions"`
	Active      bool               `bson:"active" json:"active"`
	Public      bool               `bson:"public" json:"public"`
}
