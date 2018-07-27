package main

import (
	"fmt"
	"log"
	"net/http"
)

// value path

func valueRoute(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" { // TODO: replace with switch.
		getHandler(w, r)
	} else if method == "POST" {
		postHandler(w, r)
	} else if method == "DEL" {
		postHandler(w, r)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	log.Printf("value=%s", id)
	if id == "" {
		// return all key-value pairs.
		fmt.Fprintf(w, "Requested all values")
	} else {
		// return only the value of `id`.
		fmt.Fprintf(w, "Requested one value")
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// body := r.Body
	query := r.URL.Query()
	id := query.Get("id")
	value := query.Get("value")
	if id == "" || value == "" {
		fmt.Fprintf(w, "Error: id and value is required.")
	} else {
		fmt.Fprintf(w, "Trying to store { %s: %s }", id, value)
	}
}

func delHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	if id == "" {
		fmt.Fprintf(w, "Error: id is required.")
	} else {
		fmt.Fprintf(w, "Trying to delete id=%s", id)
		d.Erase(id)
	}
}
