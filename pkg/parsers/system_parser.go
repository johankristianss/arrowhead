package parsers

import (
	"encoding/json"

	"github.com/johankristianss/arrowhead/pkg/core"
)

func UnmarshalGetSystemsResponseJSON(jsonStr string) ([]core.System, error) {
	var resp core.SystemsResponse
	err := json.Unmarshal([]byte(jsonStr), &resp)

	return resp.Systems, err
}

func UnmarshalGetSystemResponseJSON(jsonStr string) (core.System, error) {
	var system core.System
	err := json.Unmarshal([]byte(jsonStr), &system)

	return system, err
}

func UnmarshalSystemJSON(jsonStr string) (core.System, error) {
	var system core.System
	err := json.Unmarshal([]byte(jsonStr), &system)

	return system, err
}

func MarshalSystemRegistration(systemRegistration core.SystemRegistration) (string, error) {
	jsonBytes, err := json.Marshal(systemRegistration)
	return string(jsonBytes), err
}
