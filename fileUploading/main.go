package main

import (
	"log"
	"net/http"
	"test/fileUploading/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/change/{ID}", handlers.ChangeProfileImage)
	r.Handle("/", r)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
