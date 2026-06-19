package main

import (
	"fmt"
	"net/http"

	"http-haven/routes"
)

func main() {
	routes.RegisterRoutes()

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}