package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:model Language
type Permission struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	User          User               `bson:"user"`
	Group         Group              `bson:"group"`
	CanWriteGroup bool               `bson:"canWriteGroup"`
}
