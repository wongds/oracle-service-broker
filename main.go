package main

import (
	"flag"

	_ "github.com/compassorg/oracle-service-broker/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	flag.Parse()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/docs"] = "swagger"
	}
	// For Beego CORS
	var corsFilter = cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:     []string{"HEAD", "GET", "PUT", "POST", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Authorization", "Origin", "X-Requested-With", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})
	beego.InsertFilter("/*", beego.BeforeRouter, corsFilter)
	beego.Run()
}
