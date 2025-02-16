package rpc

import (
	"github.com/johankristianss/arrowhead/pkg/core"
)

func BuildOrchestrationRequest(requesterSystemName, requesterAddress string, requesterPort int, serviceName string) core.OrchestrationRequest {
	orchestrationRequest := core.OrchestrationRequest{}
	orchestrationRequest.PreferredProviders = []core.PreferredProvider{}
	orchestrationRequest.Commands = make(map[string]string)
	orchestrationRequest.QOSRequirements = make(map[string]string)
	orchestrationRequest.RequesterSystem.SystemName = requesterSystemName
	orchestrationRequest.RequesterSystem.Address = requesterAddress
	orchestrationRequest.RequesterSystem.Port = requesterPort
	orchestrationRequest.RequestedService.MetadataRequirements = make(map[string]string)
	orchestrationRequest.RequestedService.ServiceDefinitionRequirement = serviceName
	orchestrationRequest.RequestedService.InterfaceRequirements = []string{"HTTP-SECURE-JSON"}
	orchestrationRequest.RequestedService.MetadataRequirements = nil
	orchestrationRequest.OrchestrationFlags.OnlyPreferred = false
	orchestrationRequest.OrchestrationFlags.OverrideStore = true
	orchestrationRequest.OrchestrationFlags.ExternalServiceRequest = false
	orchestrationRequest.OrchestrationFlags.EnableInterCloud = false
	orchestrationRequest.OrchestrationFlags.EnableQoS = false
	orchestrationRequest.OrchestrationFlags.Matchmaking = true
	orchestrationRequest.OrchestrationFlags.MetadataSearch = false
	orchestrationRequest.OrchestrationFlags.TriggerInterCloud = false
	orchestrationRequest.OrchestrationFlags.PingProviders = false

	return orchestrationRequest
}
