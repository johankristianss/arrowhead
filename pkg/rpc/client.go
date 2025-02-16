package rpc

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/johankristianss/arrowhead/pkg/core"
	"github.com/johankristianss/arrowhead/pkg/parsers"
	log "github.com/sirupsen/logrus"
)

func (client *Client) RegisterSystem(systemReg core.SystemRegistration) error {
	jsonData, err := parsers.MarshalSystemRegistration(systemReg)
	if err != nil {
		return err
	}

	url := client.rpc.buildServiceRegistryURL("/register-system")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		log.Fatalf("Error creating POST request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	_, err = makeRequest(client.rpc.http, req, 201, "Failed to register system")
	return err
}

func (client *Client) UnregisterSystem(system *core.System) error {
	portStr := strconv.Itoa(system.Port)

	url := client.rpc.buildServiceRegistryURL("/unregister-system?address=" + system.Address + "&port=" + portStr + "&system_name=" + system.SystemName)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatalf("Error creating DELETE request: %v", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	_, err = makeRequest(client.rpc.http, req, 200, "Failed to unregister system")
	return err
}

func (client *Client) RegisterService(system *core.System, httpMethod int, serviceDefinition, serviceURI string) (core.Service, error) {
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

	url := client.rpc.buildServiceRegistryURL("/register")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		log.Errorf("Error creating POST request: %v", err)
		return core.Service{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	jsonStr, err := makeRequest(client.rpc.http, req, 201, "Failed to register service")
	if err != nil {
		return core.Service{}, err
	}

	return parsers.UnmarshalService(string(jsonStr))
}

func (client *Client) UnregisterService(systemName, serviceURI, serviceDefinition, address string, port int) error {
	// ARROWHEAD BUG: OpenAPI/Swagger documentation states that address is optional, but it is not
	// ARROWHEAD BUG: REST API is mixed up, the API should probably be DELETE /service/{serviceID} instead of DELETE /unregister, otherwise it should probably be a GET request
	url := client.rpc.buildServiceRegistryURL("/unregister?system_name=" + systemName + "&service_uri=" + serviceURI + "&service_definition=" + serviceDefinition + "&address=" + address + "&port=" + strconv.Itoa(port))
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Errorf("Error creating POST request: %v", err)
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	_, err = makeRequest(client.rpc.http, req, 200, "Failed to unregister service")
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) Orchestrate(orchestrationReq core.OrchestrationRequest) (core.OrchestrationResponse, error) {
	jsonData, err := parsers.MarshalOrchestrationRequest(orchestrationReq)
	if err != nil {
		return core.OrchestrationResponse{}, err
	}
	url := client.rpc.buildOrchestratorURL("/orchestration")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		log.Errorf("Error creating POST request: %v", err)
		return core.OrchestrationResponse{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	jsonStr, err := makeRequest(client.rpc.http, req, 200, "Failed to orchestrate")
	if err != nil {
		return core.OrchestrationResponse{}, err
	}

	return parsers.UnmarshalOrchestrationResponse(jsonStr)
}
