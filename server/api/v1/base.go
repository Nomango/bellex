package v1

import "github.com/nomango/bellex/server/controllers"

type Json map[string]interface{}

type APIController struct {
	controllers.BaseController
}

func (c *APIController) WriteJson(data interface{}, code int) {
	c.Ctx.ResponseWriter.WriteHeader(404)
	c.Data["json"] = Json{"message": "不存在指定用户"}
	c.ServeJSON()
}

// LoginValidateController ...
type LoginValidateController struct {
	APIController
}

func (c *LoginValidateController) NestPrepare() {
	if !c.IsLogin {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Data["json"] = Json{"message": "Permission denied"}
		c.ServeJSON()
		c.StopRun()
	}
}
