// Copyright (C) 2018 Nomango - All Rights Reserved

package routers

import (
	"github.com/nomango/bellex/server/api/v1"
	"github.com/nomango/bellex/server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.AdminController{}, "GET:Index")
	beego.Router("/login", &controllers.AdminController{}, "GET:Login")

	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/mechine",
			beego.NSInclude(
				&v1.MechineController{},
			),
		),
		beego.NSNamespace("/schedule",
			beego.NSInclude(
				&v1.ScheduleController{},
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
