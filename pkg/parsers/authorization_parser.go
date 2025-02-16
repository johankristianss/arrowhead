package parsers

import (
	"encoding/json"

	"github.com/johankristianss/arrowhead/pkg/core"
)

func UnmarshalGetAuthorizationsResponseJSON(jsonStr string) ([]core.Authorization, error) {
	var resp core.AuthorizationsResponse
	err := json.Unmarshal([]byte(jsonStr), &resp)

	return resp.Authorizations, err
}

func MarshalAddAuthorizationRequestJSON(authAuthRequest core.AddAuthorizationRequest) (string, error) {
	jsonStr, err := json.Marshal(authAuthRequest)

	return string(jsonStr), err
}

func UnmarshallAuthorizationJSON(jsonStr string) (core.Authorization, error) {
	var auth core.Authorization
	err := json.Unmarshal([]byte(jsonStr), &auth)

	return auth, err
}
