package encrypt

import "golang.org/x/crypto/bcrypt"

type Encrypter struct{}

func (e Encrypter) Encrypt(s string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}
