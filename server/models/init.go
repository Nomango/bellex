package models

import (
	"log"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		log.Fatalln("Register driver failed", err)
	}
	if err := orm.RegisterDataBase("default", "mysql", "bellex:Bellex2019@/bellex?charset=utf8"); err != nil {
		log.Fatalln("Register database failed", err)
	}

	orm.RegisterModel(new(Bell))
}
