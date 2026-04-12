package common_auth

import (
	"crypto/rand"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func Hash(secret string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(secret), 14)
	return string(bytes), err
}

func CompareHash(secret string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	return err == nil
}

func RandomPassword() (string, error) {
	length := 10
	b := make([]byte, length/2)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
