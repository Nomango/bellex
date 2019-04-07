package v1

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/nomango/bellex/server/models"
)

// MechineController ...
type MechineController struct {
	LoginValidateController
}

// @router /all [get]
func (c *MechineController) GetAll() {
	var (
		mechines []*models.Mechine
		err      error
	)

	switch {
	case c.User.IsNormal():
		_, err = models.Mechines().Filter("user_id", c.User.Parent).All(&mechines)
	case c.User.IsAdmin():
		_, err = models.Mechines().Filter("User", &c.User).All(&mechines)
	case c.User.IsSuperAdmin():
		_, err = models.Mechines().OrderBy("User").All(&mechines)
	}

	if err != nil {
		beego.Error(err.Error())
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
	} else {
		c.WriteJson(Json{"data": mechines}, 200)
	}
}

// @router /new [post]
func (c *MechineController) Post() {

	var mechine models.Mechine
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &mechine); err != nil {
		c.WriteJson(Json{"message": "数据格式有误"}, 400)
		return
	}

	if err := mechine.Insert(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "请求成功"}, 200)
}

// @router /:id([0-9]+) [get]
func (c *MechineController) Get() {

	mechineID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	mechine := models.Mechine{Id: mechineID}

	if err := mechine.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定主控机"}, 404)
		return
	}

	switch {
	case c.User.IsNormal() && mechine.User.Id == c.User.Parent:
		fallthrough
	case c.User.IsAdmin() && mechine.User.Id == c.User.Id:
		fallthrough
	case c.User.IsSuperAdmin():
		c.WriteJson(Json{"data": mechine}, 200)
	default:
		c.WriteJson(Json{"message": "无访问权限"}, 403)
	}
}

// @router /:id([0-9]+) [put]
func (c *MechineController) Put() {
	mechineID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	mechine := models.Mechine{Id: mechineID}

	if err := mechine.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定主控机"}, 404)
		return
	}

	switch {
	case c.User.IsNormal() && mechine.User.Id == c.User.Parent:
		fallthrough
	case c.User.IsAdmin() && mechine.User.Id == c.User.Id:
		fallthrough
	case c.User.IsSuperAdmin():
		break
	default:
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &mechine); err != nil {
		c.WriteJson(Json{"message": "数据格式有误"}, 400)
		return
	}

	if err := mechine.Update(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "更新成功"}, 201)
}

// @router /:id([0-9]+) [delete]
func (c *MechineController) Delete() {
	mechineID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	mechine := models.Mechine{Id: mechineID}

	if err := mechine.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定主控机"}, 404)
		return
	}

	switch {
	case c.User.IsNormal() && mechine.User.Id == c.User.Parent:
		fallthrough
	case c.User.IsAdmin() && mechine.User.Id == c.User.Id:
		fallthrough
	case c.User.IsSuperAdmin():
		break
	default:
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err := mechine.Delete(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "删除成功"}, 200)
}
