// @APIVersion 1.0.0
// @Title goshopping Server API
// @Description goshopping 的后台服务
// @Contact wind_ns@163.com
// @TermsOfServiceUrl http://
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/thewinds/goshopping/server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/gspapi/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
