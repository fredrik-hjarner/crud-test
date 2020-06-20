package main

import (
	"net/http"
	"os"

	"github.com/fredrik-hjarner/ztorage/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// enable CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

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

	http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router))
}
