package main

import (
	"net/http"

	"github.com/fredrik-hjarner/ztorage/routes"
)

func main() {
	// Erase the key+value from the store (and the disk).
	// d.Erase(key)

	http.HandleFunc("/", routes.Routes)

	http.ListenAndServe(":8080", nil)
}
