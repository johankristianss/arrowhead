package parsers

import (
	"encoding/json"

	"github.com/johankristianss/arrowhead/pkg/core"
)

func UnmarshalGetServicesResponseJSON(jsonStr string) ([]core.Service, error) {
	var resp core.ServicesResponse
	err := json.Unmarshal([]byte(jsonStr), &resp)

	return resp.Services, err
}

func UnmarshalGetServiceResponseJSON(jsonStr string) (core.Service, error) {
	var service core.Service
	err := json.Unmarshal([]byte(jsonStr), &service)

	return service, err
}

func MarshalService(service core.Service) (string, error) {
	jsonStr, err := json.Marshal(service)

	return string(jsonStr), err
}

func UnmarshalService(serviceStr string) (core.Service, error) {
	var service core.Service
	err := json.Unmarshal([]byte(serviceStr), &service)

	return service, err
}

func MarshalServiceReqistrationRequest(serviceRegReq core.ServiceReqistrationRequest) (string, error) {
	jsonStr, err := json.Marshal(serviceRegReq)

	return string(jsonStr), err
}
