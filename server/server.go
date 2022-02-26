package server

import (
	"fmt"
	"log"
	"net/http"

	"cloud_storage/config"

	"github.com/gorilla/mux"
)

func StartServer() {

	muxRouter := mux.NewRouter()

	muxRouter.HandleFunc("/login", Login).Methods("POST")
	muxRouter.HandleFunc("/register", Register).Methods("POST")

	muxRouter.HandleFunc("/folder", GetObjectDetails).Methods("GET")
	muxRouter.HandleFunc("/folder/{objectId}", GetObjectDetails).Methods("GET")
	muxRouter.HandleFunc("/folder/{objectId}", DeleteObject).Methods("DELETE")
	muxRouter.HandleFunc("/folder/{objectId}", CreateFolder).Methods("POST")
	muxRouter.HandleFunc("/folder/{objectId}", MoveObject).Methods("PATCH")

	muxRouter.HandleFunc("/file", GetObjectDetails).Methods("GET")
	muxRouter.HandleFunc("/file/{objectId}", DeleteObject).Methods("DELETE")
	muxRouter.HandleFunc("/file/{parentId}", nil).Methods("POST")

	muxRouter.HandleFunc("/search", SearchObjects).Methods("POST")

	log.Println("starting server on port ", config.Config.Server.Port)

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", config.Config.Server.Port),
			muxRouter,
		),
	)
}
