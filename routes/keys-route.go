package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fredrik-hjarner/ztorage/utils"
)

/*
 * TODO: This should return all keys (i.e. folders) in alphabethical order.
 *
 */
func KeysGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	prefix := query.Get("prefix")
	var slice []string

	log.Printf("prefix=%s", prefix)

	if prefix != "" {
		slice = utils.ListKeysWithPrefix(prefix)
	} else {
		// return all key-value pairs.
		slice = utils.ListKeys()
	}

	log.Println(slice)
	jsonString, err := json.Marshal(slice)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}

	// for some damn reason [] is serialized as "null" and not "[]".
	if len(slice) == 0 {
		fmt.Fprintf(w, "%s", "[]")
	} else {
		fmt.Fprintf(w, "%s", jsonString)
	}
}
