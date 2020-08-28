package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fredrik-hjarner/crud-test/requestmodels"
	"github.com/fredrik-hjarner/crud-test/storage"
)

// CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	createUserRequest := &requestmodels.CreateUserRequest{}

	err := json.NewDecoder(r.Body).Decode(createUserRequest)

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

	user := createUserRequest.ToUser()

	storage.AddUser(user)

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
