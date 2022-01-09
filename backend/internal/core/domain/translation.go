package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// swagger:model Translation
type Translation struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Content  string             `bson:"content"`
	Language Language           `bson:"language"`
	Version  int                `default:"1" bson:"version"`
	Active   bool               `bson:"active"`
	Strings  Strings            `bson:"strings"`
}
