package models

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nomango/bellex/server/modules/settings"
)

func Setup() {
	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		log.Fatalln("Register driver failed", err)
	}

	dataSource := fmt.Sprintf("%s:%s@%s", settings.DatabaseUser, settings.DatabasePassword, settings.DatabaseUri)
	if err := orm.RegisterDataBase("default", "mysql", dataSource); err != nil {
		log.Fatalln("Register database failed", err)
	}

	if err := orm.RunSyncdb("default", false, true); err != nil {
		beego.Error(err)
	}

	CreateDefaultInstitution()
	CreateAdminUser()
}
