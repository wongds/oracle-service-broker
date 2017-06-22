package services

import (
	"fmt"
	"io/ioutil"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/ghodss/yaml"
	"github.com/kubernetes-incubator/service-catalog/pkg/brokerapi"
)

func readBrokerSettings() *brokerapi.Catalog {
	section, _ := beego.AppConfig.GetSection("oracle-service-broker")

	bytes, err := ioutil.ReadFile(section["settings.path"])

	if err != nil {
		glog.Fatalln("load settings.yaml error...", err.Error())
	}

	broker := &brokerapi.Catalog{

	}

	err = yaml.Unmarshal(bytes, broker)
	if err != nil {
		glog.Fatalln("yaml unmarshal err...", err.Error())
	}

	return broker
}

func generateOracleUri(username, password, address, sid string) string {
	return fmt.Sprintf("%s:%s@%s/%s", username, password, address, sid)
}

func generateOracleUriWithSid(username, password, addressWithSid string) string {
	return fmt.Sprintf("%s:%s@%s", username, password, addressWithSid)
}
