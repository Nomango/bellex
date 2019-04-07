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
		err       error
	)

	switch {
	case c.User.IsNormal():
		_, err = models.Schedules().Filter("user_id", c.User.Parent).All(&schedules)
	case c.User.IsAdmin():
		_, err = models.Schedules().Filter("User", &c.User).All(&schedules)
	case c.User.IsSuperAdmin():
		_, err = models.Schedules().OrderBy("User").All(&schedules)
	}

	if err != nil {
		beego.Error(err.Error())
		c.WriteJson(Json{"message": "系统异常，请稍后再试"}, 400)
	} else {
		c.WriteJson(Json{"data": schedules}, 200)
	}
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

	form.Update(&schedule)
	schedule.User = &c.User

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

	switch {
	case c.User.IsNormal() && schedule.User.Id == c.User.Parent:
		fallthrough
	case c.User.IsAdmin() && schedule.User.Id == c.User.Id:
		fallthrough
	case c.User.IsSuperAdmin():
		c.WriteJson(Json{"data": schedule}, 200)
	default:
		c.WriteJson(Json{"message": "无访问权限"}, 403)
	}
}

// @router /:id([0-9]+) [put]
func (c *ScheduleController) Update() {
	scheduleID, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	schedule := models.Schedule{Id: scheduleID}

	if err := schedule.Read(); err != nil {
		c.WriteJson(Json{"message": "不存在指定时间表"}, 404)
		return
	}

	switch {
	case c.User.IsNormal() && schedule.User.Id == c.User.Parent:
		fallthrough
	case c.User.IsAdmin() && schedule.User.Id == c.User.Id:
		fallthrough
	case c.User.IsSuperAdmin():
		break
	default:
		c.WriteJson(Json{"message": "无访问权限"}, 403)
		return
	}

	var form forms.ScheduleForm
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &form); err != nil {
		c.WriteJson(Json{"message": "数据格式有误"}, 400)
		return
	}

	form.Update(&schedule)
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

	switch {
	case c.User.IsNormal() && schedule.User.Id == c.User.Parent:
		fallthrough
	case c.User.IsAdmin() && schedule.User.Id == c.User.Id:
		fallthrough
	case c.User.IsSuperAdmin():
		break
	default:
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
