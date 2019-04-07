package controllers

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
