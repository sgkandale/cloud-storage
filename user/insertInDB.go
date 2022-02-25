package user

import (
	"context"
	"errors"
	"log"

	"cloud_storage/database"
)

func (user *User) InsertInDB() error {

	_, err := database.DB.Exec(
		context.TODO(),
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		user.Username,
		user.PasswordHash,
	)

	if err != nil {
		log.Println("user.InsertInDB() : ", err)
		return errors.New("something went wrong")
	}

	return nil
}
