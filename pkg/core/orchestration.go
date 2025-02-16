package core

type OrchestrationRequest struct {
	Commands           map[string]string   `json:"commands"`
	OrchestrationFlags OrchestrationFlags  `json:"orchestrationFlags"`
	PreferredProviders []PreferredProvider `json:"preferredProviders"`
	QOSRequirements    map[string]string   `json:"qosRequirements"`
	RequestedService   RequestedService    `json:"requestedService"`
	RequesterCloud     *Cloud              `json:"requesterCloud"`
	RequesterSystem    RequesterSystem     `json:"requesterSystem"`
}

type RequesterSystem struct {
	SystemName         string            `json:"systemName"`
	Address            string            `json:"address"`
	Port               int               `json:"port"`
	AuthenticationInfo string            `json:"authenticationInfo,omitempty"`
	Metadata           map[string]string `json:"metadata,omitempty"`
}

type OrchestrationFlags struct {
	OnlyPreferred          bool `json:"onlyPreferred"`
	OverrideStore          bool `json:"overrideStore"`
	ExternalServiceRequest bool `json:"externalServiceRequest"`
	EnableInterCloud       bool `json:"enableInterCloud"`
	EnableQoS              bool `json:"enableQoS"`
	Matchmaking            bool `json:"matchmaking"`
	MetadataSearch         bool `json:"metadataSearch"`
	TriggerInterCloud      bool `json:"triggerInterCloud"`
	PingProviders          bool `json:"pingProviders"`
}

type PreferredProvider struct {
	ProviderCloud  Cloud  `json:"providerCloud"`
	ProviderSystem System `json:"providerSystem"`
}

type Cloud struct {
	AuthenticationInfo string `json:"authenticationInfo"`
	GatekeeperRelayIds []int  `json:"gatekeeperRelayIds"`
	GatewayRelayIds    []int  `json:"gatewayRelayIds"`
	Name               string `json:"name"`
	Neighbor           bool   `json:"neighbor"`
	Operator           string `json:"operator"`
	Secure             bool   `json:"secure"`
}

type RequestedService struct {
	InterfaceRequirements []string          `json:"interfaceRequirements"`
	MaxVersionRequirement *int              `json:"maxVersionRequirement"`
	MetadataRequirements  map[string]string `json:"metadataRequirements"`
	MinVersionRequirement *int              `json:"minVersionRequirement"`
	PingProviders         bool              `json:"pingProviders"`
	//ProviderAddressTypeRequirements []string          `json:"providerAddressTypeRequirements"`
	SecurityRequirements         []string `json:"securityRequirements"`
	ServiceDefinitionRequirement string   `json:"serviceDefinitionRequirement"`
	VersionRequirement           *int     `json:"versionRequirement"`
}

type OrchestrationResponse struct {
	Response []MatchedService `json:"response"`
}

type MatchedService struct {
	Provider            Provider          `json:"provider"`
	ServiceDefinition   ServiceDefinition `json:"service"`
	ServiceURI          string            `json:"serviceUri"`
	Secure              string            `json:"secure"`
	Metadata            map[string]string `json:"metadata"`
	Interfaces          []Interface       `json:"interfaces"`
	Version             int               `json:"version"`
	AuthorizationTokens map[string]string `json:"authorizationTokens"`
	Warnings            []string          `json:"warnings"`
}
