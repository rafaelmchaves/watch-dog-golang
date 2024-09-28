package infrastructure

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type User struct {
	Id       string
	Name     string
	Email    string
	UserType string
	Phone    string
	City     string
	State    string
	Fields   []string
}

type loginRequest struct {
	Username string
	Password string
}

func CheckLogin(username string, password string) *User {

	login := loginRequest{
		Username: username,
		Password: password,
	}

	jsonRequest, err := json.Marshal(login)
	if err != nil {
		log.Println("it was not possible to convert login object to json request byte ", err.Error())
		return nil
	}

	const baseURL = "localhost:8080"
	req, err := http.NewRequest("POST", baseURL+"/login", bytes.NewBuffer(jsonRequest))
	if err != nil {
		return nil
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make POST request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to make read the response body: %v", err)
		return nil
	}
	var user User

	err = json.Unmarshal(responseBody, user)
	if err != nil {
		log.Fatalf("Failed to unmarshal the response body in the user struct: %v", err)
		return nil
	}

	return &user
}
