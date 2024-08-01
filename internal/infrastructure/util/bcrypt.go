package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func ComparePassword(password string, passwordHashed []byte) error {
	return bcrypt.CompareHashAndPassword(passwordHashed, []byte(password))
}
