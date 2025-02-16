package core

import "time"

type AddAuthorizationRequest struct {
	ConsumerID           int   `json:"consumerId"`
	ProviderIDs          []int `json:"providerIds"`
	InterfaceIDs         []int `json:"interfaceIds"`
	ServiceDefinitionIDs []int `json:"serviceDefinitionIds"`
}

type AuthorizationsResponse struct {
	Authorizations []Authorization `json:"data"`
	Count          int             `json:"count"`
}

type Authorization struct {
	ID                int               `json:"id"`
	ConsumerSystem    System            `json:"consumerSystem"`
	ProviderSystem    Provider          `json:"providerSystem"`
	ServiceDefinition ServiceDefinition `json:"serviceDefinition"`
	Interfaces        []Interface       `json:"interfaces"`
	CreatedAt         time.Time         `json:"createdAt"`
	UpdatedAt         time.Time         `json:"updatedAt"`
}
