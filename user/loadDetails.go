package user

import (
	"context"
	"errors"
	"log"

	"cloud_storage/database"

	"github.com/jackc/pgx/v4"
)

func (user *User) LoadDetails() error {

	row := database.DB.QueryRow(
		context.TODO(),
		"SELECT password FROM users WHERE username = $1",
		user.Username,
	)

	err := row.Scan(
		&user.PasswordHash,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("item not found")
		}
		log.Println("user.LoadDetails() : ", err)
		return errors.New("something went wrong")
	}

	return nil
}
