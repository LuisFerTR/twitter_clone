package db

import "golang.org/x/crypto/bcrypt"

// EncryptPassword encrypts the given password
func EncryptPassword(password string) (string, error) {
	cost := 8
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(hashedPassword), err
}
