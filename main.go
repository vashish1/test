package main

import (
	"log"
	"net/http"
	"test/common"

	"test/common/datastore"
	"test/handlers"

	"github.com/gorilla/mux"
)

const (
	//WEBSERVERPORT = ":8443"
	WEBSERVERPORT = ":3000"
)

func main() {

	db, err := datastore.NewDatastore("localhost:27017")

	if err != nil {
		log.Print(err)
	}

	defer db.Close()

	env := common.Env{}

	env.DB = db

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HomeHandler)
	r.Handle("/signup", handlers.SignUpHandler(&env)).Methods("GET", "POST")

	err = http.ListenAndServe(WEBSERVERPORT, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
