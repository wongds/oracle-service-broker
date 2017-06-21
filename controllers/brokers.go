package controllers

import (

	"oracle-service-broker/services"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
)

// Operations about Brokers
type BrokersController struct {
	beego.Controller
}

// @Title List broker catalogs.
// @Description List broker catalogs and association plans.
// @Success 200 {object} map[string]string
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router /catalog [get]
func (b *BrokersController) Catalog() {

	catalog := services.ReadBrokerSettings()

	glog.Info(catalog)

	b.Data["json"] = catalog

	b.ServeJSON()
}
