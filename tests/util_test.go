package test

import (
	"flag"
	"testing"
	"io/ioutil"

	"github.com/ghodss/yaml"
	"github.com/golang/glog"
	"github.com/kubernetes-incubator/service-catalog/pkg/brokerapi"
)

func TestGetEqualPlan(t *testing.T) {

	plan := getEqualPlan("oracle-service", "default")
	if plan == nil {
		t.Errorf("Plan %s not found.Please select corrected plan.", "default")
	}

	t.Log(plan.Metadata)

}

func readBrokerSettings() *brokerapi.Catalog {
	flag.Parse()

	bytes, err := ioutil.ReadFile("../conf/settings.yaml")
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
			return &plan
		}
	}
	return nil
}
