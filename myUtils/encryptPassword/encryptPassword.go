package encryptPassword

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(passwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("bcrypt.GenerateFromPassword error: %v", err)
	}
	return string(hash), nil
}

func CompareHashAndPassword(hashedPassword string, unverifiedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(unverifiedPassword))
	if err != nil {
		return false
	} else {
		return true
	}
}
