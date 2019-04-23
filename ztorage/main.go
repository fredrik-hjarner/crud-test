package main

import (
	"net/http"

	"github.com/fredrik-hjarner/ztorage/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//////////
	// keys //
	//////////

	router.HandleFunc("/keys", routes.KeysGet).Methods("GET")

	///////////
	// value //
	///////////

	router.HandleFunc("/value", routes.GetHandler).
		Methods("GET")

	router.HandleFunc("/value", routes.PostHandler).
		Methods("POST")

	router.HandleFunc("/value", routes.DeleteOneValue).
		Methods("DELETE").
		Queries("key", "{key}")

	router.HandleFunc("/value", routes.DeleteAllValues).
		Methods("DELETE")

	//////////
	// root //
	//////////

	router.HandleFunc("/", routes.RootGet).Methods("GET")

	http.ListenAndServe(":8080", router)
}
