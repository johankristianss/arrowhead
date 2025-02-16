package rpc

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/johankristianss/arrowhead/pkg/core"
	"github.com/johankristianss/arrowhead/pkg/parsers"
	log "github.com/sirupsen/logrus"
)

const (
	GET = iota
	POST
	PUT
	DELETE
)

func httpMethodToString(method int) string {
	switch method {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	default:
		return ""
	}
}

func (mgmt *Management) RegisterSystem(systemReg core.SystemRegistration) (core.System, error) {
	jsonData, err := parsers.MarshalSystemRegistration(systemReg)
	if err != nil {
		return core.System{}, err
	}

	url := mgmt.rpc.buildServiceRegistryURL("/mgmt/systems")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		log.Fatalf("Error creating POST request: %v", err)
		return core.System{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	body, err := makeRequest(mgmt.rpc.http, req, 201, "Failed to register system")
	if err != nil {
		return core.System{}, err
	}

	return parsers.UnmarshalSystemJSON(string(body))
}

func (mgmt *Management) UnregisterSystem(system *core.System) error {
	return mgmt.UnregisterSystemByID(system.ID)
}

func (mgmt *Management) UnregisterSystemByID(systemID int) error {
	url := mgmt.rpc.buildServiceRegistryURL("/mgmt/systems/" + fmt.Sprintf("%d", systemID))
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalf("Error creating DELETE request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	_, err = makeRequest(mgmt.rpc.http, req, 200, "Failed to unregister system")
	return err
}

func (mgmt *Management) GetSystems() ([]core.System, error) {
	systems := []core.System{}
	url := mgmt.rpc.buildServiceRegistryURL("/mgmt/systems?direction=ASC&sort_field=id")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return systems, err
	}

	req.Header.Set("Accept", "*/*")

	body, err := makeRequest(mgmt.rpc.http, req, 200, "Failed to get systems")
	if err != nil {
		return systems, err
	}

	systems, err = parsers.UnmarshalGetSystemsResponseJSON(string(body))
	if err != nil {
		return systems, err
	}

	return systems, nil
}

// TODO: Is this function tested?
func (mgmt *Management) GetSystemByID(id int) (core.System, error) {
	system := core.System{}
	url := mgmt.rpc.buildServiceRegistryURL("/mgmt/systems/" + fmt.Sprintf("%d", id))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return system, err
	}

	req.Header.Set("Accept", "*/*")

	body, err := makeRequest(mgmt.rpc.http, req, 200, "Failed to get system")
	if err != nil {
		return system, err
	}

	system, err = parsers.UnmarshalGetSystemResponseJSON(string(body))
	if err != nil {
		return system, err
	}

	return system, nil
}

// TODO: Is this function tested?
func (mgmt *Management) GetSystemByName(systemName string) (core.System, error) {
	systems, err := mgmt.GetSystems()
	if err != nil {
		return core.System{}, err
	}
	for _, system := range systems {
		if system.SystemName == systemName {
			return system, nil
		}
	}

	return core.System{}, errors.New("System with name " + systemName + " not found")
}

func (mgmt *Management) RegisterService(system *core.System, httpMethod int, serviceDefinition, serviceURI string) (core.Service, error) {
	httpMethodStr := httpMethodToString(httpMethod)

	serviceRegReq := core.ServiceReqistrationRequest{}
	serviceRegReq.Metadata = map[string]string{
		"http-method": httpMethodStr,
	}
	serviceRegReq.Interfaces = []string{"HTTP-SECURE-JSON"}
	serviceRegReq.Secure = "TOKEN"
	serviceRegReq.ServiceDefinition = serviceDefinition
	serviceRegReq.ServiceUri = serviceURI

	serviceRegReq.Provider.Address = system.Address
	serviceRegReq.Provider.Port = system.Port
	serviceRegReq.Provider.SystemName = system.SystemName
	serviceRegReq.Provider.AuthenticationInfo = system.AuthenticationInfo

	jsonData, err := parsers.MarshalServiceReqistrationRequest(serviceRegReq)
	if err != nil {
		return core.Service{}, err
	}

	url := mgmt.rpc.buildServiceRegistryURL("/mgmt")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		log.Errorf("Error creating POST request: %v", err)
		return core.Service{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	jsonStr, err := makeRequest(mgmt.rpc.http, req, 201, "Failed to register service")
	if err != nil {
		return core.Service{}, err
	}

	return parsers.UnmarshalService(string(jsonStr))
}

// ARROWHEAD BUG: Why is the management API a REST API and client API a RPC API?
func (mgmt *Management) UnregisterService(systemID int) error {
	url := mgmt.rpc.buildServiceRegistryURL("/mgmt/" + fmt.Sprintf("%d", systemID))
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Errorf("Error creating POST request: %v", err)
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	_, err = makeRequest(mgmt.rpc.http, req, 200, "Failed to unregister service")
	if err != nil {
		return err
	}

	return nil
}

func (mgmt *Management) GetServices() ([]core.Service, error) {
	services := []core.Service{}
	url := mgmt.rpc.buildServiceRegistryURL("/mgmt?direction=ASC&sort_field=id")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return services, err
	}

	req.Header.Set("Accept", "*/*")

	body, err := makeRequest(mgmt.rpc.http, req, 200, "Failed to get services")
	if err != nil {
		return services, err
	}

	services, err = parsers.UnmarshalGetServicesResponseJSON(string(body))
	if err != nil {
		return services, err
	}

	return services, nil
}

