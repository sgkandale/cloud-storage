package object

import (
	"context"
	"errors"
	"log"

	"cloud_storage/database"
)

func (object *Object) ChangeParent(newparent int64) error {

	_, err := database.DB.Exec(
		context.TODO(),
		`UPDATE objects SET parent_folder_id = $1 WHERE id = $2`,
		newparent,
		object.Id,
	)

	if err != nil {
		log.Println("object.ChangeParent() : ", err)
		return errors.New("something went wrong")
	}

	object.ParentFolderId = newparent

	return nil
}
