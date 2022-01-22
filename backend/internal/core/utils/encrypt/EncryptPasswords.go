package encrypt

type Encrypt interface {
	EncryptPassword(password string) (string, error)
	CheckPassword(password string, encryptPassword string) bool
}
