package rest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	service "watchdog-go.com/internal/service"
)

type LoginHandler struct {
	loginService service.LoginService
}

func NewLoginHandler(loginService service.LoginService) *LoginHandler {
	return &LoginHandler{loginService: loginService}
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *LoginHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}

		defer r.Body.Close()
		var login Login
		err = json.Unmarshal(body, &login)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		jwt, err := h.loginService.Login(login.Email, login.Password)

		if err != nil {
			log.Println("Error to login", err)
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		}

		w.Header().Add("token", jwt)
		w.WriteHeader(http.StatusOK)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
