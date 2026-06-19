package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func Count(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Send a POST request with text to count words")
		return
	}

	body, _ := io.ReadAll(r.Body)
	fmt.Fprint(w, len(body))
}
