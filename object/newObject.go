package object

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func NewObject(name, objectType string, parent_folder_id, owner_id int64) (*Object, error) {

	if name == "" {
		return nil, errors.New("no name was provided")
	}
	if len(name) > 255 {
		return nil, errors.New("maximum name length supported is 255")
	}

	if parent_folder_id < 100_000_000 {
		return nil, errors.New("invalid parent_folder_id")
	}
	if owner_id < 100_000_000 {
		return nil, errors.New("invalid owner_id")
	}

	newObject := &Object{
		Name:           name,
		ParentFolderId: parent_folder_id,
		CreatedAt:      time.Now().Unix(),
		CreatedAtStr:   time.Now().Format("2006-01-02 15:04:05"),
		OwnerId:        owner_id,
	}

	if strings.EqualFold(objectType, "file") {
		newObject.ObjectType = "File"
		newObject.FileLink = "https://soaringeagle.biz/wp-content/uploads/2018/04/What-is-Cloud-Storage.png"
	} else if strings.EqualFold(objectType, "folder") {
		newObject.ObjectType = "Folder"
	} else {
		return nil, fmt.Errorf("unsupported object type : %s", objectType)
	}

	return newObject, nil
}
