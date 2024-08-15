package main

import (
	"fmt"
	"io"
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
		"/gateway/service1":            "http://localhost:8082/users",
		"/gateway/service1/operations": "http://localhost:8082/operations",
		"/gateway/service2":            "http://localhost:8081",
	}

	fmt.Println("path:", r.URL.Path)
	fmt.Println(r.Method)

	var response rest.CalledResponse
	route := routes[r.URL.Path]

	if r.Method == "GET" {
		response = routeGet(r, route)
	}
	if r.Method == "POST" {
		response = routePost(r, route)
	}

	w.WriteHeader(response.StatusCode)
	w.Header().Set("Content-Type", "application/json")

	w.Write(response.Body)
}

func routeGet(r *http.Request, route string) rest.CalledResponse {
	queries := r.URL.Query()
	params := make(map[string]string)
	for key, values := range queries {
		for _, value := range values {
			params[key] = value
		}
	}
	return rest.GetRequest(route, params)
}

func routePost(r *http.Request, route string) rest.CalledResponse {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return rest.CalledResponse{StatusCode: http.StatusInternalServerError}
	}
	return rest.PostRequest(route, body)
}
