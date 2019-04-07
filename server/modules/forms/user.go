package forms

import "github.com/nomango/bellex/server/models"

type UserLoginForm struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserForm struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *UserForm) Update(user *models.User) {
	user.UserName = u.UserName
	user.SetNewPassword(u.Password)
	user.Email = u.Email
}
