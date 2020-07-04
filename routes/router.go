package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	enableCORS := func(router http.Handler) http.Handler {
		headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
		originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
		methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
		return handlers.CORS(originsOk, headersOk, methodsOk)(router)
	}

	router := mux.NewRouter()

	//////////
	// keys //
	//////////

	router.HandleFunc("/keys", KeysGet).Methods("GET")

	///////////
	// value //
	///////////

	valueRouter := router.Path("/value").Subrouter()

	valueRouter.
		Methods("GET").
		Queries("key", "{key:.+}").
		HandlerFunc(GetValueByKey)

	valueRouter.
		Methods("POST").
		Queries("key", "{key:.+}", "value", "{value:.+}").
		HandlerFunc(SetValue)

	valueRouter.
		Methods("DELETE").
		Queries("key", "{key:.+}").
		HandlerFunc(DeleteOneValue)

	valueRouter.
		Methods("DELETE").
		HandlerFunc(DeleteAllValues)

	//////////
	// root //
	//////////

	router.HandleFunc("/", RootGet).Methods("GET")

	http.ListenAndServe(":9000", enableCORS(router))

	return router
}
