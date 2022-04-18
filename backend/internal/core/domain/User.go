package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// swagger:model User
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Admin    bool               `bson:"admin" json:"admin"`
	Active   bool               `bson:"active" json:"active"`
}

func (u *User) ClearPassword() {
	u.Password = ""
}
