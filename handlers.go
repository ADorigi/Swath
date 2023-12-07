/*
###### Handler template ######
func Handler_____(w http.ResponseWriter, req *http.Request) {

	log.Println("_____ Handler: Enter")

	if req.Header.Get("HX-Request") == "true" {
		w.WriteHeader(http.StatusOK)

		tmpl := template.Must(template.ParseFiles("_____"))
		tmpl.Execute(w, nil)

	} else {

		w.WriteHeader(http.StatusForbidden)
		log.Println("Direct access to /_____ forbidden")
	}

	log.Println("____ Handler: Exiting")

}
*/
package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
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

func HandlerSanityHTMX(w http.ResponseWriter, req *http.Request) {

	log.Println("Sanity HTMX Handler: Enter")

	if req.Header.Get("HX-Request") == "true" {
		w.WriteHeader(http.StatusOK)

		tmpl := template.Must(template.ParseFiles("htmx/sanity.html"))
		tmpl.Execute(w, nil)

	} else {

		w.WriteHeader(http.StatusForbidden)
		log.Println("Direct access to /sanityhtmx forbidden")
	}

	log.Println("Sanity HTMX Handler: Exiting")

}

func HandlerAboutSwath(w http.ResponseWriter, req *http.Request) {

	log.Println("About Swath Handler: Enter")

	if req.Header.Get("HX-Request") == "true" {
		w.WriteHeader(http.StatusOK)

		// tmpl := template.Must(template.ParseFiles("htmx/spinner.html"))
		// tmpl.Execute(w, nil)
		time.Sleep(500 * time.Millisecond) //// simulating artificial delay to show the spinner
		tmpl := template.Must(template.ParseFiles("htmx/about.html"))
		tmpl.Execute(w, nil)

	} else {

		w.WriteHeader(http.StatusForbidden)
		log.Println("Direct access to /aboutswath forbidden")
	}

	log.Println("About Swath Handler: Exiting")

}

// func HandlerGetLinks(w http.ResponseWriter, req *http.Request) {

// 	log.Println("GetLinks Handler: Enter")

// 	if req.Header.Get("HX-Request") == "true" {
// 		w.WriteHeader(http.StatusOK)
// 		log.Println("Sending Links")

// 		tmpl := template.Must(template.ParseFiles("htmx/links.html"))
// 		data := []string{
// 			"0.0.1",
// 			"0.0.3",
// 			"0.1.0",
// 			"0.1.1",
// 			"0.1.2",
// 			"0.1.3",
// 			"0.1.4",
// 		}
// 		tmpl.Execute(w, data)

// 	} else {

// 		w.WriteHeader(http.StatusForbidden)
// 		log.Println("Direct access to /getlinks forbidden")
// 	}

// 	log.Println("GetLinks Handler: Exiting")

// }
