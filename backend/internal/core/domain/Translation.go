package domain

import (
	"time"
)

// swagger:model Translation
type Translation struct {
	Content  string    `bson:"content" json:"content"`
	Language *Language `bson:"language" json:"language"`
	Version  int       `default:"1" bson:"version" json:"version"`
	Active   bool      `bson:"active" json:"active"`
	Author   *User     `bson:"author" json:"author"`
	Date     time.Time `bson:"date" json:"date"`
	Stage    *Stage    `bson:"stage" json:"stage"`
}
