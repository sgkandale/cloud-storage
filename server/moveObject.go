package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cloud_storage/object"
	"cloud_storage/user"

	"github.com/gorilla/mux"
)

func MoveObject(w http.ResponseWriter, r *http.Request) {

	decodedToken, err := user.DecodeToken(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	r.ParseForm()
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

	targetFolderId := r.Form.Get("targetFolderId")
	if targetFolderId == "" {
		http.Error(w, "targetFolderId is required", http.StatusBadRequest)
		return
	}

	targetFolderIdInt, err := strconv.ParseInt(targetFolderId, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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

	targetFolder, err := object.NewObjectById(targetFolderIdInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = targetFolder.LoadDetails()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if targetFolder.OwnerId != decodedToken.Id && targetFolder.OwnerId != 0 {
		http.Error(w, "you do not own the object", http.StatusUnauthorized)
		return
	}

	err = targetObject.ChangeParent(targetFolder.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(targetObject)

}
