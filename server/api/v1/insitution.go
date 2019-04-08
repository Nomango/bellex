package v1

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/modules/forms"
)

// InsititutionController ...
type InsititutionController struct {
	LoginValidateController
}

// @router /all [get]
func (c *InsititutionController) GetAll() {
	var (
		insititutions []*models.Insititution
		page          int
		limit         int
		err           error
	)

	if !c.User.IsAdmin() {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	defer func() {
		if err != nil {
			c.WriteJson(Json{"message": "输入有误"}, 400)
		} else {
			c.WriteJson(Json{"data": insititutions, "total": len(insititutions)}, 200)
		}
	}()

	if page, err = c.GetInt("page"); err != nil {
		return
	}

	if limit, err = c.GetInt("limit"); err != nil {
		return
	}

	_, err = models.Insititutions().Limit(limit, (page-1)*limit).All(&insititutions)
}

// @router /new [post]
func (c *InsititutionController) Post() {

	var (
		insititution models.Insititution
		form         forms.InsititutionForm
	)

	if !c.User.IsAdmin() {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.WriteJson(Json{"message": "数据格式有误"}, 400)
		return
	}

	if err := form.Assign(&insititution); err != nil {
		c.WriteJson(Json{"message": "数据有误"}, 400)
		return
	}

	if err := insititution.Insert(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "添加成功"}, 200)
}

// @router /:id([0-9]+) [get]
func (c *InsititutionController) Get() {

	insititutionID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	insititution := models.Insititution{Id: insititutionID}

	if !c.User.IsAdmin() {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err := insititution.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定主控机"}, 404)
		return
	}

	c.WriteJson(Json{"data": insititution}, 200)
}

// @router /:id([0-9]+) [put]
func (c *InsititutionController) Update() {
	insititutionID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	insititution := models.Insititution{Id: insititutionID}

	if !c.User.IsAdmin() {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err := insititution.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定主控机"}, 404)
		return
	}

	var form forms.InsititutionForm
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.WriteJson(Json{"message": "数据格式有误"}, 400)
		return
	}

	if err := form.Assign(&insititution); err != nil {
		c.WriteJson(Json{"message": "数据有误"}, 400)
		return
	}

	if err := insititution.Update(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "更新成功"}, 201)
}

// @router /:id([0-9]+) [delete]
func (c *InsititutionController) Delete() {
	insititutionID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	insititution := models.Insititution{Id: insititutionID}

	if !c.User.IsAdmin() {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err := insititution.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定主控机"}, 404)
		return
	}

	if err := insititution.Delete(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "删除成功"}, 200)
}
