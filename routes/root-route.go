package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// value path

func rootRoute(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method == "GET" { // TODO: replace with switch.
		rootGet(w, r)
	} else {
		// return only the value of `id`.
		fmt.Fprintf(w, "'%s' not allowed on /", method)
	}
}

/*
 * TODO: This should return all keys (i.e. folders) in alphabethical order.
 *
 */
func rootGet(w http.ResponseWriter, r *http.Request) {
	// Return documentation of API.
	t, err := template.ParseFiles("templates/root.html")
	if err != nil {
		s, _ := json.Marshal(err)
		fmt.Fprintf(w, string(s))
		return
	}
	t.Execute(w, nil)
}
