package rest

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type CalledResponse struct {
	StatusCode int
	Body       []byte
}

func GetRequest(baseURL string, params map[string]string) CalledResponse {
	fmt.Println("url called", baseURL)

	u, err := url.Parse(baseURL)
	if err != nil {
		log.Fatalf("Failed to parse URL: %v", err)
	}

	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
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
		Body:       body,
	}
}
