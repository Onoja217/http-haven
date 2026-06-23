package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Stretch Goal: Check Content-Type
	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
		http.Error(
			w,
			"Content-Type must be application/x-www-form-urlencoded",
			http.StatusUnsupportedMediaType,
		)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "failed to parse form", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	language := r.FormValue("language")

	if username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}

	if language == "" {
		http.Error(w, "language is required", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(
		w,
		"Hello %s, you are coding in %s!",
		username,
		language,
	)
}