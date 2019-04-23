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

	if prefix != "" {
		log.Printf("prefix=%s", prefix)
		slice := utils.ListKeysWithPrefix(prefix)
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
	} else {
		// return all key-value pairs.
		slice := utils.ListKeys()
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
}
