package v1

import (
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/nomango/bellex/server/models"
)

// ScheduleController ...
type ScheduleController struct {
	LoginValidateController
}

// Post ...
func (b *ScheduleController) Post() {
	result := Json{
		"ok": false,
	}

	defer func() {
		b.Data["json"] = result
		b.ServeJSON()
	}()

	var ob models.Schedule
	if err := json.Unmarshal(b.Ctx.Input.RequestBody, &ob); err != nil {
		result["message"] = err.Error()
		return
	}

	if err := ob.Insert(); err != nil {
		beego.Error(err)
		result["message"] = err.Error()
		return
	}

	result["ok"] = true
}

// Get ...
func (b *ScheduleController) Get() {
	defer b.ServeJSON()

	scheduleID, err := strconv.Atoi(b.Ctx.Input.Param(":schedule_id"))
	if err != nil {
		b.Data["json"] = Json{"ok": false, "message": err.Error()}
		return
	}

	schedule := models.Schedule{Id: scheduleID}
	if err := schedule.Read(); err != nil {
		b.Data["json"] = Json{"ok": false, "message": err.Error()}
		return
	}
	b.Data["json"] = Json{"ok": true, "data": schedule}
}

// GetAll ...
func (b *ScheduleController) GetAll() {
	var schedules []*models.Schedule
	if _, err := models.Schedules().All(&schedules); err != nil {
		b.Data["json"] = Json{"ok": false, "message": err.Error()}
	} else {
		b.Data["json"] = Json{"ok": true, "data": schedules}
	}
	b.ServeJSON()
}

// Put ...
func (b *ScheduleController) Put() {
	result := Json{
		"ok": false,
	}

	defer func() {
		b.Data["json"] = result
		b.ServeJSON()
	}()

	var ob models.Schedule
	if err := json.Unmarshal(b.Ctx.Input.RequestBody, &ob); err != nil {
		result["message"] = err.Error()
		return
	}

	scheduleID, err := strconv.Atoi(b.Ctx.Input.Param(":schedule_id"))
	if err != nil {
		result["message"] = err.Error()
		return
	}

	ob.Id = scheduleID
	if err := ob.InsertOrUpdate(); err != nil {
		beego.Error(err)
		result["message"] = err.Error()
		return
	}

	result["ok"] = true
}

// Delete ...
func (b *ScheduleController) Delete() {
	result := Json{
		"ok": false,
	}

	defer func() {
		b.Data["json"] = result
		b.ServeJSON()
	}()

	scheduleID, err := strconv.Atoi(b.Ctx.Input.Param(":schedule_id"))
	if err != nil {
		result["message"] = err.Error()
		return
	}

	ob := models.Schedule{Id: scheduleID}
	if err := ob.Delete(); err != nil {
		result["message"] = err.Error()
		return
	}

	result["ok"] = true
}
