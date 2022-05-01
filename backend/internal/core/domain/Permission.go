package domain

// swagger:model Permission
type Permission struct {
	User     *User `bson:"user" json:"user"`
	CanWrite bool  `bson:"canWrite" json:"canWrite"`
}
