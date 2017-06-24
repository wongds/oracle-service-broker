package services

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/astaxie/beego"
	"github.com/ghodss/yaml"
	"github.com/golang/glog"
	"github.com/kubernetes-incubator/service-catalog/pkg/brokerapi"
)

func readBrokerSettings() *brokerapi.Catalog {
	section, _ := beego.AppConfig.GetSection("oracle-service-broker")

	bytes, err := ioutil.ReadFile(section["settings.path"])

	if err != nil {
		glog.Fatalln("load settings.yaml error...", err.Error())
	}

	broker := &brokerapi.Catalog{}

	err = yaml.Unmarshal(bytes, broker)
	if err != nil {
		glog.Fatalln("yaml unmarshal err...", err.Error())
	}

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
			return plan
		}
	}
	return nil
}

func generateOracleUri(username, password, address, sid string) string {
	return fmt.Sprintf("%s:%s@%s/%s", username, password, address, sid)
}

func generateOracleUriWithSid(username, password, addressWithSid string) string {
	return fmt.Sprintf("%s:%s@%s", username, password, addressWithSid)
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
