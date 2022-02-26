package object

import (
	"errors"
)

func NewObjectById(object_id int64) (*Object, error) {

	if object_id < 100_000_000 {
		return nil, errors.New("invalid object id")
	}

	return &Object{
		Id: object_id,
	}, nil
}
