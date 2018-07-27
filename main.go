package main

import (
	"fmt"
	"net/http"

	"github.com/peterbourgon/diskv"
)

func routesHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/value":
		valueRoute(w, r)
	default:
		fmt.Fprintf(w, "Unrecognized path: %s", path)
	}
}

// Initialize a new diskv store, rooted at "my-data-dir", with a 1MB cache.
var d = diskv.New(diskv.Options{
	BasePath:     "my-data-dir",
	CacheSizeMax: 1024 * 1024,
})

func main() {
	// Write three bytes to the key "alpha".
	key := "alpha"
	d.Write(key, []byte{'1', '2', '3'})

	// Read the value back out of the store.
	value, _ := d.Read(key)
	fmt.Printf("%v\n", value)

	// Erase the key+value from the store (and the disk).
	// d.Erase(key)

	http.HandleFunc("/", routesHandler)

	http.ListenAndServe(":8080", nil)
}