// TODO: Is this function tested?
func (mgmt *Management) GetServiceDefinitionsIDsForProvider(providerID int, serviceDef string) ([]int, error) {
	services, err := mgmt.GetServices()
	if err != nil {
		return nil, err
	}

	serviceDefinitionIDs := []int{}
	for _, service := range services {
		if service.Provider.ID == providerID && service.ServiceDefinition.ServiceDefinition == serviceDef {
			serviceDefinitionIDs = append(serviceDefinitionIDs, service.ServiceDefinition.ID)
		}
	}

	return serviceDefinitionIDs, nil
}

// TODO: Is this function tested?
func (mgmt *Management) GetInterfaceIDsForProvider(providerID int) ([]int, error) {
	services, err := mgmt.GetServices()
	if err != nil {
		return nil, err
	}

	interfaceIDs := []int{}
	for _, service := range services {
		if service.Provider.ID == providerID {
			for _, interfaceID := range service.Interfaces {
				interfaceIDs = append(interfaceIDs, interfaceID.ID)
			}
			break
		}
	}
	return interfaceIDs, nil
}

func (mgmt *Management) GetServiceByID(id int) (core.Service, error) {
	service := core.Service{}
	url := mgmt.rpc.buildServiceRegistryURL("/mgmt/" + fmt.Sprintf("%d", id))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return service, err
	}

	req.Header.Set("Accept", "*/*")

	body, err := makeRequest(mgmt.rpc.http, req, 200, "Failed to get service")
	if err != nil {
		return service, err
	}

	service, err = parsers.UnmarshalGetServiceResponseJSON(string(body))
	if err != nil {
		return service, err
	}

	return service, nil
}

func (mgmt *Management) AddAuthorization(consumerName, providerName, serviceDef string) (core.Authorization, error) {
	addAuthReq := core.AddAuthorizationRequest{}

	consumer, err := mgmt.GetSystemByName(consumerName)
	if err != nil {
		return core.Authorization{}, err
	}

	provider, err := mgmt.GetSystemByName(providerName)
	if err != nil {
		return core.Authorization{}, err
	}

	serviceDefinitionIDs, err := mgmt.GetServiceDefinitionsIDsForProvider(provider.ID, serviceDef)
	if err != nil {
		return core.Authorization{}, err
	}

	interfaceIDs, err := mgmt.GetInterfaceIDsForProvider(provider.ID)
	if err != nil {
		return core.Authorization{}, err
	}

	addAuthReq.ConsumerID = consumer.ID
	addAuthReq.ProviderIDs = []int{provider.ID}
	addAuthReq.ServiceDefinitionIDs = serviceDefinitionIDs
	addAuthReq.InterfaceIDs = interfaceIDs

	jsonData, err := parsers.MarshalAddAuthorizationRequestJSON(addAuthReq)
	if err != nil {
		return core.Authorization{}, err
	}

	url := mgmt.rpc.buildAuthorizationURL("/mgmt/intracloud")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return core.Authorization{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	body, err := makeRequest(mgmt.rpc.http, req, 201, "Failed to add authorization rule")
	if err != nil {
		return core.Authorization{}, err
	}

	auths, err := parsers.UnmarshalGetAuthorizationsResponseJSON(string(body))
	if err != nil {
		return core.Authorization{}, err
	}

	if len(auths) == 0 {
		return core.Authorization{}, errors.New("Failed to add authorization rule")
	}

	if len(auths) > 1 {
		return core.Authorization{}, errors.New("More than one authorization rule added")
	}

	return auths[0], nil
}

func (mgmt *Management) GetAuthorizations() ([]core.Authorization, error) {
	auths := []core.Authorization{}
	url := mgmt.rpc.buildAuthorizationURL("/mgmt/intracloud?direction=ASC&sort_field=id")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return auths, err
	}

	req.Header.Set("Accept", "application/json")

	body, err := makeRequest(mgmt.rpc.http, req, 200, "Failed to get authorizations")
	if err != nil {
		return auths, err
	}

	auths, err = parsers.UnmarshalGetAuthorizationsResponseJSON(string(body))
	if err != nil {
		return auths, err
	}

	return auths, nil
}

func (mgmt *Management) RemoveAuthorization(authID int) error {
	url := mgmt.rpc.buildAuthorizationURL("/mgmt/intracloud/" + fmt.Sprintf("%d", authID))

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")

	_, err = makeRequest(mgmt.rpc.http, req, 200, "Failed to remove authorization rule")
	if err != nil {
		return err
	}

	return nil
}
