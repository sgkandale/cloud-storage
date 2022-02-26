package object

import (
	"context"
	"errors"
	"log"
	"time"

	"cloud_storage/database"
)

func CreateUserRoot(userid int64) error {

	_, err := database.DB.Exec(
		context.TODO(),
		`INSERT INTO objects 
		(id, name, parent_folder_id, object_type, created_at, owner_id)
		VALUES (nextval('objects_id_seq'), 'root', $1, 'Folder', $2, $3)`,
		0,
		time.Now().Unix(),
		userid,
	)

	if err != nil {
		log.Println("object.CreateUserRoot() : ", err)
		return errors.New("something went wrong")
	}

	return nil
}
