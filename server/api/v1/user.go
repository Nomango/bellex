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

	var (
		users []*models.User
		err   error
	)

	switch {
	case u.User.IsAdmin():
		_, err = models.Users().Filter("Parent", u.User.Id).All(&users)
	case u.User.IsSuperAdmin():
		_, err = models.Users().OrderBy("User").All(&users)
	default:
		u.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err != nil {
		beego.Error(err.Error())
		u.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
	} else {
		u.WriteJson(Json{"data": users}, 200)
	}
}

// @router /:id([0-9]+) [get]
func (u *UserController) Get() {
	userID, _ := strconv.Atoi(u.Ctx.Input.Param(":id"))
	user := &models.User{Id: userID}

	if err := user.Read(); err != nil {
		u.WriteJson(Json{"message": "不存在指定用户"}, 404)
		return
	}

	switch {
	case u.User.IsAdmin() && user.Parent == u.User.Id:
		fallthrough
	case u.User.IsSuperAdmin():
		u.WriteJson(Json{"data": user}, 200)
	default:
		u.WriteJson(Json{"message": "无访问权限"}, 403)
	}
}

func (u *UserController) Post() {
	defer u.ServeJSON()
}

func (u *UserController) Delete() {
	defer u.ServeJSON()
}
