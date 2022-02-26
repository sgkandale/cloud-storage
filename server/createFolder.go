package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cloud_storage/object"
	"cloud_storage/user"
)

func CreateFolder(w http.ResponseWriter, r *http.Request) {

	decodedToken, err := user.DecodeToken(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	r.ParseForm()

	parentFolderIdInt := decodedToken.Id
	parentFolderId := r.Form.Get("parentFolderId")

	if parentFolderId != "" {
		parentFolderIdInt, err = strconv.ParseInt(parentFolderId, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	parentFolder, err := object.NewObjectById(parentFolderIdInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = parentFolder.LoadDetails()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if parentFolder.OwnerId != decodedToken.Id && parentFolder.OwnerId != 0 {
		http.Error(w, "you do not own the parent folder", http.StatusUnauthorized)
		return
	}

	if parentFolder.OwnerId == 0 {
		parentFolder.OwnerId = decodedToken.Id
	}

	err = parentFolder.LoadChildren()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newFolder, err := object.NewObject(r.Form.Get("folderName"), "Folder", parentFolder.Id, decodedToken.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if parentFolder.HaveChildren(newFolder.Name) {
		http.Error(w, "folder already exists", http.StatusBadRequest)
		return
	}

	err = newFolder.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newFolder)
}
