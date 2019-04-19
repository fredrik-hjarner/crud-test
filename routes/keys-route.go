package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fredrik-hjarner/ztorage/utils"
)

// value path

func keysRoute(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" { // TODO: replace with switch.
		keysGet(w, r)
	} else {
		// return only the value of `id`.
		fmt.Fprintf(w, "'%s' not allowed on /keys", method)
	}
}

/*
 * TODO: This should return all keys (i.e. folders) in alphabethical order.
 *
 */
func keysGet(w http.ResponseWriter, r *http.Request) {
	// return all key-value pairs.
	slice := utils.ListKeys()
	jsonString, _ := json.Marshal(slice)
	fmt.Fprintf(w, "%s", jsonString)
}
