package main

import (
	"net/http"

	"github.com/fredrik-hjarner/crud-test/routes"
	"github.com/fredrik-hjarner/crud-test/storage"
)

func main() {
	storage.Init()
	httpHandler := routes.CreateHTTPHandler()
	http.ListenAndServe(":9000", httpHandler)
}
