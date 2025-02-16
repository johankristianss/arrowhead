package rpc

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/johankristianss/arrowhead/pkg/parsers"
	log "github.com/sirupsen/logrus"
)

func (rpc *RPC) buildServiceRegistryURL(path string) string {
	protocol := "http"
	if rpc.Config.TLS {
		protocol = "https"
	}
	return fmt.Sprintf("%s://%s:%d/serviceregistry%s", protocol, rpc.Config.ServiceRegistryHost, rpc.Config.ServiceRegistryPort, path)
}

func (rpc *RPC) buildAuthorizationURL(path string) string {
	protocol := "http"
	if rpc.Config.TLS {
		protocol = "https"
	}
	return fmt.Sprintf("%s://%s:%d/authorization%s", protocol, rpc.Config.AuthorizationHost, rpc.Config.AuthorizationPort, path)
}

func (rpc *RPC) buildOrchestratorURL(path string) string {
	protocol := "http"
	if rpc.Config.TLS {
		protocol = "https"
	}
	return fmt.Sprintf("%s://%s:%d/orchestrator%s", protocol, rpc.Config.OrchestratorHost, rpc.Config.OrchestratorPort, path)
}

func makeRequest(http *http.Client, req *http.Request, expectedStatusCode int, errMsg string) ([]byte, error) {
	resp, err := http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = checkRPCError(expectedStatusCode, resp, body, errMsg)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func checkRPCError(expectedStatusCode int, resp *http.Response, body []byte, errMsg string) error {
	if resp.StatusCode != expectedStatusCode {
		errorResp := parsers.UnmarshalErrorResponse(body)
		log.WithFields(log.Fields{
			"StatusCode":    resp.StatusCode,
			"ErrorMessage":  errorResp.ErrorMessage,
			"ErrorCode":     errorResp.ErrorCode,
			"ExceptionType": errorResp.ExceptionType}).Error(errMsg)
		return errors.New(errMsg + ":" + errorResp.ErrorMessage)
	}
	return nil
}
