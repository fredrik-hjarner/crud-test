package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fredrik-hjarner/crud-test/requestmodels"
	"github.com/fredrik-hjarner/crud-test/storage"
	"github.com/gorilla/mux"
)

// UpdateUser ...
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	updateUserRequest := &requestmodels.UpdateUserRequest{}

	err := json.NewDecoder(r.Body).Decode(updateUserRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Could not parse request params.")
		log.Println(err.Error())
		if r.Body != nil {
			/*
			 *httptest.ResponseRecorder has no Body for some reason...
			 so tests fail without this if.
			*/
			r.Body.Close()
		}
		return
	}

	// Check if user exists, else not found.
	vars := mux.Vars(r)
	id := vars["id"]

	user, getUserByIDError := storage.GetUserByID(id)

	if getUserByIDError != nil {
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

	// For each field replace in `user`.
	// TODO: this is very suboptimal

	{
		userJSON, err := json.Marshal(user)
		if err != nil {
			log.Println("Marshal error.")
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
		fmt.Fprintf(w, "%s", userJSON)
	}

}
