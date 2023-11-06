package main

import (
	"log"
	"net/http"
)

func HandlerGetHome(w http.ResponseWriter, req *http.Request) {

	// fmt.Fprintf(w, "<h1>Hello</h1>")
	log.Println("Get Home Handler: ")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to Swath\n"))
	log.Println("Get Home Handler: Exiting")

}
