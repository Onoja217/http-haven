package handlers

import (
	"net/http"
)

func Legacy(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
}

func V2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to version 2"))
}
