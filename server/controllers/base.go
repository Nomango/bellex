package controllers

import (
	"github.com/astaxie/beego"
	"github.com/nomango/bellex/server/models"
)

const (
	sessionUserKey = "auth_user_id"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
	User    models.User
}

type NestPreparer interface {
	NestPrepare()
}

// Prepare implements Prepare method for beego.Controller
func (b *BaseController) Prepare() {
	b.IsLogin = b.GetUserFromSession(&b.User)

	if b.IsLogin {
		// if user forbided then do logout
		if b.User.IsForbid {
			b.LogoutUser()
			b.FlashError("用户权限已被封锁")
			b.Redirect("/login", 302)
			return
		}
	}

	if app, ok := b.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (b *BaseController) LoginUser(user *models.User) {
	b.SetSession(sessionUserKey, user.Id)
	// ctx.Input.CruSession.SessionRelease(ctx.ResponseWriter)
	// ctx.Input.CruSession = beego.GlobalSessions.SessionRegenerateID(ctx.ResponseWriter, ctx.Request)
	// ctx.Input.CruSession.Set(sessionUserKey, user.Id)
}

func (b *BaseController) LogoutUser() {
	b.DelSession(sessionUserKey)
	// ctx.Input.CruSession.Delete(sessionUserKey)
	// ctx.Input.CruSession.Flush()
	// beego.GlobalSessions.SessionDestroy(ctx.ResponseWriter, ctx.Request)
}

func (b *BaseController) GetUserFromSession(user *models.User) bool {
	id, ok := b.GetSession(sessionUserKey).(int)
	if !ok || id <= 0 {
		return false
	}
	u := models.User{Id: id}
	if u.Read() == nil {
		*user = u
		return true
	}
	return false
}

func (b *BaseController) FlashSuccess(msg string, args ...interface{}) {
	flash := beego.NewFlash()
	flash.Success(msg, args...)
	flash.Store(&b.Controller)
}

func (b *BaseController) FlashNotice(msg string, args ...interface{}) {
	flash := beego.NewFlash()
	flash.Notice(msg, args...)
	flash.Store(&b.Controller)
}

func (b *BaseController) FlashWarning(msg string, args ...interface{}) {
	flash := beego.NewFlash()
	flash.Warning(msg, args...)
	flash.Store(&b.Controller)
}

func (b *BaseController) FlashError(msg string, args ...interface{}) {
	flash := beego.NewFlash()
	flash.Error(msg, args...)
	flash.Store(&b.Controller)
}
