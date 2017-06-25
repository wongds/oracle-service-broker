package controllers

import (
	"net/http"

	"oracle-service-broker/services"

	"github.com/astaxie/beego"
	"log"
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
	log.Println("111111111111")
	catalog := services.OracleServiceBrokerInstance().Catalog()
	log.Println("222222222222")
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