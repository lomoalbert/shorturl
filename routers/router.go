// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/lomoalbert/shorturl/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/shorturl",
			beego.NSInclude(
				&controllers.ShortURLController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.Router("/:code([0-9a-zA-Z]+)", &controllers.RedirectController{})
	beego.Router("/", &controllers.HomeController{})
}
