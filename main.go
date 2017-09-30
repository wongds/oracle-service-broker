package main

import (
	"flag"
	"strings"
	"net/http"

	_ "github.com/compassorg/oracle-service-broker/routers"
	"github.com/compassorg/oracle-service-broker/util"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"encoding/base64"
)

func main() {
	flag.Parse()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/docs"] = "swagger"
	}

	var tokenFilter = func(ctx *context.Context) {
		if strings.ToUpper(ctx.Request.Method) != "OPTIONS" {
			authorization := ctx.Request.Header.Get("Authorization")
			if authorization == "" {
				ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
			}
			if authorization != "" {
				section, _ := beego.AppConfig.GetSection("auth")
				username := section["username"]
				password := section["password"]
				localAuthStr := username+":"+password
				localBase64AuthStr := base64.StdEncoding.EncodeToString([]byte(localAuthStr))
				reqBase64Str := util.GetToken(authorization)
				if !util.CheckEqual(reqBase64Str, localBase64AuthStr) {
					ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
				}
			}
		}
	}
	beego.InsertFilter("/*", beego.BeforeRouter, tokenFilter)

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