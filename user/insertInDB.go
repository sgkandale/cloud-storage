package user

import (
	"context"
	"errors"
	"log"
	"time"

	"cloud_storage/database"
)

func (user *User) InsertInDB() error {

	err := database.DB.QueryRow(
		context.TODO(),
		`INSERT INTO users
		(id, username, password, created_at)
		VALUES (nextval('users_id_seq'), $1, $2, $3)
		RETURNING id`,
		user.Username,
		user.PasswordHash,
		user.CreatedAt,
	).Scan(&user.Id)

	if err != nil {
		log.Println("user.InsertInDB() : ", err)
		return errors.New("something went wrong")
	}

	user.CreatedAtStr = time.Unix(user.CreatedAt, 0).Format("2006-01-02 15:04:05")

	return nil
}
