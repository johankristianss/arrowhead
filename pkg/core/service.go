package core

import (
	"time"
)

// ARROWHEAD BUG: Inconsitent field names, e.g. sometimes Version is an int, sometimes a string
// Sometimes ServiceDefinition is a string, sometimes a struct, which makes it very
// confusing to work with the API.

type ServiceReqistrationRequest struct {
	EndOfValidity string            `json:"endOfValidity"`
	Interfaces    []string          `json:"interfaces"`
	Metadata      map[string]string `json:"metadata"`
	// ARROWHEAD BUG:  Provider cannot be a Provider struct, because the providerSystem field some
	// fields are different from the Provider struct, e.g. ID and created/updated at fields.
	Provider          ProviderSystem `json:"providerSystem"`
	Secure            string         `json:"secure"`
	ServiceDefinition string         `json:"serviceDefinition"`
	ServiceUri        string         `json:"serviceUri"`
	Version           string         `json:"version"`
}

type ProviderSystem struct {
	SystemName         string            `json:"systemName"`
	Address            string            `json:"address"`
	Port               int               `json:"port"`
	AuthenticationInfo string            `json:"authenticationInfo"`
	Metadata           map[string]string `json:"metadata,omitempty"`
}

type ServicesResponse struct {
	Services []Service `json:"data"`
	Count    int       `json:"count"`
}

type Service struct {
	ID                int               `json:"id"`
	ServiceDefinition ServiceDefinition `json:"serviceDefinition"`
	Provider          Provider          `json:"provider"`
	ServiceURI        string            `json:"serviceUri"`
	Secure            string            `json:"secure"`
	Version           int               `json:"version"`
	Interfaces        []Interface       `json:"interfaces"`
	Metadata          map[string]string `json:"metadata,omitempty"`
	CreatedAt         time.Time         `json:"createdAt"`
	UpdatedAt         time.Time         `json:"updatedAt"`
	EndOfValidity     time.Time         `json:"endOfValidity"`
}

type ServiceDefinition struct {
	ID                int       `json:"id"`
	ServiceDefinition string    `json:"serviceDefinition"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type Provider struct {
	ID                 int               `json:"id"`
	SystemName         string            `json:"systemName"`
	Address            string            `json:"address"`
	Port               int               `json:"port"`
	AuthenticationInfo string            `json:"authenticationInfo"`
	Metadata           map[string]string `json:"metadata,omitempty"`
	CreatedAt          time.Time         `json:"createdAt"`
	UpdatedAt          time.Time         `json:"updatedAt"`
}

type Interface struct {
	ID            int       `json:"id"`
	InterfaceName string    `json:"interfaceName"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
