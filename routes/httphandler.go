package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// CreateHTTPHandler ...
func CreateHTTPHandler() http.Handler {
	// TODO:
	enableCORS := func(router http.Handler) http.Handler {
		headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
		originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
		methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
		return handlers.CORS(originsOk, headersOk, methodsOk)(router)
	}

	router := mux.NewRouter()

	///////////
	// users //
	///////////

	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUserByID).Methods("GET")

	///////////
	// value //
	///////////

	// valueRouter := router.Path("/value").Subrouter()

	// valueRouter.
	// 	Methods("GET").
	// 	Queries("namespace", "{namespace:.+}", "key", "{key:.+}").
	// 	HandlerFunc(GetValueByKey)

	// valueRouter.
	// 	Methods("POST").
	// 	Queries("namespace", "{namespace:.+}", "key", "{key:.+}", "value", "{value:.+}").
	// 	HandlerFunc(SetValue)

	// valueRouter.
	// 	Methods("DELETE").
	// 	Queries("namespace", "{namespace:.+}", "key", "{key:.+}").
	// 	HandlerFunc(DeleteOneValue)

	// valueRouter.
	// 	Methods("DELETE").
	// 	HandlerFunc(DeleteAllValues)

	//////////
	// root //
	//////////

	router.HandleFunc("/", RootGet).Methods("GET")

	return enableCORS(router)
}
