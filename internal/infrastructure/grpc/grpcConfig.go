package grpc

import (
	"log"
	"sync"

	"google.golang.org/grpc"
)

var (
	connection *grpc.ClientConn
	once       sync.Once
)

// CreateConnectionConfig initializes the gRPC connection
func CreateConnectionConfig() {
	once.Do(func() { // Ensure the connection is created only once
		var err error
		connection, err = grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect to gRPC server: %v", err)
		}
		log.Println("gRPC server connected on :50051")
	})
}

// GetClientConn returns the gRPC connection
func GetClientConn() *grpc.ClientConn {
	if connection == nil {
		log.Fatal("gRPC connection is not initialized. Call CreateConnectionConfig first.")
	}
	return connection
}
