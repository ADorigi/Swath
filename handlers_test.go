package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerGetHome(t *testing.T) {
	// create a test server
	s := httptest.NewServer(http.HandlerFunc(HandlerGetHome))

	// create a request to be made for this test
	req, err := http.NewRequest(http.MethodGet, s.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	// create the request and store the response
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	want := "Welcome to Swath\n"
	if string(body) != want {
		t.Errorf("Unexpected return. want: %s \t got: %s", want, string(body))
	}

}
