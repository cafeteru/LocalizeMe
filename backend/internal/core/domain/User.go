package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// swagger:model User
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"Email"`
	Password string             `bson:"Password"`
	Admin    bool               `bson:"Admin"`
	Active   bool               `bson:"Active"`
}
