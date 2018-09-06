package main

import (
	"fmt"
	"log"
	"net/http"
)

// value path

func keysRoute(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" { // TODO: replace with switch.
		keysGet(w, r)
	} else {
		// return only the value of `id`.
		fmt.Fprintf(w, "'%s' not allowed on /keys", method)
	}
}

/*
 * TODO: This should return all keys (i.e. folders) in alphabethical order.
 *
 */
func keysGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id") // TODO: this is wrong.
	log.Printf("value=%s", id)
	if id == "" {
		// return all key-value pairs.
		fmt.Fprintf(w, "Requested all values")
	} else {
		// return only the value of `id`.
		fmt.Fprintf(w, "Requested one value")
	}
}
