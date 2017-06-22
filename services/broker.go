package services

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/golang/glog"
	"github.com/kubernetes-incubator/service-catalog/pkg/brokerapi"
)

var (
	oracleServiceBroker *OracleServiceBroker
)

type OracleServiceBroker struct {
	rwMutex sync.RWMutex
	//TODO: Need to be replace by etcd v3.
	instanceMap map[string]*userProvidedServiceInstance
}

func OracleServiceBrokerInstance() *OracleServiceBroker {
	if oracleServiceBroker == nil {
		//TODO: Need to be replace by etcd v3.
		var instanceMap = make(map[string]*userProvidedServiceInstance)
		oracleServiceBroker = &OracleServiceBroker{
			instanceMap: instanceMap,
		}
	}
	return oracleServiceBroker
}

func (o *OracleServiceBroker) Catalog() *brokerapi.Catalog {
	return readBrokerSettings()
}

func (o *OracleServiceBroker) ServiceInstance(id string) (string, error) {
	result := ""

	//TODO: Need to be replace by etcd v3.
	if _, ok := o.instanceMap[id]; ok {
		result = id
	}

	glog.Info("instance map len :", len(o.instanceMap))
	glog.Info("instance map :", o.instanceMap)

	return result, nil
}

func (o *OracleServiceBroker) Provision(id string, req *brokerapi.CreateServiceInstanceRequest) (*brokerapi.CreateServiceInstanceResponse, error) {
	o.rwMutex.Lock()
	defer o.rwMutex.Unlock()

	credential, ok := req.Parameters["credentials"]

	if !ok {
		o.instanceMap[id] = &userProvidedServiceInstance{
			Name: id,
			Credential: &brokerapi.Credential{
				"special-key-1": "special-value-1",
				"special-key-2": "special-value-2",
			},
		}
	} else {
		jsonCred, err := json.Marshal(credential)
		if err != nil {
			glog.Errorf("Failed to marshal credentials: %v", err)
			return nil, err
		}
		var cred brokerapi.Credential
		err = json.Unmarshal(jsonCred, &cred)
		o.instanceMap[id] = &userProvidedServiceInstance{
			Name:       id,
			Credential: &cred,
		}
	}

	glog.Info("instance map len :", len(o.instanceMap))
	glog.Info("instance map :", o.instanceMap)

	return &brokerapi.CreateServiceInstanceResponse{
		Operation: "Provision",
	}, nil
}

func (o *OracleServiceBroker) DeProvision(id string) (*brokerapi.DeleteServiceInstanceResponse, error) {
	o.rwMutex.Lock()
	defer o.rwMutex.Unlock()

	//TODO: Need to be replace by etcd v3.
	_, ok := o.instanceMap[id]

	if ok {
		delete(o.instanceMap, id)
		return &brokerapi.DeleteServiceInstanceResponse{}, nil
	}

	return &brokerapi.DeleteServiceInstanceResponse{
		Operation: "DeProvision",
	}, nil
}

func (o *OracleServiceBroker) Binding(instanceID, bindingID string, req *brokerapi.BindingRequest) (*brokerapi.CreateServiceBindingResponse, error) {
	o.rwMutex.RLock()
	defer o.rwMutex.RUnlock()

	//TODO: Need to be replace by etcd v3.
	instance, ok := o.instanceMap[instanceID]
	if !ok {
		return nil, errors.New("no such instance: " + instanceID)
	}
	cred := instance.Credential
	return &brokerapi.CreateServiceBindingResponse{Credentials: *cred}, nil
}

func (o *OracleServiceBroker) UnBinding(instanceId, bindingId string) error {
	// Since we don't persist the binding, there's nothing to do here.
	return nil
}
