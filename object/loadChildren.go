package object

import (
	"context"
	"errors"
	"log"
	"time"

	"cloud_storage/database"
)

func (object *Object) LoadChildren() error {

	rows, err := database.DB.Query(
		context.TODO(),
		`SELECT id, name, parent_folder_id, object_type, created_at, owner_id
		 FROM objects 
		 WHERE parent_folder_id = $1`,
		object.Id,
	)

	if err != nil {
		log.Println("object.LoadDetails() : ", err)
		return errors.New("something went wrong")
	}

	for rows.Next() {
		eachChildren := &Object{}

		err := rows.Scan(
			&eachChildren.Id,
			&eachChildren.Name,
			&eachChildren.ParentFolderId,
			&eachChildren.ObjectType,
			&eachChildren.CreatedAt,
			&eachChildren.OwnerId,
		)
		if err != nil {
			log.Println("object.LoadDetails() : ", err)
			return errors.New("something went wrong")
		}

		if eachChildren.ObjectType == "DISK" {
			continue
		}

		object.CreatedAtStr = time.Unix(object.CreatedAt, 0).Format("2006-01-02 15:04:05")

		object.Children = append(object.Children, eachChildren)
	}

	return nil

}
