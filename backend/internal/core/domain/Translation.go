package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// swagger:model Translation
type Translation struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Content    string             `bson:"content" json:"content"`
	Language   Language           `bson:"language" json:"language"`
	Version    int                `default:"1" bson:"version" json:"version"`
	Active     bool               `bson:"active" json:"active"`
	BaseString BaseString         `bson:"baseString" json:"baseString"`
	Author     User               `bson:"author" json:"author"`
	Date       time.Time          `bson:"date" json:"date"`
	Stage      Stage              `bson:"stage" json:"stage"`
}
