package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fredrik-hjarner/crud-test/diskv"
)

func GetValueByKey(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	namespace := query.Get("namespace")
	key := query.Get("key")

	value := diskv.Diskv.ReadString(fmt.Sprintf("%v/%v", namespace, key))
	// if error
	if value == "" {
		w.WriteHeader(http.StatusNotFound)
		if r.Body != nil {
			/*
			 *httptest.ResponseRecorder has no Body for some reason...
			 so tests fail without this if.
			*/
			r.Body.Close()
		}
	} else {
		fmt.Fprintf(w, "%s", value)
	}
}

func SetValue(w http.ResponseWriter, r *http.Request) {
	// body := r.Body
	query := r.URL.Query()
	namespace := query.Get("namespace")
	key := query.Get("key")
	value := query.Get("value")
	diskv.Diskv.WriteString(fmt.Sprintf("%v/%v", namespace, key), value)
}

func DeleteOneValue(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	namespace := query.Get("namespace")
	key := query.Get("key")
	diskv.Diskv.Erase(fmt.Sprintf("%v/%v", namespace, key))
}

func DeleteAllValues(w http.ResponseWriter, r *http.Request) {
	diskv.Diskv.EraseAll()
	os.Mkdir("data-dir", 0777) // EraseAll deletes folder??
}
