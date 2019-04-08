package v1

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/modules/forms"
)

// InstitutionController ...
type InstitutionController struct {
	LoginValidateController
}

// @router /all [get]
func (c *InstitutionController) GetAll() {
	var (
		institutions []*models.Institution
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
			c.WriteJson(Json{"data": institutions, "total": len(institutions)}, 200)
		}
	}()

	if page, err = c.GetInt("page"); err != nil {
		return
	}

	if limit, err = c.GetInt("limit"); err != nil {
		return
	}

	if page == 0 && limit == 0 {
		_, err = models.Institutions().All(&institutions)
	} else {
		_, err = models.Institutions().Limit(limit, (page-1)*limit).All(&institutions)
	}
}

// @router /new [post]
func (c *InstitutionController) Post() {

	var (
		institution models.Institution
		form         forms.InstitutionForm
	)

	if !c.User.IsAdmin() {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.WriteJson(Json{"message": "数据格式有误"}, 400)
		return
	}

	if err := form.Assign(&institution); err != nil {
		c.WriteJson(Json{"message": "数据有误"}, 400)
		return
	}

	if err := institution.Insert(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "添加成功"}, 200)
}

// @router /:id([0-9]+) [get]
func (c *InstitutionController) Get() {

	institutionID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	institution := models.Institution{Id: institutionID}

	if !c.User.IsAdmin() {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err := institution.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定主控机"}, 404)
		return
	}

	c.WriteJson(Json{"data": institution}, 200)
}

// @router /:id([0-9]+) [put]
func (c *InstitutionController) Update() {
	institutionID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	institution := models.Institution{Id: institutionID}

	if !c.User.IsAdmin() {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err := institution.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定主控机"}, 404)
		return
	}

	var form forms.InstitutionForm
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.WriteJson(Json{"message": "数据格式有误"}, 400)
		return
	}

	if err := form.Assign(&institution); err != nil {
		c.WriteJson(Json{"message": "数据有误"}, 400)
		return
	}

	if err := institution.Update(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "更新成功"}, 201)
}

// @router /:id([0-9]+) [delete]
func (c *InstitutionController) Delete() {
	institutionID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	institution := models.Institution{Id: institutionID}

	if !c.User.IsAdmin() {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err := institution.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定主控机"}, 404)
		return
	}

	if err := institution.Delete(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "删除成功"}, 200)
}
