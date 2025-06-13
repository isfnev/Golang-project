package main

import (
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password []byte) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 10)
	return string(bytes), err
}
