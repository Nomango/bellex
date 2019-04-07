package v1

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/nomango/bellex/server/models"
)

type UserController struct {
	LoginValidateController
}

// @router /all [get]
func (u *UserController) GetAll() {
	var users []*models.User

	switch u.User.Role {
	case models.UserRoleAdmin:
		if _, err := models.Users().Filter("Parent", u.User.Id).All(&users); err != nil {
			beego.Error(err.Error())
		}
	case models.UserRoleSuperAdmin:
		if _, err := models.Users().All(&users); err != nil {
			beego.Error(err.Error())
		}
	default:
		u.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	u.WriteJson(users, 200)
}

// @router /:id([0-9]+) [get]
func (u *UserController) Get() {
	id, _ := strconv.Atoi(u.Ctx.Input.Param(":id"))

	user := &models.User{Id: id}
	err := user.Read()

	switch {
	case err != nil:
		u.WriteJson(Json{"message": "不存在指定用户"}, 404)
		return
	case u.User.IsAdmin() && user.Parent == u.User.Id:
		break
	case u.User.IsSuperAdmin():
		break
	default:
		u.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	u.WriteJson(user, 200)
}

func (u *UserController) Post() {
	defer u.ServeJSON()
}

func (u *UserController) Delete() {
	defer u.ServeJSON()
}
