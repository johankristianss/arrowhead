package rpc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRPCBuildServiceRegistryURL(t *testing.T) {
	arrowhead, err := CreateArrowhead(testConfig)
	assert.Nil(t, err)
	path := "/path"
	expected := "https://localhost:8443/serviceregistry/path"
	if url := arrowhead.buildServiceRegistryURL(path); url != expected {
		t.Errorf("Expected %s, got %s", expected, url)
	}
}

func TestRPCBuildAuthorizationURL(t *testing.T) {
	arrowhead, err := CreateArrowhead(testConfig)
	assert.Nil(t, err)
	path := "/path"
	expected := "https://localhost:8445/authorization/path"
	if url := arrowhead.buildAuthorizationURL(path); url != expected {
		t.Errorf("Expected %s, got %s", expected, url)
	}
}

func TestRPCBuildOrchestratorURL(t *testing.T) {
	arrowhead, err := CreateArrowhead(testConfig)
	assert.Nil(t, err)
	path := "/path"
	expected := "https://localhost:8441/orchestrator/path"
	if url := arrowhead.buildOrchestratorURL(path); url != expected {
		t.Errorf("Expected %s, got %s", expected, url)
	}
}
