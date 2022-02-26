package object

import (
	"context"
	"errors"
	"log"
	"time"

	"cloud_storage/database"

	"github.com/jackc/pgx/v4"
)

func (object *Object) LoadDetails() error {

	if object.Id < 1_000_000_000 {
		object.Name = "root"
		object.ObjectType = "Folder"
		object.Id = 1000000000

		return nil
	}

	row := database.DB.QueryRow(
		context.TODO(),
		`SELECT id, name, parent_folder_id, object_type, created_at, owner_id
		 FROM objects 
		 WHERE id = $1 LIMIT 1`,
		object.Id,
	)

	err := row.Scan(
		&object.Id,
		&object.Name,
		&object.ParentFolderId,
		&object.ObjectType,
		&object.CreatedAt,
		&object.OwnerId,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return errors.New("object not found")
		}
		log.Println("object.LoadDetails() : ", err)
		return errors.New("something went wrong")
	}

	object.CreatedAtStr = time.Unix(object.CreatedAt, 0).Format("2006-01-02 15:04:05")

	return nil

}
