package domain

// swagger:model Permission
type Permission struct {
	User          User `bson:"user" json:"user"`
	CanWriteGroup bool `bson:"canWriteGroup" json:"canWriteGroup"`
}
