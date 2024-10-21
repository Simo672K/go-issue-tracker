package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(rawPassword string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	if err != nil {
		return nil, err
	}

	return hashedPassword, nil
}
