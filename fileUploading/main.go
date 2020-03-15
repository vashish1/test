package main

import (
	"log"
	"net/http"
	"test/fileUploading/handlers"
	// "test/fileUploading/database"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/change/{ID}", handlers.ChangeProfileImage).Methods("GET","POST")
	http.Handle("/", r)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}