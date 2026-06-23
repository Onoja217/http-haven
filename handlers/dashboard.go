package handlers

import (
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-API-Key")

	if key != "secret123" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Welcome to dashboard"))
}