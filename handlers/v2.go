package handlers

import (
	"fmt"
	"net/http"
)

func V2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to version 2")
}
