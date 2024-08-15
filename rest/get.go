package rest

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetRequest(url string) {
	fmt.Println("url called", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Println("Response:", string(body))
}
