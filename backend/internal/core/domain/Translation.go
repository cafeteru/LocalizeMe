package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// swagger:model Translation
type Translation struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Content    string             `bson:"Content"`
	Language   Language           `bson:"Language"`
	Version    int                `default:"1" bson:"Version"`
	Active     bool               `bson:"Active"`
	BaseString BaseString         `bson:"BaseString"`
	Author     User               `bson:"Author"`
	Date       time.Time          `bson:"Date"`
	Stage      Stage              `bson:"Stage"`
}
