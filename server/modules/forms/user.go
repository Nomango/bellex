package forms

import "github.com/nomango/bellex/server/models"

type UserLoginForm struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterForm struct {
	UserName       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	InsititutionID int    `json:"insititution_id"`
}

func (u *UserRegisterForm) Assign(user *models.User) error {
	insititution := &models.Insititution{Id: u.InsititutionID}
	if err := insititution.Read(); err != nil {
		return err
	}

	user.UserName = u.UserName
	user.SetNewPassword(u.Password)
	user.Email = u.Email
	user.Insititution = insititution
	return nil
}

type UserForm struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
}

func (u *UserForm) Assign(user *models.User) error {
	user.UserName = u.UserName
	user.Email = u.Email
	return nil
}
