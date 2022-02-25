package user

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (user *User) HashPassword() error {

	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordText), 8)
	if err != nil {
		log.Println("user.HashPassword() :", err)
		return errors.New("something went wrong")
	}

	user.PasswordHash = string(hash)

	return nil
}
