package object

import (
	"cloud_storage/database"
	"context"
	"errors"
	"log"
	"time"
)

func (object *Object) Save() error {

	err := database.DB.QueryRow(
		context.TODO(),
		`INSERT INTO objects (id, name, parent_folder_id, object_type, created_at, owner_id, file_link)
		 VALUES (nextval('objects_id_seq'), $1, $2, $3, $4, $5, $6)
		 RETURNING id`,
		object.Name,
		object.ParentFolderId,
		object.ObjectType,
		object.CreatedAt,
		object.OwnerId,
		object.FileLink,
	).Scan(&object.Id)

	if err != nil {
		log.Println("user.InsertInDB() : ", err)
		return errors.New("something went wrong")
	}

	object.CreatedAtStr = time.Unix(object.CreatedAt, 0).Format("2006-01-02 15:04:05")

	return nil
}
