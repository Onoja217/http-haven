package handlers

import (
	"fmt"
	"net/http"
)

// r.Header.Get() returns an empty string if the header was not sent.

func Headers(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("X-Custom-Token")

	if token == "" {
		http.Error(w, "X-Custom-Token header is missing", http.StatusBadRequest)
		return
	}

	contentType := r.Header.Get("Content-Type")

	if contentType == "" {
		contentType = "Content-Type not provided"
	} else {
		contentType = "Content-Type: " + contentType
	}

	fmt.Fprintf(w,
		"Token received: %s\n%s",
		token,
		contentType,
	)
}
