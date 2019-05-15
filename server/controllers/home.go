package controllers

import (
	"github.com/nomango/bellex/server/models"
)

const downloadBucketURL = ""

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	if !c.IsLogin {
		c.Redirect("/login", 302)
		return
	}

	c.TplName = "index.html"
}

func (c *HomeController) Login() {
	if c.IsLogin {
		c.Redirect("/", 302)
		return
	}

	c.TplName = "login.html"
}

func (c *HomeController) DownloadLatest() {
	if ver, err := models.GetLatestVersion(); err != nil {
		c.Abort("400")
	} else {
		c.Redirect(ver.URL, 302)
	}
}
