package object

import (
	"context"
	"errors"
	"log"
	"time"

	"cloud_storage/database"
)

func SearchObjects(query string, owner_id int64) ([]*Object, error) {

	rows, err := database.DB.Query(
		context.TODO(),
		`SELECT id, name, parent_folder_id, object_type, created_at, owner_id
		 FROM objects 
		 WHERE name LIKE '%`+query+`%' AND owner_id = $1`,
		// query,
		owner_id,
	)

	if err != nil {
		log.Println("object.SearchObjects() : ", err)
		return nil, errors.New("something went wrong")
	}

	var objects []*Object

	for rows.Next() {
		eachObject := &Object{}

		err := rows.Scan(
			&eachObject.Id,
			&eachObject.Name,
			&eachObject.ParentFolderId,
			&eachObject.ObjectType,
			&eachObject.CreatedAt,
			&eachObject.OwnerId,
		)
		if err != nil {
			log.Println("object.SearchObjects() : ", err)
			return nil, errors.New("something went wrong")
		}

		eachObject.CreatedAtStr = time.Unix(eachObject.CreatedAt, 0).Format("2006-01-02 15:04:05")

		objects = append(objects, eachObject)
	}

	return objects, nil

}
