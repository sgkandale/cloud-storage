package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cloud_storage/object"
	"cloud_storage/user"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {

	decodedToken, err := user.DecodeToken(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var maxSize int64 = 2048000

	err = r.ParseMultipartForm(maxSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	parentFolderId := r.Form.Get("parentFolderId")
	if parentFolderId == "" {
		http.Error(w, "parentFolderId is required", http.StatusBadRequest)
		return
	}

	parentFolderIdInt, err := strconv.ParseInt(parentFolderId, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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

	fileName := r.Form.Get("fileName")
	if fileName == "" {
		http.Error(w, "fileName is requierd", http.StatusBadRequest)
		return
	}
	_, _, err = r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newFile, err := object.NewObject(fileName, "File", parentFolder.Id, decodedToken.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if parentFolder.HaveChildren(newFile.Name) {
		http.Error(w, "file already exists", http.StatusBadRequest)
		return
	}

	err = newFile.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newFile)

}
