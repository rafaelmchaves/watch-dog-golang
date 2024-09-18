package rest

import (
	"encoding/json"
	"io"
	"net/http"

	internal "watchdog-go.com/internal"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleLogin(w http.ResponseWriter, r *http.Request) {
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

		var loginSvc internal.LoginService = newLoginServiceImpl()
		loginHandler := internal.NewLoginHandler(loginSvc)

		jwt, err := loginHandler.Login(login.Email, login.Password)

		if err != nil {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		}

		w.Header().Add("token", jwt)
		w.WriteHeader(http.StatusOK)

	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func newLoginServiceImpl() *internal.LoginServiceImpl {
	return &internal.LoginServiceImpl{}
}
