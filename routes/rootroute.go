package routes

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func RootGet(w http.ResponseWriter, r *http.Request) {
	// Return documentation of API.
	t, err := template.ParseFiles("templates/root.html")
	if err != nil {
		s, _ := json.Marshal(err)
		fmt.Fprintf(w, string(s))
		return
	}
	t.Execute(w, nil)
}
