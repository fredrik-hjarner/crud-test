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
	/*
	 * TODO: Wtf is this?????
	 * Should "id" be key????
	 */
	key := query.Get("key")
	log.Printf("key=%s", key)

	value, err := d.Read(key)
	// if error
	if err != nil {
		fmt.Fprintf(w, "'%s' does not exist", key)
	} else {
		fmt.Fprintf(w, "%s", value)
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
