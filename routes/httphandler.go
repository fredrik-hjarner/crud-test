package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// CreateHTTPHandler ...
func CreateHTTPHandler() http.Handler {
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

	userHandler := NewUserHandler()

	router.HandleFunc("/users", userHandler.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetUserID).Methods("GET")
	router.HandleFunc("/users", userHandler.Create).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.Update).Methods("PUT")

	return enableCORS(router)
}
