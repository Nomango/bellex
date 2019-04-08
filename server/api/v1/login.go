package v1

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/modules/forms"
	"github.com/nomango/bellex/server/modules/utils"
)

// UserLoginController ...
type UserLoginController struct {
	APIController
}

// @router /login [post]
func (u *UserLoginController) Login() {

	if u.IsLogin {
		u.WriteJson(Json{"message": "请先退出当前帐号后再登录"}, 400)
		return
	}

	var form forms.UserLoginForm
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &form); err != nil {
		u.WriteJson(Json{"message": "账号数据有误"}, 400)
		return
	}

	if !models.HasUser(form.UserName) {
		u.WriteJson(Json{"message": "用户不存在"}, 400)
		return
	}

	user, err := models.FindUser(form.UserName)
	if err != nil {
		beego.Error(err.Error())
		u.WriteJson(Json{"message": "数据异常"}, 400)
		return
	}

	if !verifyPassword(form.Password, user.Password) {
		u.WriteJson(Json{"message": "密码错误，请重新输入"}, 400)
		return
	}

	u.LoginUser(user)
	u.WriteJson(Json{"message": "登录成功", "redirect_url": "/"}, 200)
}

// @router /logout [post]
func (u *UserLoginController) Logout() {
	if u.IsLogin {
		u.LogoutUser()
	}

	u.WriteJson(Json{"message": "退出成功"}, 200)
}

// @router /status [get]
func (u *UserLoginController) Status() {
	if u.IsLogin {
		u.WriteJson(Json{"user": &u.User}, 200)
	} else {
		u.WriteJson(Json{"message": "用户未登录"}, 400)
	}
}

// VerifyPassword compare raw password and encoded password
func verifyPassword(rawPwd, encodedPwd string) bool {

	// split
	var salt, encoded string
	if len(encodedPwd) > 11 {
		salt = encodedPwd[:10]
		encoded = encodedPwd[11:]
	}

	return utils.EncodePassword(rawPwd, salt) == encoded
}
