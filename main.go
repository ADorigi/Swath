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
	log.Println("Starting to listen on port 9009")
	// http.ListenAndServe(":9009", nil)
	// http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9009", r))
}
