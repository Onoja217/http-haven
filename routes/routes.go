package routes

import (
	"net/http"

	"http-haven/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/render", handlers.Render)
	http.HandleFunc("/status", handlers.Status)
	http.HandleFunc("/echo", handlers.Echo)
	http.HandleFunc("/form", handlers.FormHandler)
	http.HandleFunc("/headers", handlers.Headers)
	http.HandleFunc("/method-inspector", handlers.MethodInspector)

	http.HandleFunc("/ping", handlers.Ping)
	http.HandleFunc("/hello", handlers.Hello)
	http.HandleFunc("/count", handlers.Count)
	http.HandleFunc("/calculate", handlers.Calculate)

	http.HandleFunc("/agent", handlers.Agent)
	http.HandleFunc("/dashboard", handlers.Dashboard)
	http.HandleFunc("/legacy", handlers.Legacy)
	http.HandleFunc("/v2", handlers.V2)

	// NEW API V1 (simple version)
	http.HandleFunc("/api/v1/ping", handlers.APIPing)
	http.HandleFunc("/api/v1/greet", handlers.APIGreet)
}
