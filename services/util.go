package services

import (
	"io/ioutil"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/ghodss/yaml"
	"github.com/pivotal-cf/brokerapi"
)

func ReadBrokerSettings() *brokerapi.CatalogResponse {
	section, _ := beego.AppConfig.GetSection("oracle-service-broker")

	bytes, err := ioutil.ReadFile(section["settings.path"])

	if err != nil {
		glog.Fatalln("load settings.yaml error...", err.Error())
	}

	broker := &brokerapi.CatalogResponse{

	}

	err = yaml.Unmarshal(bytes, broker)
	if err != nil {
		glog.Fatalln("yaml unmarshal err...", err.Error())
	}

	return broker
}
