package v1

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/modules/forms"
)

// ScheduleController ...
type ScheduleController struct {
	LoginValidateController
}

// @router /all [get]
func (c *ScheduleController) GetAll() {
	var (
		schedules []*models.Schedule
		page      int
		limit     int
		total     int64
		err       error
	)

	if page, err = c.GetInt("page"); err != nil {
		c.WriteJson(Json{"message": "输入有误"}, 400)
		return
	}

	if limit, err = c.GetInt("limit"); err != nil {
		c.WriteJson(Json{"message": "输入有误"}, 400)
		return
	}

	qs := models.Schedules()

	switch {
	case c.User.IsNormal():
		qs = qs.Filter("Insititution", c.User.Insititution.Id)
	case c.User.IsAdmin():
		qs = qs.OrderBy("Insititution")
	}

	if total, err = qs.Count(); err != nil {
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	if _, err = qs.Limit(limit, (page-1)*limit).All(&schedules); err != nil {
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"data": schedules, "total": total}, 200)
}

// @router /new [post]
func (c *ScheduleController) Post() {

	var (
		schedule models.Schedule
		form     forms.ScheduleForm
	)

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.WriteJson(Json{"message": "数据格式有误"}, 400)
		return
	}

	if err := form.Assign(&schedule); err != nil {
		c.WriteJson(Json{"message": "数据有误"}, 400)
		return
	}
	schedule.Insititution = c.User.Insititution

	if err := schedule.Insert(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "添加成功"}, 200)
}

// @router /:id([0-9]+) [get]
func (c *ScheduleController) Get() {

	scheduleID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	schedule := models.Schedule{Id: scheduleID}

	if err := schedule.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定时间表"}, 404)
		return
	}

	if c.User.IsNormal() && schedule.Insititution.Id != c.User.Insititution.Id {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	c.WriteJson(Json{"data": schedule}, 200)
}

// @router /:id([0-9]+) [put]
func (c *ScheduleController) Update() {
	scheduleID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	schedule := models.Schedule{Id: scheduleID}

	if err := schedule.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定时间表"}, 404)
		return
	}

	if c.User.IsNormal() && schedule.Insititution.Id != c.User.Insititution.Id {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	var form forms.ScheduleForm
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.WriteJson(Json{"message": "数据格式有误"}, 400)
		return
	}

	if err := form.Assign(&schedule); err != nil {
		c.WriteJson(Json{"message": "数据有误"}, 400)
		return
	}

	if err := schedule.Update(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "更新成功"}, 201)
}

// @router /:id([0-9]+) [delete]
func (c *ScheduleController) Delete() {
	scheduleID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	schedule := models.Schedule{Id: scheduleID}

	if err := schedule.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定时间表"}, 404)
		return
	}

	if c.User.IsNormal() && schedule.Insititution.Id != c.User.Insititution.Id {
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	if err := schedule.Delete(); err != nil {
		beego.Error(err)
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
		return
	}

	c.WriteJson(Json{"message": "删除成功"}, 200)
}
