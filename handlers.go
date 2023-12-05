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

	w.Write([]byte("Server is sane\n"))
	w.WriteHeader(http.StatusOK)

	log.Println("Sanity Handler: Exiting")

}

func HandlerGetLinks(w http.ResponseWriter, req *http.Request) {

	log.Println("GetLinks Handler: Enter")

	if req.Header.Get("HX-Request") == "true" {
		w.WriteHeader(http.StatusOK)
		log.Println("Sending Links")

		tmpl := template.Must(template.ParseFiles("htmx/links.html"))
		data := []string{
			"0.0.1",
			"0.0.3",
			"0.1.0",
			"0.1.1",
			"0.1.2",
			"0.1.3",
			"0.1.4",
		}
		tmpl.Execute(w, data)

	} else {

		w.WriteHeader(http.StatusForbidden)
		log.Println("Direct access to /getlinks forbidden")
	}

	log.Println("GetLinks Handler: Exiting")

}
