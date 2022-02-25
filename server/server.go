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

	muxRouter.HandleFunc("/login", Login).Methods("OPTIONS", "POST")
	muxRouter.HandleFunc("/register", Register).Methods("OPTIONS", "POST")

	log.Println("starting server on port ", config.Config.Server.Port)

	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", config.Config.Server.Port),
			muxRouter,
		),
	)
}
