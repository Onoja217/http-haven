package handlers

import (
	"fmt"
	"net/http"
)

func APIPing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func APIGreet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Stranger"
	}

	fmt.Fprintf(w, "Greetings, %s!", name)
}