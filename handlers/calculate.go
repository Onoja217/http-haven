package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func Calculate(w http.ResponseWriter, r *http.Request) {
	op := r.URL.Query().Get("op")

	a, errA := strconv.Atoi(r.URL.Query().Get("a"))
	b, errB := strconv.Atoi(r.URL.Query().Get("b"))
	
	fmt.Println("a =", r.URL.Query().Get("a"))
	fmt.Println("b =", r.URL.Query().Get("b"))
	fmt.Println("errA =", errA)
	fmt.Println("errB =", errB)

	if errA != nil || errB != nil {
		http.Error(w, "invalid parameters", http.StatusBadRequest)
		return
	}

	var result int

	switch op {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "multiply":
		result = a * b
	default:
		http.Error(w, "unknown operation", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Result: %d", result)
}
