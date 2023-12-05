package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HandlerGetHome).Methods("GET")
	r.HandleFunc("/sanity", HandlerSanity).Methods("GET")
	r.HandleFunc("/getlinks", HandlerGetLinks).Methods("GET")

	log.Println("Starting to listen on port 9009")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./htmx/static/"))) //to server the contents of static folder in htmx
	// http.ListenAndServe(":9009", nil)
	// http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9009", r))
}
