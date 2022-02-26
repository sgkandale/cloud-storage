package user

import (
	"fmt"
	"time"
)

func NewUserInstance(username, password string) (*User, error) {

	if username == "" {
		return nil, fmt.Errorf("username is empty")
	}
	if password == "" {
		return nil, fmt.Errorf("password is empty")
	}

	if len(username) < 3 || len(username) > 32 {
		return nil, fmt.Errorf("username length must be between 3 and 32")
	}
	if len(password) < 8 || len(password) > 32 {
		return nil, fmt.Errorf("password length must be between 8 and 32")
	}

	return &User{
		Username:     username,
		PasswordText: password,
		CreatedAt:    time.Now().Unix(),
		CreatedAtStr: time.Now().Format("2006-01-02 15:04:05"),
	}, nil
}
