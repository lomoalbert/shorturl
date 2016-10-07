package models

import (

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.Debug = beego.AppConfig.DefaultBool("DBDebug", false)

	orm.RegisterModel(new(SURL),
	)
	orm.RunSyncdb("default", false, true)
}

func NewOrm() orm.Ormer {
	return orm.NewOrm()
}