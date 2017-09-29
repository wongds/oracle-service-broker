package services

import (
	"encoding/json"
	"errors"

	"github.com/golang/glog"
	"github.com/kubernetes-incubator/service-catalog/pkg/brokerapi"

	"github.com/astaxie/beego/logs"
)

var (
	oracleServiceBroker *OracleServiceBroker
)

type OracleServiceBroker struct {
}

func OracleServiceBrokerInstance() *OracleServiceBroker {
	if oracleServiceBroker == nil {
		oracleServiceBroker = &OracleServiceBroker{
		}
	}
	return oracleServiceBroker
}

func (o *OracleServiceBroker) Catalog() *brokerapi.Catalog {
	return readBrokerSettings()
}

func (o *OracleServiceBroker) ServiceInstance(id string) (string, error) {
	result := ""

	client := GetEtcdClientInstance()
	if client == nil {
		return "", errors.New("Create etcd client instance failure.")
	}
	response, err := client.Get("/serviceinstance/" + id)
	if err != nil {
		return "", errors.New("Get instance failre. The instance id is " + id)
	}
	var serviceInstance userProvidedServiceInstance
	json.Unmarshal([]byte(response.Node.Value), &serviceInstance)
	ok := true
	if serviceInstance.Name == "" {
		ok = false
	}
	if ok {
		result = id
	}

	return result, nil
}

func (o *OracleServiceBroker) Provision(id string, req *brokerapi.CreateServiceInstanceRequest) (*brokerapi.CreateServiceInstanceResponse, error) {

	DashboardURL := ""

	credential, ok := req.Parameters["credentials"]

	if !ok {
		return nil, errors.New("Parameters need to be provided \\'credential\\'")
	} else {
		jsonCred, err := json.Marshal(credential)
		if err != nil {
			glog.Errorf("Failed to marshal credentials: %v", err)
			return nil, err
		}
		cred := &brokerapi.Credential{}
		err = json.Unmarshal(jsonCred, cred)

		connectURI := (*cred)["connect_uri"]
		serviceName := (*cred)["service_name"]
		planName := (*cred)["plan_name"]
		if connectURI == "" {
			return nil, errors.New("Parameters need to be provided \\'connect_uri\\'")
		}
		if serviceName == "" {
			return nil, errors.New("Parameters need to be provided \\'service_name\\'")
		}
		if planName == "" {
			return nil, errors.New("Parameters need to be provided \\'plan_name\\'")
		}
		logs.Info("Parameters \"connect_uri\" : ", connectURI)
		logs.Info("Parameters \"service_name\" : ", serviceName)
		logs.Info("Parameters \"plan_name\" : ", planName)


		plan := getEqualPlan(serviceName.(string), planName.(string))
		if plan == nil {
			return nil, errors.New("Plan not found.Please select corrected plan.")
		}
		valueUnit := getPlanValue(plan)
		if valueUnit == "" {
			logs.Info("Get plan value equals \"\".")
			return nil, errors.New("Get plan value equals \"\".")
		}

		databaseName, userName, userPassword, err := createDatabaseAndUser(connectURI.(string), valueUnit, true)
		if err != nil {
			logs.Info(err)
			return nil, errors.New("CRUD - Create database and user error.")
		}
		DashboardURL = generateOracleUri(userName, userPassword, connectURI.(string), databaseName)

		(*cred)["gen_database"] = databaseName
		(*cred)["gen_username"] = userName
		(*cred)["gen_password"] = userPassword

		//o.instanceMap[id] = &userProvidedServiceInstance{
		//	Name:       id,
		//	Credential: cred,
		//}

		contents, err := json.Marshal(userProvidedServiceInstance{
			Name:       id,
			Credential: cred,
		})
		if err != nil {
			logs.Info("Instance info switch failue.")
			return nil, errors.New("Instance info switch failue.")
		}
		client := GetEtcdClientInstance()
		if client == nil {
			logs.Info("Create etcd client instance failure.")
			return nil, errors.New("Create etcd client instance failure.")
		}
		logs.Info(string(contents))
		client.Set("/serviceinstance/"+id, string(contents))
	}

	//glog.Info("instance map len :", len(o.instanceMap))
	//glog.Info("instance map :", o.instanceMap)

	return &brokerapi.CreateServiceInstanceResponse{
		Operation:    "Provision",
		DashboardURL: DashboardURL,
	}, nil
}

func (o *OracleServiceBroker) DeProvision(id string) (*brokerapi.DeleteServiceInstanceResponse, error) {
	//o.rwMutex.Lock()
	//defer o.rwMutex.Unlock()

	//TODO: Need to be replace by etcd v3.
	//instance, ok := o.instanceMap[id]
	client := GetEtcdClientInstance()
	if client == nil {
		return nil, errors.New("Create etcd client instance failure.")
	}
	response, err := client.Get("/serviceinstance/" + id)
	if err != nil {
		return nil, errors.New("Get instance failre. The instance id is " + id)
	}
	var serviceInstance userProvidedServiceInstance
	json.Unmarshal([]byte(response.Node.Value), &serviceInstance)
	ok := true
	if serviceInstance.Name == "" {
		ok = false
	}

	if ok {
		cred := serviceInstance.Credential
		connectURI := (*cred)["connect_uri"]
		if connectURI == "" {
			return nil, errors.New("Parameters cann't to be obtain \\'connect_uri\\'")
		}
		databaseName := (*cred)["gen_database"]
		if databaseName == "" {
			return nil, errors.New("Parameters cann't to be obtain \\'gen_database\\'")
		}
		userName := (*cred)["gen_username"]
		if userName == "" {
			return nil, errors.New("Parameters cann't to be obtain \\'gen_username\\'")
		}

		err := deleteDatabaseAndUser(connectURI.(string), databaseName.(string), userName.(string))
		if err != nil {
			return nil, errors.New("CRUD - Delete database and user error.")
		}

		//delete(o.instanceMap, id)
		client.Delete("/serviceinstance/" + id)

		return &brokerapi.DeleteServiceInstanceResponse{
			Operation: "DeProvision",
		}, nil
	}

	return &brokerapi.DeleteServiceInstanceResponse{
		Operation: "DeProvision",
	}, nil
}

func (o *OracleServiceBroker) Binding(instanceID, bindingID string, req *brokerapi.BindingRequest) (*brokerapi.CreateServiceBindingResponse, error) {
	//o.rwMutex.RLock()
	//defer o.rwMutex.RUnlock()

	//TODO: Need to be replace by etcd v3.
	//instance, ok := o.instanceMap[instanceID]
	client := GetEtcdClientInstance()
	if client == nil {
		return nil, errors.New("Create etcd client instance failure.")
	}
	response, err := client.Get("/serviceinstance/" + instanceID)
	if err != nil {
		return nil, errors.New("Get instance failre. The instance id is " + instanceID)
	}
	var serviceInstance userProvidedServiceInstance
	json.Unmarshal([]byte(response.Node.Value), &serviceInstance)
	if serviceInstance.Name == "" {
		return nil, errors.New("no such instance: " + instanceID)
	}
	//if !ok {
	//	return nil, errors.New("no such instance: " + instanceID)
	//}
	//cred := instance.Credential
	// remove connect_uri from service instance credential.
	credential := serviceInstance.Credential
	delete(*credential, "connect_uri")
	serviceInstance.Credential = credential
	return &brokerapi.CreateServiceBindingResponse{
		Credentials: *(serviceInstance.Credential),
	}, nil
}

func (o *OracleServiceBroker) UnBinding(instanceId, bindingId string) error {
	// Since we don't persist the binding, there's nothing to do here.
	return nil
}
