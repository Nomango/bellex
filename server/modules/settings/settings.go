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

func IsDevelopeMode() bool {
	return os.Getenv("BELLEX_MODE") == "develope"
}

func Setup() {
	beego.SetViewsPath("views")
	beego.SetStaticPath("/static", "static")

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "bellex_session"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./session"

	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.EnableXSRF = false

	// flash name
	beego.BConfig.WebConfig.FlashName = "BELLEX_FLASH"
	beego.BConfig.WebConfig.FlashSeparator = "BELLEXLASH"

	if IsDevelopeMode() {
		Mode = ModeDevelope
	} else {
		Mode = ModeProduct
	}
}
