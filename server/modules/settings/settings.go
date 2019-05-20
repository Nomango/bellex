package settings

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/astaxie/beego"
	"github.com/spf13/viper"
)

var (
	Debug   bool
	AppName string
	AppAddr string
	AppPort string

	SessionName           string
	SessionProvider       string
	SessionProviderConfig string

	DatabaseUser     string
	DatabasePassword string
	DatabaseUri      string

	TcpPort string
)

const (
	configDevName  = "dev"
	configProdName = "prod"
	configPath     = "../conf/"
	configFormat   = "yaml"
)

func IsDevelopeMode() bool {
	return os.Getenv("BELLEX_MODE") == "develope"
}

func Setup() {
	ReadSettings()

	beego.SetViewsPath("views")
	beego.SetStaticPath("/static", "static")
	beego.SetStaticPath("/download/desktop", "download")

	if Debug {
		beego.BConfig.RunMode = "dev"
	} else {
		beego.BConfig.RunMode = "prod"
	}

	beego.BConfig.AppName = AppName
	beego.BConfig.Listen.HTTPPort, _ = strconv.Atoi(AppPort)
	beego.BConfig.Listen.HTTPAddr = AppAddr

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = SessionName
	beego.BConfig.WebConfig.Session.SessionProvider = SessionProvider
	beego.BConfig.WebConfig.Session.SessionProviderConfig = SessionProviderConfig

	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.EnableXSRF = false

	// flash name
	beego.BConfig.WebConfig.FlashName = "BELLEX_FLASH"
	beego.BConfig.WebConfig.FlashSeparator = "BELLEXLASH"
}

func ReadSettings() {
	v := getViper()
	v.SetConfigType(configFormat)
	v.AddConfigPath(configPath)

	if IsDevelopeMode() {
		v.SetConfigName(configDevName)
	} else {
		v.SetConfigName(configProdName)
	}

	if err := v.ReadInConfig(); err != nil {
		log.Panicln("load config file failed", err)
	}

	Debug = v.GetBool("app.debug")
	AppName = v.GetString("app.name")
	AppAddr = v.GetString("app.address")
	AppPort = v.GetString("app.port")

	SessionName = v.GetString("app.session.name")
	SessionProvider = v.GetString("app.session.provider")
	SessionProviderConfig = v.GetString("app.session.providerConfig")

	DatabaseUser = v.GetString("database.user")
	DatabasePassword = v.GetString("database.password")
	DatabaseUri = v.GetString("database.uri")

	TcpPort = v.GetString("tcp.port")
}

var (
	once      sync.Once
	vInstance *viper.Viper
)

func getViper() *viper.Viper {
	once.Do(func() {
		vInstance = viper.New()
	})
	return vInstance
}
