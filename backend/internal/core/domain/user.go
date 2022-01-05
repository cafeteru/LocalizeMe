package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// swagger:model User
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Password string             `bson:"password"`
	Role     string             `bson:"role"`
	Active   bool               `bson:"active"`
}

func (user User) EncryptPassword() (User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
	return user, err
}

func (user User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
