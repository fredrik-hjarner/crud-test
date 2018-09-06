package main

import (
	"fmt" // format
	"net/http"
)

func routes(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/keys":
		keysRoute(w, r)
	case "/value":
		valueRoute(w, r)
	default:
		fmt.Fprintf(w, "Unrecognized path: %s", path)
	}
}
