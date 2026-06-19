package handlers

import (
	"fmt"
	"net/http"
)

func Agent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Header.Get("User-Agent"))
}
