package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fredrik-hjarner/crud-test/storage"
)

// GetAllUsers ...
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	usersJSON, err := json.Marshal(storage.Users)
	if err != nil {
		log.Println("Cannot encode to JSON ", err)
	}

	// for some damn reason [] is serialized as "null" and not "[]".
	if len(storage.Users) == 0 {
		fmt.Fprintf(w, "%s", "[]")
	} else {
		fmt.Fprintf(w, "%s", usersJSON)
	}
}
