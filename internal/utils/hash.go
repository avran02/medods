package utils

import (
	"crypto/sha256"

	"golang.org/x/crypto/bcrypt"
)

func Hash(input string) (string, error) {
	sha256Hash := sha256.Sum256([]byte(input))

	bcryptHash, err := bcrypt.GenerateFromPassword(sha256Hash[:], bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bcryptHash), nil
}

func CompareHashAndPassword(password, hashedPassword string) error {
	sha256Hash := sha256.Sum256([]byte(password))

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), sha256Hash[:]); err != nil {
		return err
	}
	return nil
}
