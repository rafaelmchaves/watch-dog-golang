package rest

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type CalledResponse struct {
	StatusCode int
}

func GetRequest(url string) CalledResponse {
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

	return CalledResponse{
		StatusCode: resp.StatusCode,
	}
}
