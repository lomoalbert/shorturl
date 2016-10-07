package routers

import (
	"github.com/lomoalbert/shorturl/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
