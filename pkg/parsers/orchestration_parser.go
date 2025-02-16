package parsers

import (
	"encoding/json"

	"github.com/johankristianss/arrowhead/pkg/core"
)

func MarshalOrchestrationRequest(orchestrationRequest core.OrchestrationRequest) ([]byte, error) {
	// make indented
	return json.MarshalIndent(orchestrationRequest, "", "  ")
	//return json.Marshal(orchestrationRequest)
}

func UnmarshalOrchestrationResponse(data []byte) (core.OrchestrationResponse, error) {
	var orchestrationResponse core.OrchestrationResponse
	err := json.Unmarshal(data, &orchestrationResponse)
	return orchestrationResponse, err
}
