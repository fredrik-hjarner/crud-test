package main

import (
	"net/http"

	"github.com/fredrik-hjarner/ztorage/routes"
)

func main() {
	httpHandler := routes.CreateHttpHandler()
	http.ListenAndServe(":9000", httpHandler)
}
