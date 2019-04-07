package v1

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/modules/forms"
)

type UserController struct {
	LoginValidateController
}

// @router /all [get]
func (c *UserController) GetAll() {

	var (
		users []*models.User
		err   error
	)

	switch {
	case c.User.IsAdmin():
		_, err = models.Users().Filter("Parent", c.User.Id).All(&users)
	case c.User.IsSuperAdmin():
		_, err = models.Users().OrderBy("User").All(&users)
	default:
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err != nil {
		beego.Error(err.Error())
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
	} else {
		c.WriteJson(Json{"data": users}, 200)
	}
}

// @router /:id([0-9]+) [get]
func (c *UserController) Get() {
	userID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	user := &models.User{Id: userID}

	if err := user.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定用户"}, 404)
		return
	}

	switch {
	case c.User.IsAdmin() && user.Parent == c.User.Id:
		fallthrough
	case c.User.IsSuperAdmin():
		c.WriteJson(Json{"data": user}, 200)
	default:
		c.WriteJson(Json{"message": "无访问权限"}, 403)
	}
}

// @router /new [post]
func (c *UserController) Post() {
	if c.User.IsNormal() {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	var (
		user models.User
		form forms.UserForm
	)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.WriteJson(Json{"message": "数据格式有误"}, 400)
		return
	}

	form.Update(&user)
	user.Parent = c.User.Id

	if c.User.IsSuperAdmin() {
		user.Role = models.UserRoleAdmin
	} else {
		user.Role = models.UserRoleNormal
	}

	if err := user.Insert(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "添加成功"}, 200)
}

// @router /:id([0-9]+) [put]
func (c *UserController) Update() {
	userID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	user := models.User{Id: userID}

	if err := user.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定用户"}, 404)
		return
	}

	switch {
	case c.User.IsNormal() && c.User.Id == user.Id:
		fallthrough
	case c.User.IsAdmin() && c.User.Id == user.Parent:
		fallthrough
	case c.User.IsSuperAdmin():
		break
	default:
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	var form forms.UserForm
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.WriteJson(Json{"message": "数据格式有误"}, 400)
		return
	}

	form.Update(&user)
	if err := user.Update(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "更新成功"}, 201)
}

// @router /:id([0-9]+) [delete]
func (c *UserController) Delete() {
	userID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	user := models.User{Id: userID}

	if err := user.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定用户"}, 404)
		return
	}

	switch {
	case c.User.IsAdmin() && user.Parent == c.User.Id:
		fallthrough
	case c.User.IsSuperAdmin():
		break
	default:
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err := user.Delete(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "删除成功"}, 200)
}
