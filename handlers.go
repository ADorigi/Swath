package main

import (
	"net/http"
)

func HandlerGetHome(w http.ResponseWriter, req *http.Request) {

	// fmt.Fprintf(w, "<h1>Hello</h1>")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to Swath"))

}
