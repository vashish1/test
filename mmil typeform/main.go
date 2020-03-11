package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func main(){
	r:=mux.NewRouter()
	r.HandleFunc("/Registeration",register)
	http.Handle("/",r)
	http.ListenAndServe(":9090",nil)
}

func register(w http.ResponseWriter,r *http.Request){
	t, err := template.ParseFiles("C:/Users/yashi/go/src/test/validationkit/test.html")
	//fmt.Println("anushiv testing 3")
	if err != nil {
		log.Fatal("Could not parse template files:", err)
	}
	er := t.Execute(w,"")
	if er != nil {
		log.Fatal("could not execute the files\n:", er)
	}
}