package v1

import (
	"encoding/json"
	"regexp"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/modules/forms"
)

// MechineController ...
type MechineController struct {
	LoginValidateController
}

// @router /all [get]
func (c *MechineController) GetAll() {
	var (
		mechines []*models.Mechine
		page     int
		limit    int
		err      error
	)

	defer func() {
		if err != nil {
			c.WriteJson(Json{"message": "输入有误"}, 400)
		} else {
			for _, m := range mechines {
				m.UpdateStatus()
			}
			c.WriteJson(Json{"data": mechines, "total": len(mechines)}, 200)
		}
	}()

	if page, err = c.GetInt("page"); err != nil {
		return
	}

	if limit, err = c.GetInt("limit"); err != nil {
		return
	}

	switch {
	case c.User.IsNormal():
		_, err = models.Mechines().Filter("Institution", c.User.Institution.Id).Limit(limit, (page-1)*limit).RelatedSel().All(&mechines)
	case c.User.IsAdmin():
		_, err = models.Mechines().OrderBy("Institution").Limit(limit, (page-1)*limit).RelatedSel().All(&mechines)
	}
}

// @router /new [post]
func (c *MechineController) Post() {

	var (
		mechine models.Mechine
		form    forms.MechineForm
	)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.WriteJson(Json{"message": "数据格式有误"}, 400)
		return
	}

	if err := form.Assign(&mechine); err != nil {
		c.WriteJson(Json{"message": "数据有误"}, 400)
		return
	}
	mechine.Institution = c.User.Institution
	mechine.SetNewSecret()

	if err := mechine.Insert(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "添加成功"}, 200)
}

// @router /:id([0-9]+) [get]
func (c *MechineController) Get() {

	mechineID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	mechine := models.Mechine{Id: mechineID}

	if err := mechine.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定主控机"}, 404)
		return
	}

	if c.User.IsNormal() && mechine.Institution.Id != c.User.Institution.Id {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	mechine.UpdateStatus()
	c.WriteJson(Json{"data": mechine}, 200)
}

// @router /:id([0-9]+) [put]
func (c *MechineController) Update() {
	mechineID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	mechine := models.Mechine{Id: mechineID}

	if err := mechine.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定主控机"}, 404)
		return
	}

	if c.User.IsNormal() && mechine.Institution.Id != c.User.Institution.Id {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	var form forms.MechineForm
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.WriteJson(Json{"message": "数据格式有误"}, 400)
		return
	}

	if err := form.Assign(&mechine); err != nil {
		c.WriteJson(Json{"message": "数据有误"}, 400)
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
	mechine := &models.Mechine{Id: mechineID}

	if err := mechine.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定主控机"}, 404)
		return
	}

	if c.User.IsNormal() && mechine.Institution.Id != c.User.Institution.Id {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	mechine.UpdateStatus()
	if mechine.Accept {
		models.DeleteConnection(mechine)
	}

	if err := mechine.Delete(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "删除成功"}, 200)
}

// @router /:id([0-9]+)/start/current [post]
func (c *MechineController) StartCurrent() {
	mechineID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	mechine := models.Mechine{Id: mechineID}

	if err := mechine.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定主控机"}, 404)
		return
	}

	if c.User.IsNormal() && mechine.Institution.Id != c.User.Institution.Id {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err := mechine.SendData(`bell:current;`); err != nil {
		c.WriteJson(Json{"message": "主控机未连接"}, 403)
		return
	}

	c.WriteJson(Json{"message": "发送成功"}, 200)
}

// @router /:id([0-9]+)/start [post]
func (c *MechineController) Start() {
	mechineID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	mechine := models.Mechine{Id: mechineID}

	if err := mechine.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定主控机"}, 404)
		return
	}

	if c.User.IsNormal() && mechine.Institution.Id != c.User.Institution.Id {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	time := c.GetString("time")
	matched, _ := regexp.MatchString(`^\d{2}:\d{2}$`, time)
	if !matched {
		c.WriteJson(Json{"message": "数据格式有误"}, 403)
		return
	}

	if err := mechine.SendData(`bell:` + time[:2] + time[3:] + ";"); err != nil {
		c.WriteJson(Json{"message": "主控机未连接"}, 403)
		return
	}

	c.WriteJson(Json{"message": "发送成功"}, 200)
}

// @router /:id([0-9]+)/close [post]
func (c *MechineController) Close() {
	mechineID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	mechine := models.Mechine{Id: mechineID}

	if err := mechine.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定主控机"}, 404)
		return
	}

	if c.User.IsNormal() && mechine.Institution.Id != c.User.Institution.Id {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	models.DeleteConnection(&mechine)
	c.WriteJson(Json{"message": "断开成功"}, 200)
}
