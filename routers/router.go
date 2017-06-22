// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"oracle-service-broker/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v2",
		beego.NSNamespace("/catalog",
			beego.NSInclude(
				&controllers.CatalogsController{},
			),
		),

		beego.NSNamespace("/service_instances",
			beego.NSInclude(
				&controllers.InstancesController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
