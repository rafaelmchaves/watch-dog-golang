package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// Define the routes and their respective backend services
	routes := map[string]string{
		"/service1": "http://localhost:8082/users",
		"/service2": "http://localhost:8081",
	}

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Create a reverse proxy for each route
	for route, target := range routes {
		targetURL, err := url.Parse(target)
		if err != nil {
			log.Fatalf("Failed to parse URL: %v", err)
		}

		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		// Create a handler for the route
		mux.Handle(route, proxy)
	}

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
