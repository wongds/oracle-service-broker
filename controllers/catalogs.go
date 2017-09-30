package controllers

import (
	"net/http"

	"github.com/compassorg/oracle-service-broker/services"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Operations about Catalogs
type CatalogsController struct {
	beego.Controller
}

// @Title List broker catalogs.
// @Description List broker catalogs and association plans.
// @Param    Authorization        header     string    "Basic token"    true    "Authorization information."
// @Success 200 {object} map[string]string
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router / [get]
func (c *CatalogsController) Catalogs() {
	logs.Info("Invoke Oracle Service Broker API Reques Header : ", c.Ctx.Request.Header)
	catalog := services.OracleServiceBrokerInstance().Catalog()

	writeCatalogResponse(c, http.StatusOK, catalog)
}

// WriteResponse will serialize 'object' to the HTTP ResponseWriter
// using the 'code' as the HTTP status code
func writeCatalogResponse(c *CatalogsController, code int, object interface{}) {
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.Ctx.Output.Header("Content-Type", "application/json")
	c.Data["json"] = object
	c.ServeJSON()
}
