package services

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"time"

	"github.com/compassorg/oracle-service-broker/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/client"
	"github.com/ghodss/yaml"
	"github.com/golang/glog"
	"github.com/kubernetes-incubator/service-catalog/pkg/brokerapi"
	"golang.org/x/net/context"
	yamlV2 "gopkg.in/yaml.v2"
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
			return &plan
		}
	}
	return nil
}

func getPlanValue(plan *brokerapi.ServicePlan) string {
	bytes, err := yamlV2.Marshal(plan.Metadata)
	if err != nil {
		glog.Errorf("Yaml Unmarshal error. %s", err.Error())
		return ""
	}

	metadata := models.Metadata{}
	err = yamlV2.Unmarshal(bytes, &metadata)
	if err != nil {
		glog.Errorf("Yaml Marshal error. %s", err.Error())
		return ""
	}

	// ignore all values but first value
	valueUnit := ""
	if len(metadata.Costs) > 0 {
		cost := metadata.Costs[0]
		valueUnit, _ = fmt.Printf("%s%s", cost.Amount.Value, cost.Unit)
	}

	return valueUnit
}

func generateOracleUri(username, password, address, sid string) string {
	return fmt.Sprintf("%s/%s@%s/%s", username, password, address, sid)
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

// For etcd client util.
var etcdClient *EtcdClient = nil

func GetEtcdClientInstance() *EtcdClient {
	if etcdClient == nil {
		section, _ := beego.AppConfig.GetSection("etcd")
		endpointString := section["endpoints"]
		endpoints := strings.Split(endpointString, ",")
		config := client.Config{
			Endpoints:               endpoints,
			Transport:               client.DefaultTransport,
			HeaderTimeoutPerRequest: time.Second * 5,
		}
		c, err := client.New(config)
		if err != nil {
			logs.Error("Can not init etcd client.", err)
			return etcdClient
		}
		etcdClient = &EtcdClient{
			EtcdApi: client.NewKeysAPI(c),
		}
	}
	return etcdClient
}

type EtcdClient struct {
	EtcdApi client.KeysAPI
}

func (e *EtcdClient) Get(key string) (*client.Response, error) {
	n := 5

RETRY:
	resp, err := e.EtcdApi.Get(context.Background(), key, nil)
	if err != nil {
		logs.Error("Can not get "+key+" from etcd", err)
		n--
		if n > 0 {
			goto RETRY
		}

		return nil, err
	} else {
		logs.Debug("Successful get " + key + " from etcd. value is " + resp.Node.Value)
		return resp, nil
	}
}

func (e *EtcdClient) Set(key string, value string) (*client.Response, error) {
	n := 5

RETRY:
	resp, err := e.EtcdApi.Set(context.Background(), key, value, nil)
	if err != nil {
		logs.Error("Can not set "+key+" from etcd", err)
		n--
		if n > 0 {
			goto RETRY
		}

		return nil, err
	} else {
		logs.Debug("Successful set " + key + " from etcd. value is " + value)
		return resp, nil
	}
}

func (e *EtcdClient) Delete(key string) (*client.Response, error) {
	return e.EtcdApi.Delete(context.Background(), key, &client.DeleteOptions{
		Recursive: true,
		Dir:       true,
	})
}
