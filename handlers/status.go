package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func Status(w http.ResponseWriter, r *http.Request) {
	codeStr := r.URL.Query().Get("code")

	if codeStr == "" {
		http.Error(w, "code parameter is required", http.StatusBadRequest)
		return
	}

	code, err := strconv.Atoi(codeStr)
	if err != nil {
		http.Error(w, "code must be a valid integer", http.StatusBadRequest)
		return
	}

	if code < 100 || code > 599 {
		http.Error(w, "code must be a valid HTTP status code (100–599)", http.StatusBadRequest)
		return
	}

	// IMPORTANT: WriteHeader comes FIRST
	w.WriteHeader(code)

	fmt.Fprintf(w, "Responding with status %d %s", code, http.StatusText(code))
}