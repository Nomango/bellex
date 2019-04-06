package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Index() {
	c.TplName = "index.html"
}

func (c *AdminController) Login() {
	c.TplName = "login.html"
}
