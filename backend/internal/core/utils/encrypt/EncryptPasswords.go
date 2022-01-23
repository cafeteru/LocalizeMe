package encrypt

type Encrypt interface {
	EncryptPassword(password string) (string, error)
	CheckPassword(encryptPassword, password string) bool
}
