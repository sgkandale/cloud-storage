package server

import (
	"fmt"
	"net/http"

	"cloud_storage/user"
)

func Login(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	username := r.Form.Get("username")
	password := r.Form.Get("password")

	userInstance, err := user.NewUserInstance(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = userInstance.LoadDetails()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// match password

	jwtoken, err := userInstance.EncodeJWT()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(
		[]byte(
			fmt.Sprintf(
				`{"token":"%s"}`,
				jwtoken,
			),
		),
	)
}
