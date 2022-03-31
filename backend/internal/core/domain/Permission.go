package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:model Permission
type Permission struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	User          User               `bson:"user" json:"user"`
	Group         Group              `bson:"group" json:"group"`
	CanWriteGroup bool               `bson:"canWriteGroup" json:"canWriteGroup"`
}
