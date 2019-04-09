package forms

import (
	"errors"

	"github.com/nomango/bellex/server/models"
)

type UserLoginForm struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Password string `json:"password"`
}

type UserRegisterForm struct {
	UserName      string `json:"username"`
	NickName      string `json:"nickname"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	InstitutionID int    `json:"institution_id"`
}

func (u *UserRegisterForm) Assign(user *models.User) error {
	institution := &models.Institution{Id: u.InstitutionID}
	if err := institution.Read(); err != nil {
		return errors.New("指定机构不存在")
	}

	if models.HasUser(user.UserName) {
		return errors.New("用户已被占用")
	}

	if models.HasUser(user.Email) {
		return errors.New("邮箱已被占用")
	}

	user.UserName = u.UserName
	user.NickName = u.NickName
	user.Password = u.Password
	user.Email = u.Email
	user.Institution = institution
	return nil
}

type UserForm struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
}

func (u *UserForm) Assign(user *models.User) error {
	if user.UserName != u.UserName && models.HasUser(u.UserName) {
		return errors.New("用户已被占用")
	}

	if user.Email != u.Email && models.HasUser(u.Email) {
		return errors.New("邮箱已被占用")
	}

	user.UserName = u.UserName
	user.NickName = u.NickName
	user.Email = u.Email
	return nil
}
