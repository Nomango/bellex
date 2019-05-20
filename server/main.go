package main

import (
	"github.com/nomango/bellex/server/models"
	"github.com/nomango/bellex/server/modules/settings"
	_ "github.com/nomango/bellex/server/routers"

	"github.com/astaxie/beego"
)

func main() {

	settings.Setup()
	models.Setup()

	beego.Run(":" + settings.AppPort)
}
