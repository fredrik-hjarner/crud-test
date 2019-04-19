package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fredrik-hjarner/ztorage/diskv"
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
	key := query.Get("key")
	log.Printf("key=%s", key)

	value := diskv.Diskv.ReadString(key)
	// if error
	if value == "" {
		fmt.Fprintf(w, "'%s' does not exist or value was an empty string", key)
	} else {
		fmt.Fprintf(w, "%s", value)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// body := r.Body
	query := r.URL.Query()
	key := query.Get("key")
	value := query.Get("value")
	if key == "" || value == "" {
		fmt.Fprintf(w, "Error: key and value is required.")
	} else {
		fmt.Fprintf(w, "Trying to store { %s: %s }", key, value)
		diskv.Diskv.WriteString(key, value)
	}
}

func delHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")
	if key == "" {
		fmt.Fprintf(w, "Error: key is required.")
	} else {
		fmt.Fprintf(w, "Trying to delete key=%s", key)
		diskv.Diskv.Erase(key)
	}
}
