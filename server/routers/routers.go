// Copyright (C) 2018 Nomango - All Rights Reserved

package routers

import (
	"github.com/nomango/bellex/server/api/v1"
	"github.com/nomango/bellex/server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.AdminController{})

	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/bell",
			beego.NSInclude(
				&v1.BellController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&v1.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
