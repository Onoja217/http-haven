package routes

import (
	"net/http"

	"http-haven/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/ping", handlers.Ping)
	http.HandleFunc("/hello", handlers.Hello)
	http.HandleFunc("/count", handlers.Count)
	http.HandleFunc("/calculate", handlers.Calculate)
	http.HandleFunc("/agent", handlers.Agent)
	http.HandleFunc("/dashboard", handlers.Dashboard)
	http.HandleFunc("/legacy", handlers.Legacy)
	http.HandleFunc("/v2", handlers.V2)
}