package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"watchdog-go.com/internal/infrastructure/grpc"
	"watchdog-go.com/internal/service"
	"watchdog-go.com/rest"
)

func main() {

	// Initialize gRPC connection
	grpc.CreateConnectionConfig()

	// Handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Shutting down gracefully...")
		if conn := grpc.GetClientConn(); conn != nil {
			conn.Close()
		}
		os.Exit(0)
	}()

	var loginSvc service.LoginService = newLoginServiceImpl()
	loginHandler := rest.NewLoginHandler(loginSvc)

	http.HandleFunc("/gateway/*", corsMiddleware(handleService))
	http.HandleFunc("/login", corsMiddleware(loginHandler.HandleLogin))
	http.HandleFunc("/proposals", corsMiddleware(rest.HandleJob))

	// Start the HTTP server
	log.Println("Starting server on :8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}

// corsMiddleware adds CORS headers to each response
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "token")

		// Handle preflight requests (OPTIONS)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Call the next handler
		next(w, r)
	}
}

func handleService(w http.ResponseWriter, r *http.Request) {

	routes := map[string]string{
		"/gateway/service1":                         "http://localhost:8082/users",
		"/gateway/service1/operations":              "http://localhost:8082/operations",
		"/gateway/service1/operations/{id}/clients": "http://localhost:8082/operations/{id}/clients",
		"/gateway/service2":                         "http://localhost:8081",
	}

	fmt.Println("path:", r.URL.Path)
	fmt.Println(r.Method)

	var response rest.CalledResponse
	route := routes[r.URL.Path]

	if route == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

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

func newLoginServiceImpl() *service.LoginServiceImpl {
	return &service.LoginServiceImpl{}
}
