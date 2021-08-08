package auth_service

import "github.com/cuongtop4598/QuizWithGo/QuizApp/models"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}
