package main

import (
	"fmt" // format
	"net/http"
	"strings"

	"github.com/peterbourgon/diskv"
)

func keyTransform(key string) *diskv.PathKey {
	path := strings.Split(key, "/")
	last := len(path) - 1
	return &diskv.PathKey{
		Path:     path[:last],
		FileName: path[last],
	}
}

// If you provide an AdvancedTransform, you must also provide its
// inverse:
func keyInverseTransform(pathKey *diskv.PathKey) (key string) {
	txt := pathKey.FileName[len(pathKey.FileName):]
	if txt != ".txt" {
		panic("Invalid file found in storage folder!")
	}
	return strings.Join(pathKey.Path, "/") + pathKey.FileName[:len(pathKey.FileName)]
}

// Initialize a new diskv store, rooted at "my-data-dir", with a 1MB cache.
var d = diskv.New(diskv.Options{
	BasePath:          "data-dir",
	AdvancedTransform: keyTransform,
	InverseTransform:  keyInverseTransform,
	CacheSizeMax:      1024 * 1024,
})

func main() {
	// Write three bytes to the key "alpha".
	key := "alpha/beta"
	d.Write(key, []byte{'1', '2', '3'})

	// Read the value back out of the store.
	value, _ := d.Read(key)
	fmt.Printf("%v\n", value)

	// Erase the key+value from the store (and the disk).
	// d.Erase(key)

	http.HandleFunc("/", routes)

	http.ListenAndServe(":8080", nil)
}
