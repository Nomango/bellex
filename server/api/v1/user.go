package v1

import (
	"encoding/json"
	"errors"
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
		page  int
		limit int
		err   error
	)

	defer func() {
		if err != nil {
			c.WriteJson(Json{"message": err.Error()}, 400)
		} else {
			c.WriteJson(Json{"data": users, "total": len(users)}, 200)
		}
	}()

	if page, err = c.GetInt("page"); err != nil {
		err = errors.New("请求数据有误")
		return
	}

	if limit, err = c.GetInt("limit"); err != nil {
		err = errors.New("请求数据有误")
		return
	}

	qs := models.Users().OrderBy("-CreateTime")
	if !c.User.IsAdmin() {
		qs = qs.Filter("Insititution", c.User.Insititution)
	}
	_, err = qs.Exclude("Id", c.User.Id).Limit(limit, (page-1)*limit).All(&users)
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
	case c.User.IsNormal() && user.Insititution.Id == c.User.Insititution.Id:
		fallthrough
	case c.User.IsAdmin():
		c.WriteJson(Json{"data": user}, 200)
	default:
		c.WriteJson(Json{"message": "无访问权限"}, 403)
	}
}

// @router /new [post]
func (c *UserController) Post() {

	var (
		user models.User
		form forms.UserRegisterForm
	)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.WriteJson(Json{"message": "数据格式有误"}, 400)
		return
	}

	if err := form.Assign(&user); err != nil {
		c.WriteJson(Json{"message": "数据有误"}, 400)
		return
	}
	user.Role = models.UserRoleNormal

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

	if c.User.IsNormal() && c.User.Id != userID {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err := user.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定用户"}, 404)
		return
	}

	var form forms.UserForm
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.WriteJson(Json{"message": "数据格式有误"}, 400)
		return
	}

	if err := form.Assign(&user); err != nil {
		c.WriteJson(Json{"message": "数据有误"}, 400)
		return
	}

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

	if c.User.IsNormal() {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err := user.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定用户"}, 404)
		return
	}

	if err := user.Delete(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "删除成功"}, 200)
}
