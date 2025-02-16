package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsersUnmarshalGetAuthorizationsResponseJSON(t *testing.T) {
	jsonStr := `
{
  "data": [
    {
      "id": 1,
      "consumerSystem": {
        "id": 8,
        "systemName": "cardemoconsumer",
        "address": "localhost",
        "port": 8080,
        "createdAt": "2025-02-09T09:16:11Z",
        "updatedAt": "2025-02-09T09:16:11Z"
      },
      "providerSystem": {
        "id": 7,
        "systemName": "cardemoprovider",
        "address": "127.0.0.1",
        "port": 8888,
        "authenticationInfo": "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAw75K6GkGDIn/r1KH+ChGyP5iuo7BjK95fLWnagWDDNs+fVQoRYCHE97sY7DtaFYe1dbnw7Wwa1Rk3IEHCNGvjMEw1nvcq1EUJWS7C+fe9PS611Egracj/ium0wY+GJBUy7htilsrs1EOhlZxvfilSCm2+ucY3hUnWIs66OvR0LeAeBPjw+ytIPrwy+B1j3kScpzJ3l5fqWHp+B01K4AIYc10U9kTAFskd55aNmRBYcW87Zs/5CqJQzI+mrzl5EftOLejIs4sxhW4pEwwtnmcsRoYHj8Tbhl6Yi1AKoD1eAb1/LVBx+8e1YAtJ06ch9f+R4sXaN4hsMPrU7V/kg6orQIDAQAB",
        "createdAt": "2025-02-09T09:14:56Z",
        "updatedAt": "2025-02-09T09:14:56Z"
      },
      "serviceDefinition": {
        "id": 36,
        "serviceDefinition": "create-car",
        "createdAt": "2025-02-09T09:14:56Z",
        "updatedAt": "2025-02-09T09:14:56Z"
      },
      "interfaces": [
        {
          "id": 1,
          "interfaceName": "HTTP-SECURE-JSON",
          "createdAt": "2023-06-08T12:01:21Z",
          "updatedAt": "2023-06-08T12:01:21Z"
        }
      ],
      "createdAt": "2025-02-09T09:24:42Z",
      "updatedAt": "2025-02-09T09:24:42Z"
    },
    {
      "id": 2,
      "consumerSystem": {
        "id": 8,
        "systemName": "cardemoconsumer",
        "address": "localhost",
        "port": 8080,
        "createdAt": "2025-02-09T09:16:11Z",
        "updatedAt": "2025-02-09T09:16:11Z"
      },
      "providerSystem": {
        "id": 7,
        "systemName": "cardemoprovider",
        "address": "127.0.0.1",
        "port": 8888,
        "authenticationInfo": "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAw75K6GkGDIn/r1KH+ChGyP5iuo7BjK95fLWnagWDDNs+fVQoRYCHE97sY7DtaFYe1dbnw7Wwa1Rk3IEHCNGvjMEw1nvcq1EUJWS7C+fe9PS611Egracj/ium0wY+GJBUy7htilsrs1EOhlZxvfilSCm2+ucY3hUnWIs66OvR0LeAeBPjw+ytIPrwy+B1j3kScpzJ3l5fqWHp+B01K4AIYc10U9kTAFskd55aNmRBYcW87Zs/5CqJQzI+mrzl5EftOLejIs4sxhW4pEwwtnmcsRoYHj8Tbhl6Yi1AKoD1eAb1/LVBx+8e1YAtJ06ch9f+R4sXaN4hsMPrU7V/kg6orQIDAQAB",
        "createdAt": "2025-02-09T09:14:56Z",
        "updatedAt": "2025-02-09T09:14:56Z"
      },
      "serviceDefinition": {
        "id": 37,
        "serviceDefinition": "get-car",
        "createdAt": "2025-02-09T09:14:56Z",
        "updatedAt": "2025-02-09T09:14:56Z"
      },
      "interfaces": [
        {
          "id": 1,
          "interfaceName": "HTTP-SECURE-JSON",
          "createdAt": "2023-06-08T12:01:21Z",
          "updatedAt": "2023-06-08T12:01:21Z"
        }
      ],
      "createdAt": "2025-02-09T09:24:42Z",
      "updatedAt": "2025-02-09T09:24:42Z"
    }
  ],
  "count": 2
}	
	`
	auths, err := UnmarshalGetAuthorizationsResponseJSON(jsonStr)
	assert.Nil(t, err)

	assert.True(t, len(auths) > 0)
	// TODO Add more tests
}
