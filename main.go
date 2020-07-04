package main

import (
	"net/http"
	"os"

	"github.com/fredrik-hjarner/ztorage/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func enableCORS(router http.Handler) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	return handlers.CORS(originsOk, headersOk, methodsOk)(router)
}

func main() {
	router := mux.NewRouter()

	//////////
	// keys //
	//////////

	router.HandleFunc("/keys", routes.KeysGet).Methods("GET")

	///////////
	// value //
	///////////

	router.HandleFunc("/value", routes.GetValueByKey).
		Methods("GET").
		Queries("key", "{key:.+}")

	router.HandleFunc("/value", routes.SetValue).
		Methods("POST").
		Queries("key", "{key:.+}", "value", "{value:.+}")

	router.HandleFunc("/value", routes.DeleteOneValue).
		Methods("DELETE").
		Queries("key", "{key:.+}")

	router.HandleFunc("/value", routes.DeleteAllValues).
		Methods("DELETE")

	//////////
	// root //
	//////////

	router.HandleFunc("/", routes.RootGet).Methods("GET")

	http.ListenAndServe(":9000", enableCORS(router))
}
