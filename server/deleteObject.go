package server

import (
	"fmt"
	"net/http"
	"strconv"

	"cloud_storage/object"
	"cloud_storage/user"

	"github.com/gorilla/mux"
)

func DeleteObject(w http.ResponseWriter, r *http.Request) {

	decodedToken, err := user.DecodeToken(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	objectId := vars["objectId"]

	if objectId == "" {
		http.Error(w, "objectId is required", http.StatusBadRequest)
		return
	}
	objectIdInt, err := strconv.ParseInt(objectId, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	object, err := object.NewObjectById(objectIdInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = object.LoadDetails()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if object.OwnerId != decodedToken.Id {
		http.Error(w, "you do not own the object", http.StatusUnauthorized)
		return
	}

	err = object.Delete()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"message":"object deleted","objectId":%d}`, objectIdInt)))

}
