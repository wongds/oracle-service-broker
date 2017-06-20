package services

import (
	"context"

	"github.com/pivotal-cf/brokerapi"
)

type ServiceBroker interface {
	DoProvision(context context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error)
	DoDeProvision(context context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error)
	DoBind(context context.Context, instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error)
	DoUnbind(context context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails) error
	DoUpdate(context context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error)
	DoLastOperation(context context.Context, instanceID, operationData string) (brokerapi.LastOperation, error)
	DoServices(context context.Context) []brokerapi.Service
}

type ServiceBrokerHandler struct {
	handler ServiceBroker
}

func (s *ServiceBrokerHandler) DoProvision(context context.Context, instanceID string, details brokerapi.ProvisionDetails,
	asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	return s.handler.DoProvision(context, instanceID, details, asyncAllowed)
}

func (s *ServiceBrokerHandler) DoDeProvision(context context.Context, instanceID string, details brokerapi.DeprovisionDetails,
	asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	return s.handler.DoDeProvision(context, instanceID, details, asyncAllowed)
}

func (s *ServiceBrokerHandler) DoBind(context context.Context, instanceID, bindingID string,
	details brokerapi.BindDetails) (brokerapi.Binding, error) {
	return s.handler.DoBind(context, instanceID, bindingID, details)
}

func (s *ServiceBrokerHandler) DoUnbind(context context.Context, instanceID, bindingID string,
	details brokerapi.UnbindDetails) error {
	return s.handler.DoUnbind(context, instanceID, bindingID, details)
}

func (s *ServiceBrokerHandler) DoUpdate(context context.Context, instanceID string, details brokerapi.UpdateDetails,
	asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return s.handler.DoUpdate(context, instanceID, details, asyncAllowed)
}

func (s *ServiceBrokerHandler) DoLastOperation(context context.Context, instanceID,
	operationData string) (brokerapi.LastOperation, error) {
	return s.handler.DoLastOperation(context, instanceID, operationData)
}

func (s *ServiceBrokerHandler) DoServices(context context.Context) []brokerapi.Service {
	return s.handler.DoServices(context)
}
