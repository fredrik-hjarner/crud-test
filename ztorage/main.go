package main

import (
	"fmt" // format
	"net/http"

	"github.com/fredrik-hjarner/ztorage/diskv"
	"github.com/fredrik-hjarner/ztorage/routes"
)

func main() {
	// Write three bytes to the key "alpha".
	key := "alpha/beta"
	diskv.Diskv.Write(key, []byte{'1', '2', '3'})

	// Read the value back out of the store.
	value, _ := diskv.Diskv.Read(key)
	fmt.Printf("%v\n", value)

	// Erase the key+value from the store (and the disk).
	// d.Erase(key)

	http.HandleFunc("/", routes.Routes)

	http.ListenAndServe(":8080", nil)
}
