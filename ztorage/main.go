package main

import (
	"net/http"

	"github.com/fredrik-hjarner/ztorage/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// keys
	r.HandleFunc("/keys", routes.KeysGet).Methods("GET")

	// value
	r.HandleFunc("/value", routes.GetHandler).Methods("GET")
	r.HandleFunc("/value", routes.PostHandler).Methods("POST")
	r.HandleFunc("/value", routes.DelHandler).Methods("DELETE")

	// route
	r.HandleFunc("/", routes.RootGet).Methods("GET")

	http.ListenAndServe(":8080", r)
}
