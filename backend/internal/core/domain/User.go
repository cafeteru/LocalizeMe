package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// swagger:model User
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
	IsAdmin  bool               `bson:"isAdmin"`
	IsActive bool               `bson:"isActive"`
}
