package internal

type LoginHandler struct {
	loginService LoginService
}

func NewLoginHandler(loginService LoginService) *LoginHandler {
	return &LoginHandler{loginService: loginService}
}

type LoginService interface {
	Login(login string, password string) (string, error)
}

type LoginServiceImpl struct{}

func (loginService *LoginServiceImpl) Login(login string, password string) (string, error) {
	//TODO call login function
	//TODO generate token

	return "", nil
}
