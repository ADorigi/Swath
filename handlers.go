package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func HandlerGetHome(w http.ResponseWriter, req *http.Request) {

	log.Println("Get Home Handler: Enter")
	log.Println(os.Getenv("SWATH_VERSION"))
	w.WriteHeader(http.StatusOK)

	// w.Write([]byte("Welcome to Swath\n"))
	tmpl := template.Must(template.ParseFiles("htmx/index.html"))
	data := map[string]string{
		"Version": os.Getenv("SWATH_VERSION"),
	}
	tmpl.Execute(w, data)

	log.Println("Get Home Handler: Exiting")

}

func HandlerSanity(w http.ResponseWriter, req *http.Request) {

	log.Println("Sanity Handler: Enter")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Server is sane\n"))

	log.Println("Sanity Handler: Exiting")

}
