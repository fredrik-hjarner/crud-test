package routes

import (
	"fmt"
	"net/http"
)

func Routes(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/":
		rootRoute(w, r)
	case "/keys":
		keysRoute(w, r)
	case "/value":
		valueRoute(w, r)
	default:
		fmt.Fprintf(w, "Unrecognized path: %s", path)
	}
}
