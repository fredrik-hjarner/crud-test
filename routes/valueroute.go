package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fredrik-hjarner/ztorage/diskv"
)

func GetValueByKey(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")
	log.Printf("key=%s", key)

	value := diskv.Diskv.ReadString(key)
	// if error
	if value == "" {
		w.WriteHeader(http.StatusNotFound)
		r.Body.Close()
	} else {
		fmt.Fprintf(w, "%s", value)
	}
}

func SetValue(w http.ResponseWriter, r *http.Request) {
	// body := r.Body
	query := r.URL.Query()
	key := query.Get("key")
	value := query.Get("value")
	diskv.Diskv.WriteString(key, value)
}

func DeleteOneValue(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")
	diskv.Diskv.Erase(key)
}

func DeleteAllValues(w http.ResponseWriter, r *http.Request) {
	diskv.Diskv.EraseAll()
	os.Mkdir("data-dir", 0777) // EraseAll deletes folder??
}
