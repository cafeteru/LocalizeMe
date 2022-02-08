package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// swagger:model Translation
type Translation struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Content    string             `bson:"content"`
	Language   Language           `bson:"language"`
	Version    int                `default:"1" bson:"version"`
	Active     bool               `bson:"active"`
	BaseString BaseString         `bson:"strings"`
	Author     User               `bson:"author"`
	Date       time.Time          `bson:"date"`
	Stage      Stage              `bson:"stage"`
}
