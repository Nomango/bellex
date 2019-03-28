package models

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Setup() {
	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		log.Fatalln("Register driver failed", err)
	}
	if err := orm.RegisterDataBase("default", "mysql", "bellex:Bellex2019@tcp(127.0.0.1:3306)/bellex?charset=utf8"); err != nil {
		log.Fatalln("Register database failed", err)
	}

	if err := orm.RunSyncdb("default", false, true); err != nil {
		beego.Error(err)
	}
}
