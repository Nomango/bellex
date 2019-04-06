package v1

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/modules/utils"
)

type UserController struct {
	BaseController
}

// @router /all [get]
func (u *UserController) GetAll() {
	defer u.ServeJSON()

	if u.IsLogin {
		var users []*models.User
		if u.User.IsAdmin() {
			if _, err := models.Users().Filter("Parent", u.User.Id).All(&users); err != nil {
				beego.Error(err.Error())
			}
		} else if u.User.IsSuperAdmin() {
			if _, err := models.Users().All(&users); err != nil {
				beego.Error(err.Error())
			}
		}
		u.Data["json"] = Json{
			"success": true,
			"data":    users,
		}
	}

	u.Data["json"] = Json{
		"success": false,
		"message": "权限不足",
	}
}

func (u *UserController) Get() {
	defer u.ServeJSON()

	id, err := u.GetInt("user_id")

	if err != nil {
		u.Data["json"] = Json{
			"success": false,
			"message": "User ID 必须为整数",
		}
		return
	}

	user := &models.User{Id: id}

	if err := user.Read(); err != nil {
		u.Data["json"] = Json{
			"success": false,
			"message": "用户不存在",
		}
		return
	}

	if u.IsLogin {
		if u.User.IsAdmin() && user.Parent == u.User.Id {
			u.Data["json"] = Json{
				"success": true,
				"data":    user,
			}
			return
		}
		if u.User.IsSuperAdmin() {
			u.Data["json"] = Json{
				"success": true,
				"data":    user,
			}
			return
		}
	}

	u.Data["json"] = Json{
		"success": false,
		"message": "权限不足",
	}
}

func (u *UserController) Post() {
	defer u.ServeJSON()
}

func (u *UserController) Delete() {
	defer u.ServeJSON()
}

// @router /login [post]
func (u *UserController) Login() {
	defer u.ServeJSON()

	result := Json{
		"success": false,
		"message": "",
	}
	u.Data["json"] = &result

	if u.IsLogin {
		result["message"] = "请先退出当前帐号后再登录"
		return
	}

	var params struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.Unmarshal(u.Ctx.Input.RequestBody, &params); err != nil {
		result["message"] = "数据有误"
		return
	}

	if !models.HasUser(params.UserName) {
		result["message"] = "用户不存在"
		return
	}

	user, err := models.FindUser(params.UserName)
	if err != nil {
		beego.Error(err.Error())
		result["message"] = "数据异常"
		return
	}

	if !verifyPassword(params.Password, user.Password) {
		result["message"] = "密码错误，请重新输入"
		return
	}

	u.LoginUser(user)
	result["success"] = true
	result["message"] = "登录成功"
}

// @router /logout [post]
func (u *UserController) Logout() {
	defer u.ServeJSON()

	if u.IsLogin {
		u.Logout()
	}

	u.Data["json"] = Json{
		"success": true,
	}
}

// @router /status [get]
func (u *UserController) Status() {
	defer u.ServeJSON()

	if u.IsLogin {
		u.Data["json"] = Json{
			"is_login": true,
			"user":     &u.User,
		}
	}

	u.Data["json"] = Json{
		"is_login": false,
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
