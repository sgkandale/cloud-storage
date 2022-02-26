package server

import (
	"encoding/json"
	"net/http"

	"cloud_storage/object"
	"cloud_storage/user"
)

func SearchObjects(w http.ResponseWriter, r *http.Request) {

	decodedToken, err := user.DecodeToken(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	r.ParseForm()

	query := r.Form.Get("query")
	if len(query) < 3 {
		http.Error(w, "query should be atleast 3 characters long", http.StatusBadRequest)
		return
	}

	objects, err := object.SearchObjects(query, decodedToken.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objects)

}
