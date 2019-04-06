package settings

import (
	"os"

	"github.com/astaxie/beego"
)

var (
	Mode int
)

const (
	ModeDevelope = iota
	ModeProduct
)

func Setup() {
	beego.SetViewsPath("views")
	beego.SetStaticPath("/static", "static")

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "bellex-key"
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"

	if os.Getenv("BELLEX_MODE") == "develope" {
		Mode = ModeDevelope
	} else {
		Mode = ModeProduct
	}
}
