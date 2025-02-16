package rpc

import (
	"os"
	"testing"
)

var testConfig *Config

func TestMain(m *testing.M) {
	testConfig = &Config{
		TLS:                 true,
		AuthorizationHost:   "localhost",
		AuthorizationPort:   8445,
		ServiceRegistryHost: "localhost",
		ServiceRegistryPort: 8443,
		OrchestratorHost:    "localhost",
		OrchestratorPort:    8441,
		KeystorePath:        "../security/testcerts/sysop.p12",
		TruststorePath:      "../security/testcerts/truststore.pem",
		Password:            "123456",
	}

	code := m.Run()
	os.Exit(code)
}
