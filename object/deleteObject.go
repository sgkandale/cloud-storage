package object

import (
	"context"
	"errors"
	"log"

	"cloud_storage/database"
)

func (object *Object) Delete() error {

	if object.Id < 1_000_000_000 {
		return errors.New("root object cannot be deleted")
	}

	_, err := database.DB.Exec(
		context.TODO(),
		`DELETE FROM objects WHERE id = $1`,
		object.Id,
	)

	if err != nil {
		log.Println("object.Delete() : ", err)
		return errors.New("something went wrong")
	}

	return nil

}
