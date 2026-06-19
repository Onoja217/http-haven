package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func Calculate(w http.ResponseWriter, r *http.Request) {
	op := r.URL.Query().Get("op")
	a, _ := strconv.Atoi(r.URL.Query().Get("a"))
	b, _ := strconv.Atoi(r.URL.Query().Get("b"))

	var result int

	switch op {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "multiply":
		result = a * b
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%d", result)
}
