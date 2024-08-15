package main

import (
	"log"
	"net/http"

	"watchdog-go.com/rest"
)

func main() {

	http.HandleFunc("/gateway/*", handleService)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func handleService(w http.ResponseWriter, r *http.Request) {

	routes := map[string]string{
		"/gateway/service1": "http://localhost:8082/users",
		"/gateway/service2": "http://localhost:8081",
	}

	route := routes[r.URL.Path]

	rest.GetRequest(route)

}
