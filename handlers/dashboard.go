package handlers

import (
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-Key") != "secret123" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Welcome"))
}
