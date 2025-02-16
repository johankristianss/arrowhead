package parsers

import (
	"testing"

	"github.com/johankristianss/arrowhead/pkg/core"
)

func TestParsersUnmarshalOrchestrationsResponse(t *testing.T) {
	jsonStr := `{"response":[{"provider":{"id":7,"systemName":"cardemoprovider","address":"127.0.0.1","port":8888,"authenticationInfo":"MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAw75K6GkGDIn/r1KH+ChGyP5iuo7BjK95fLWnagWDDNs+fVQoRYCHE97sY7DtaFYe1dbnw7Wwa1Rk3IEHCNGvjMEw1nvcq1EUJWS7C+fe9PS611Egracj/ium0wY+GJBUy7htilsrs1EOhlZxvfilSCm2+ucY3hUnWIs66OvR0LeAeBPjw+ytIPrwy+B1j3kScpzJ3l5fqWHp+B01K4AIYc10U9kTAFskd55aNmRBYcW87Zs/5CqJQzI+mrzl5EftOLejIs4sxhW4pEwwtnmcsRoYHj8Tbhl6Yi1AKoD1eAb1/LVBx+8e1YAtJ06ch9f+R4sXaN4hsMPrU7V/kg6orQIDAQAB","createdAt":"2025-02-13T21:31:17Z","updatedAt":"2025-02-13T21:31:17Z"},"service":{"id":36,"serviceDefinition":"create-car","createdAt":"2025-02-13T21:31:17Z","updatedAt":"2025-02-13T21:31:17Z"},"serviceUri":"/car","secure":"TOKEN","metadata":{"http-method":"POST"},"interfaces":[{"id":1,"interfaceName":"HTTP-SECURE-JSON","createdAt":"2023-06-08T12:01:21Z","updatedAt":"2023-06-08T12:01:21Z"}],"version":1,"authorizationTokens":{"HTTP-SECURE-JSON":"eyJhbGciOiJSU0EtT0FFUC0yNTYiLCJlbmMiOiJBMjU2Q0JDLUhTNTEyIiwiY3R5IjoiSldUIn0.AmvDyWqAgcMQZKauLu8Tw-4m0-K9ENPm-6dQk7hUYYgZitDnLss560L8lb_e5uR5msPi0inggu64Hw5VnadMy5975a8wWl2fUw1Y39EooeyOFnRrLz10-ZYFHWuDuZM1aFGr9kpKL9Crx0wzh-6R9oB8EA2mQDDlcCBPjtlSXqMHnN-ANyGayMST3oB3orzeGYntc3wA14PoN3-GPVRG6HoPAmJMElCk2hYEgY4HLrivjBHX0Gx9ysUbsd8p8KK43BI6j4BWx7theBGK2EJpcLCB1g-6xW3MoNjDtHkWtaYAs5mfXn7yCebSd5VfWshLn2tolYZp3Qzq0NTXCQxmog.rDe5GJJTG675GIZnwsAdpQ.9eJMJorSvm5JS0IyVjYbjqfeueznuwOUmeNiVaxiAD1TTkQSAO-tU85Ta4ZhmxaYKRXEESFUybB7qBg1A2TH1e-Lli1irHXAV5ioTLbfb2I3ArjFne0JD2iCzQPaPvS1BRXUctF7BhtuarMLYzckXzhrSWk9If0dZCHQoG7ngCin71cw_wRhDZzYg7gLYPLPTQCiIWeAZdz9KUzvReNYjvxOs8hlyRsOgiZFP5vXfdjjOF6_mR6zwZ_SBwQT14Ix_uPf2jmY3KveTpoLXwkGFVeMy5j6GngkEWTjL_HstkEiNSG0WY-OJi3YMNdmndavJ4CzzzHMcxvimodljn_qMnrl5QuyhuQ8VoXUqoyvqSYYLN2ih8KJGFuoDaBldhW2hocMKNOSPyZcl2IgBt3b66MpQjHD1Vk2mmGrOJE0CZ1tkYwfZaRPhVRC_cdRKW6qe0V5UF5fDamqpGdKzc1q_TUhrF7E_MR-GsdyP42oGB9XxB-xb0mc5sCO-9bd1SMNi1EhRfEzmkb4KLiB7BXYc_gtUT3M63MDM33e66T3OVDlXC0U8kUtQL3bX-D5AzR1XNz3qGB8kkuZqZTn12T0PHT1bK6xzQlLyUoWVqsYqIArZms2oMhp_MQldcMn8jtbvlspC4gf0lOrHNIEIrWiy6MSdU2vo4G_1sx7Y5RNOaKgRKsqPu4GW5CazsvgMxG2NaPouHAgKDSS9XUT7ksBwbh_CwmGgZdUhh-A3dSAqVbzH6QTL7CY1ljV-_7QplPhvowiZGoyEbqP3qYppD4S4Q.PDAKEIS2dl7zKsuwbT02vVVyoYBVuz7dKe271we1aVA"},"warnings":["TTL_UNKNOWN"]}]}`

	_, err := UnmarshalOrchestrationResponse([]byte(jsonStr))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	// TODO: Add more test
}

func TestParsersMarshalOrchestrationsRequest(t *testing.T) {
	orchestrationRequest := core.OrchestrationRequest{}
	orchestrationRequest.PreferredProviders = []core.PreferredProvider{}
	orchestrationRequest.Commands = make(map[string]string)
	orchestrationRequest.QOSRequirements = make(map[string]string)
	orchestrationRequest.RequesterSystem.SystemName = "cardemorconsumer"
	orchestrationRequest.RequesterSystem.AuthenticationInfo = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAws6zCv0nPnL8lAqSGBB8rMokDSrUEBTrMw1tb0IIsgkH1pbreBhkxyhPlDC7OstPiS6T2DZfo15Fz0ib8QmxJ81s4SJqfeh527DVrEgoDXKBemPzipfSQNBRP1jGDo7RDmNvNMtuQ50rWtMvMJJoo+jOxsS+a0pscr4Te5huaQQnLPDemjFA1YjE38BL3vEsdWt0mL4ueO819wj+BVpUcXQkNspDrRStzh3m8MAsdIzN/KPrPPsBlw8GlhCy+YAxSQ4+uA3M3wYN+FkZd77iNQZP4OgK9P3f/+6/ul2N9QCf0UvQy77JnE/DczXdKz+9y36wU88T8UEaHBaJp8sm6wIDAQAB"
	orchestrationRequest.RequesterSystem.Address = "localhost"
	orchestrationRequest.RequesterSystem.Port = 8091
	orchestrationRequest.RequestedService.ServiceDefinitionRequirement = "testservice"
	orchestrationRequest.RequestedService.InterfaceRequirements = []string{"HTTP-SECURE-JSON"}
	orchestrationRequest.OrchestrationFlags.OnlyPreferred = false
	orchestrationRequest.OrchestrationFlags.OverrideStore = true
	orchestrationRequest.OrchestrationFlags.ExternalServiceRequest = false
	orchestrationRequest.OrchestrationFlags.EnableInterCloud = false
	orchestrationRequest.OrchestrationFlags.EnableQoS = false
	orchestrationRequest.OrchestrationFlags.Matchmaking = true
	orchestrationRequest.OrchestrationFlags.MetadataSearch = false
	orchestrationRequest.OrchestrationFlags.TriggerInterCloud = false
	orchestrationRequest.OrchestrationFlags.PingProviders = false

	//json, err := MarshalOrchestrationRequest(orchestrationRequest)
	_, err := MarshalOrchestrationRequest(orchestrationRequest)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	// TODO Add more tests
}
