package main

import (
	"html/template"
	"log"
	"net/http"
)

func HandlerGetHome(w http.ResponseWriter, req *http.Request) {

	log.Println("Get Home Handler: Enter")

	w.WriteHeader(http.StatusOK)
	// w.Write([]byte("Welcome to Swath\n"))
	tmpl := template.Must(template.ParseFiles("htmx/index.html"))
	tmpl.Execute(w, nil)

	log.Println("Get Home Handler: Exiting")

}

func HandlerSanity(w http.ResponseWriter, req *http.Request) {

	log.Println("Sanity Handler: Enter")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is sane\n"))

	log.Println("Sanity Handler: Exiting")

}
