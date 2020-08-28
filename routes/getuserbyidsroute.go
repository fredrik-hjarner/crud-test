package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fredrik-hjarner/crud-test/storage"
	"github.com/gorilla/mux"
)

// GetUserByID ...
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := storage.GetUserByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		if r.Body != nil {
			/*
			 *httptest.ResponseRecorder has no Body for some reason...
			 so tests fail without this if.
			*/
			r.Body.Close()
		}
		return
	}

	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
		w.WriteHeader(http.StatusInternalServerError)
		if r.Body != nil {
			/*
			 *httptest.ResponseRecorder has no Body for some reason...
			 so tests fail without this if.
			*/
			r.Body.Close()
		}
		return
	}

	userJSON, err := json.Marshal(user)

	fmt.Fprintf(w, "%s", userJSON)
}
