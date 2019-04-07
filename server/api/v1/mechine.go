package v1

import (
	"encoding/json"
	"strconv"

	"github.com/nomango/bellex/server/models"

	"github.com/astaxie/beego"
)

// MechineController ...
type MechineController struct {
	LoginValidateController
}

// Post ...
func (b *MechineController) Post() {
	result := Json{
		"ok": false,
	}

	defer func() {
		b.Data["json"] = result
		b.ServeJSON()
	}()

	var ob models.Mechine
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
func (b *MechineController) Get() {
	defer b.ServeJSON()

	mechineID, err := strconv.Atoi(b.Ctx.Input.Param(":mechine_id"))
	if err != nil {
		b.Data["json"] = Json{"ok": false, "message": err.Error()}
		return
	}

	mechine := models.Mechine{Id: mechineID}
	if err := mechine.Read(); err != nil {
		b.Data["json"] = Json{"ok": false, "message": err.Error()}
		return
	}
	b.Data["json"] = Json{"ok": true, "data": mechine}
}

// GetAll ...
func (b *MechineController) GetAll() {
	var mechines []*models.Mechine
	if _, err := models.Mechines().All(&mechines); err != nil {
		b.Data["json"] = Json{"ok": false, "message": err.Error()}
	} else {
		b.Data["json"] = Json{"ok": true, "data": mechines}
	}
	b.ServeJSON()
}

// Put ...
func (b *MechineController) Put() {
	result := Json{
		"ok": false,
	}

	defer func() {
		b.Data["json"] = result
		b.ServeJSON()
	}()

	var ob models.Mechine
	if err := json.Unmarshal(b.Ctx.Input.RequestBody, &ob); err != nil {
		result["message"] = err.Error()
		return
	}

	mechineID, err := strconv.Atoi(b.Ctx.Input.Param(":mechine_id"))
	if err != nil {
		result["message"] = err.Error()
		return
	}

	ob.Id = mechineID
	if err := ob.InsertOrUpdate(); err != nil {
		beego.Error(err)
		result["message"] = err.Error()
		return
	}

	result["ok"] = true
}

// Delete ...
func (b *MechineController) Delete() {
	result := Json{
		"ok": false,
	}

	defer func() {
		b.Data["json"] = result
		b.ServeJSON()
	}()

	mechineID, err := strconv.Atoi(b.Ctx.Input.Param(":mechine_id"))
	if err != nil {
		result["message"] = err.Error()
		return
	}

	ob := models.Mechine{Id: mechineID}
	if err := ob.Delete(); err != nil {
		result["message"] = err.Error()
		return
	}

	result["ok"] = true
}
