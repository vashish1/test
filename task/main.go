package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var done chan string
var state int

const(
	running=0
	stopping=1
	pausing=2
	resuming=3
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/start", start).Methods("GET")
	r.HandleFunc("/stop", stopit).Methods("GET")
	r.HandleFunc("/pause", pauseit).Methods("GET")
	r.HandleFunc("/resume", resumeit).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func process(){
	fmt.Println("working")
	//do rest of the work here
}

func run(done chan string) {
let:
	for {
		switch state{
		case 1:
			fmt.Println("done")
			break let
		case 2:
			continue
		case 3:
			setstate(running)		
		case 0:
           process() 	
	}
}
}
func setstate(st int){
	state=st
}

func start(w http.ResponseWriter,r *http.Request){
	  setstate(running)
	  fmt.Print("1")
	  go run(done)
	  fmt.Print("2")
}

func stopit(w http.ResponseWriter,r *http.Request){
	setstate(stopping)
	fmt.Println("stopping state")
}

func pauseit(w http.ResponseWriter,r *http.Request){
	setstate(pausing)
	fmt.Println("pausing state")
}

func resumeit(w http.ResponseWriter,r *http.Request){
	setstate(resuming)
	fmt.Println("resuming state")
}