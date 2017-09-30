package controllers

import (
	"net/http"

	"github.com/compassorg/oracle-service-broker/services"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/kubernetes-incubator/service-catalog/pkg/brokerapi"
)

// Operations about ServiceInstances
type InstancesController struct {
	beego.Controller
}

// @Title Get instance with instanceID.
// @Description Get service instance with instanceID.
// @Param    Authorization        header     string    "Basic token"    true    "Authorization information."
// @Param instance_id path string true "Instance id."
// @Success 200 {object} map[string]string
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router /:instance_id [get]
func (i *InstancesController) Instance() {
	instanceId := i.GetString(":instance_id")

	glog.Infof("Get service instance %s...\n", instanceId)

	if result, err := services.OracleServiceBrokerInstance().ServiceInstance(instanceId); err == nil {
		writeResponse(i, http.StatusOK, result)
	} else {
		writeErrResponse(i, http.StatusBadRequest, err)
	}
}

// @Title Provision service instance.
// @Description Provision service instance with serviceClass & plan.
// @Param    Authorization        header     string    "Basic token"    true    "Authorization information."
// @Param instance_id path string true "Instance id."
// @Param body body brokerapi.CreateServiceInstanceRequest true "Request Body."
// @Success 200 {object} map[string]string
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router /:instance_id [put]
func (i *InstancesController) Provision() {
	instanceId := i.GetString(":instance_id")

	glog.Infof("Provision service instance %s...\n", instanceId)

	var req brokerapi.CreateServiceInstanceRequest
	if err := BodyToObject(i.Ctx.Request, &req); err != nil {
		glog.Errorf("error unmarshalling: %v", err)
		writeErrResponse(i, http.StatusBadRequest, err)
		return
	}

	// TODO: Check if parameters are required, if not, this thing below is ok to leave in,
	// if they are ,they should be checked. Because if no parameters are passed in, this will
	// fail because req.Parameters is nil.
	if req.Parameters == nil {
		req.Parameters = make(map[string]interface{})
	}

	if result, err := services.OracleServiceBrokerInstance().Provision(instanceId, &req); err == nil {
		writeResponse(i, http.StatusCreated, result)
	} else {
		writeErrResponse(i, http.StatusBadRequest, err)
	}
}

// @Title DeProvision service instance.
// @Description DeProvision service instance with instanceId.
// @Param    Authorization        header     string    "Basic token"    true    "Authorization information."
// @Param instance_id path string true "Instance id."
// @Success 200 {object} map[string]string
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router /:instance_id [delete]
func (i *InstancesController) DeProvision() {
	instanceId := i.GetString(":instance_id")

	glog.Infof("DeProvision service instance %s...\n", instanceId)

	if result, err := services.OracleServiceBrokerInstance().DeProvision(instanceId); err == nil {
		writeResponse(i, http.StatusOK, result)
	} else {
		writeErrResponse(i, http.StatusBadRequest, err)
	}
}

// @Title Binding service instance.
// @Description Binding service instance with instanceId and bindingId.
// @Param    Authorization        header     string    "Basic token"    true    "Authorization information."
// @Param instance_id path string true "Instance id."
// @Param binding_id path string true "Binding id."
// @Success 200 {object} map[string]string
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router /:instance_id/service_bindings/:binding_id [put]
func (i *InstancesController) Binding() {
	instanceId := i.GetString(":instance_id")
	bindingId := i.GetString(":binding_id")

	glog.Infof("Binding binding_id=%s, instance_id=%s\n", bindingId, instanceId)

	var req brokerapi.BindingRequest

	if err := BodyToObject(i.Ctx.Request, &req); err != nil {
		glog.Errorf("Failed to unmarshall request: %v", err)
		writeErrResponse(i, http.StatusBadRequest, err)
		return
	}

	// TODO: Check if parameters are required, if not, this thing below is ok to leave in,
	// if they are ,they should be checked. Because if no parameters are passed in, this will
	// fail because req.Parameters is nil.
	if req.Parameters == nil {
		req.Parameters = make(map[string]interface{})
	}

	// Pass in the instanceId to the template.
	req.Parameters["instanceId"] = instanceId

	if result, err := services.OracleServiceBrokerInstance().Binding(instanceId, bindingId, &req); err == nil {
		writeResponse(i, http.StatusOK, result)
	} else {
		writeErrResponse(i, http.StatusBadRequest, err)
	}

}

// @Title UnBinding service instance.
// @Description UnBinding service instance with instanceId and bindingId.
// @Param    Authorization        header     string    "Basic token"    true    "Authorization information."
// @Param instance_id path string true "Instance id."
// @Param binding_id path string true "Binding id."
// @Success 200 {object} map[string]string
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router /:instance_id/service_bindings/:binding_id [delete]
func (i *InstancesController) UnBinding() {
	instanceId := i.GetString(":instance_id")
	bindingId := i.GetString(":binding_id")

	glog.Infof("UnBinding srvice instance guid: %s:%s", instanceId, bindingId)

	if err := services.OracleServiceBrokerInstance().UnBinding(instanceId, bindingId); err == nil {
		writeResponse(i, http.StatusOK, bindingId)
	} else {
		writeErrResponse(i, http.StatusBadRequest, err)
	}
}

// @Title Get instance last operation with instanceID.
// @Description Get service instance with instanceID.
// @Param    Authorization        header     string    "Basic token"    true    "Authorization information."
// @Param instance_id path string true "Instance id."
// @Param service_id query string true "Service id."
// @Param plan_id query string true "Plan id."
// @Param operation query string true "Operation."
// @Success 200 {object} map[string]string
// @Failure 400 {object} error
// @Failure 500 {object} error
// @router /:instance_id/last_operation [get]
func (i *InstancesController) ServiceInstanceLastOperation() {
	instanceId := i.GetString(":instance_id")
	serviceId := i.GetString("service_id")
	planId := i.GetString("plan_id")
	operation := i.GetString("operation")

	glog.Infof("Get service instance last opertaion %s...\n", instanceId)

	if result, err := services.OracleServiceBrokerInstance().ServiceInstanceLastOperation(instanceId, serviceId, planId, operation); err == nil {
		writeResponse(i, http.StatusOK, result)
	} else {
		writeErrResponse(i, http.StatusBadRequest, err)
	}
}

// WriteResponse will serialize 'object' to the HTTP ResponseWriter
// using the 'code' as the HTTP status code
func writeResponse(i *InstancesController, code int, object interface{}) {
	i.Ctx.ResponseWriter.WriteHeader(code)
	i.Ctx.Output.Header("Content-Type", "application/json")
	i.Data["json"] = object
	i.ServeJSON()
}

// WriteResponse will serialize 'object' to the HTTP ResponseWriter
// using the 'code' as the HTTP status code
func writeErrResponse(i *InstancesController, code int, err error) {
	i.Ctx.ResponseWriter.WriteHeader(code)
	i.Ctx.Output.Header("Content-Type", "application/json")
	i.Data["json"] = err.Error()
	i.ServeJSON()
}
