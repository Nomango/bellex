package v1

import (
	"encoding/json"
	"strconv"

	"github.com/nomango/bellex/server/models"

	"github.com/astaxie/beego"
)

// BellController ...
type BellController struct {
	beego.Controller
}

// Post ...
func (b *BellController) Post() {
	result := Json{
		"ok": false,
	}

	defer func() {
		b.Data["json"] = result
		b.ServeJSON()
	}()

	var ob models.Bell
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
func (b *BellController) Get() {
	defer b.ServeJSON()

	bellID, err := strconv.Atoi(b.Ctx.Input.Param(":bellId"))
	if err != nil {
		b.Data["json"] = Json{"ok": false, "message": err.Error()}
		return
	}

	bell := models.Bell{Id: bellID}
	if err := bell.Read(); err != nil {
		b.Data["json"] = Json{"ok": false, "message": err.Error()}
		return
	}
	b.Data["json"] = Json{"ok": true, "data": bell}
}

// GetAll ...
func (b *BellController) GetAll() {
	var bells []*models.Bell
	if _, err := models.Bells().All(&bells); err != nil {
		b.Data["json"] = Json{"ok": false, "message": err.Error()}
	} else {
		b.Data["json"] = Json{"ok": true, "data": bells}
	}
	b.ServeJSON()
}

// Put ...
func (b *BellController) Put() {
	result := Json{
		"ok": false,
	}

	defer func() {
		b.Data["json"] = result
		b.ServeJSON()
	}()

	var ob models.Bell
	if err := json.Unmarshal(b.Ctx.Input.RequestBody, &ob); err != nil {
		result["message"] = err.Error()
		return
	}

	bellID, err := strconv.Atoi(b.Ctx.Input.Param(":bellId"))
	if err != nil {
		result["message"] = err.Error()
		return
	}

	ob.Id = bellID
	if err := ob.InsertOrUpdate(); err != nil {
		beego.Error(err)
		result["message"] = err.Error()
		return
	}

	result["ok"] = true
}

// Delete ...
func (b *BellController) Delete() {
	result := Json{
		"ok": false,
	}

	defer func() {
		b.Data["json"] = result
		b.ServeJSON()
	}()

	bellID, err := strconv.Atoi(b.Ctx.Input.Param(":bellId"))
	if err != nil {
		result["message"] = err.Error()
		return
	}

	ob := models.Bell{Id: bellID}
	if err := ob.Delete(); err != nil {
		result["message"] = err.Error()
		return
	}

	result["ok"] = true
}
