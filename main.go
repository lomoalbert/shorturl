package main

import (
	_ "github.com/lomoalbert/shorturl/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
