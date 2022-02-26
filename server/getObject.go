package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cloud_storage/object"
	"cloud_storage/user"

	"github.com/gorilla/mux"
)

func GetObjectDetails(w http.ResponseWriter, r *http.Request) {

	decodedToken, err := user.DecodeToken(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	objectIdInt := decodedToken.Id
	vars := mux.Vars(r)
	objectId := vars["objectId"]

	if objectId != "" {
		objectIdInt, err = strconv.ParseInt(objectId, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	targetObject, err := object.NewObjectById(objectIdInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = targetObject.LoadDetails()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if targetObject.OwnerId != decodedToken.Id && targetObject.OwnerId != 0 {
		http.Error(w, "you do not own the object", http.StatusUnauthorized)
		return
	}

	if targetObject.OwnerId == 0 {
		targetObject.OwnerId = decodedToken.Id
	}

	err = targetObject.LoadChildren()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(targetObject)

}
