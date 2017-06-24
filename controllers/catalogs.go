package controllers

import (
	"oracle-service-broker/services"

	"github.com/astaxie/beego"
)

// Operations about Catalogs
type CatalogsController struct {
	beego.Controller
}

// @Title List broker catalogs.
// @Description List broker catalogs and association plans.
// @Success 200 {object} map[string]string
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router / [get]
func (c *CatalogsController) Catalogs() {
	catalog := services.OracleServiceBrokerInstance().Catalog()

	c.Data["json"] = catalog

	c.ServeJSON()
}
