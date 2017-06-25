package services

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"flag"
	"io/ioutil"

	"github.com/astaxie/beego"
	"github.com/ghodss/yaml"
	"github.com/golang/glog"
	"github.com/kubernetes-incubator/service-catalog/pkg/brokerapi"
	"log"
)

func readBrokerSettings() *brokerapi.Catalog {
	log.Println("333333333333333")
	flag.Parse()
	log.Println("444444444444444")
	section, _ := beego.AppConfig.GetSection("oracle-service-broker")
	log.Println("555555555555555")
	bytes, err := ioutil.ReadFile(section["settings.path"])
	log.Println("666666666666666")
	if err != nil {
		log.Println("888888888888888888")
		glog.Fatalln("load settings.yaml error...", err.Error())
	}
	log.Println("777777777777777")
	broker := &brokerapi.Catalog{}
	log.Println("99999999999999999")
	err = yaml.Unmarshal(bytes, broker)
	log.Println("ffffffffffffffff")
	if err != nil {
		log.Println("aaaaaaaaaaaaaaaaaaaaaaaa")
		glog.Fatalln("yaml unmarshal err...", err.Error())
		log.Println("bbbbbbbbbbbbbbbbbbbbbbb")
	}
	log.Println("ccccccccccccccccccccc")
	return broker
}

func getEqualService(serviceName string) *brokerapi.Service {
	catalogs := readBrokerSettings()
	for _, service := range catalogs.Services {
		if serviceName == service.Name {
			return service
		}
	}
	return nil
}

func getEqualPlan(serviceName, planName string) *brokerapi.ServicePlan {
	service := getEqualService(serviceName)
	for _, plan := range service.Plans {
		if planName == plan.Name {
			return &plan
		}
	}
	return nil
}

func generateOracleUri(username, password, address, sid string) string {
	return fmt.Sprintf("%s:%s@%s/%s", username, password, address, sid)
}

func generateGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return generateMd5(base64.URLEncoding.EncodeToString(b))
}

func generateMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
