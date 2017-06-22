package services

import (
	"github.com/kubernetes-incubator/service-catalog/pkg/brokerapi"
)

type userProvidedServiceInstance struct {
	Name       string
	Credential *brokerapi.Credential
}