package v1

import (
	"github.com/astaxie/beego"
	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/modules/utils"
)

type UserController struct {
	BaseController
}

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

func (u *UserController) Login() {
	defer u.ServeJSON()

	if u.IsLogin {
		u.Data["json"] = Json{
			"success": false,
			"message": "请先退出当前帐号后再登录",
		}
		return
	}

	var (
		username string
		password string
	)

	username = u.GetString("username")

	if !models.HasUser(username) {
		u.Data["json"] = Json{
			"success": false,
			"message": "用户不存在",
		}
		return
	}

	user, err := models.FindUser(username)
	if err != nil {
		beego.Error(err.Error())
		u.Data["json"] = Json{
			"success": false,
			"message": "数据异常",
		}
		return
	}

	if !verifyPassword(password, user.Password) {
		u.Data["json"] = Json{
			"success": false,
			"message": "密码错误，请重新输入",
		}
		return
	}

	u.LoginUser(user)
	u.Data["json"] = Json{
		"success": true,
		"message": "登录成功",
	}
}

func (u *UserController) Logout() {
	defer u.ServeJSON()

	if u.IsLogin {
		u.Logout()
	}

	u.Data["json"] = Json{
		"success": true,
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
