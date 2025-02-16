package cli

import (
	"os"
	"strconv"

	"github.com/johankristianss/arrowhead/pkg/build"
	"github.com/johankristianss/arrowhead/pkg/rpc"
	log "github.com/sirupsen/logrus"
)

func checkEnv() {
	envVars := []string{
		"ARROWHEAD_AUTHORIZATION_HOST",
		"ARROWHEAD_AUTHORIZATION_PORT",
		"ARROWHEAD_SERVICEREGISTRY_HOST",
		"ARROWHEAD_SERVICEREGISTRY_PORT",
		"ARROWHEAD_ORCHESTRATOR_HOST",
		"ARROWHEAD_ORCHESTRATOR_PORT",
		"ARROWHEAD_KEYSTORE_PASSWORD",
		"ARROWHEAD_ROOT_KEYSTORE",
		"ARROWHEAD_ROOT_KEYSTORE_ALIAS",
		"ARROWHEAD_CLOUD_KEYSTORE",
		"ARROWHEAD_CLOUD_KEYSTORE_ALIAS",
		"ARROWHEAD_SYSOPS_KEYSTORE",
		"ARROWHEAD_TRUSTSTORE",
		"ARROWHEAD_TLS",
		"ARROWHEAD_ASCII",
		"ARROWHEAD_VERBOSE",
	}

	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			log.WithFields(log.Fields{"Error": envVar + " is not set", "BuildVersion": build.BuildVersion, "BuildTime": build.BuildTime}).Error("Environment variable not set")
		}
	}
}

func parseEnv() {
	ASCIIEnv := os.Getenv("ARROWHEAD_ASCII")
	if ASCIIEnv == "true" {
		ASCII = true
	} else if ASCIIEnv == "false" {
		ASCII = false
	}

	VerboseEnv := os.Getenv("ARROWHEAD_VERBOSE")
	if VerboseEnv == "true" {
		Verbose = true
	}

	if Verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	KeystorePassword = os.Getenv("ARROWHEAD_KEYSTORE_PASSWORD")
	RootKeystorePath = os.Getenv("ARROWHEAD_ROOT_KEYSTORE")
	RootKeystoreAlias = os.Getenv("ARROWHEAD_ROOT_KEYSTORE_ALIAS")
	CloudKeystorePath = os.Getenv("ARROWHEAD_CLOUD_KEYSTORE")
	CloudKeystoreAlias = os.Getenv("ARROWHEAD_CLOUD_KEYSTORE_ALIAS")
	SysOpsKeystorePath = os.Getenv("ARROWHEAD_SYSOPS_KEYSTORE")
	Truststore = os.Getenv("ARROWHEAD_TRUSTSTORE")

	ArrowheadTLSEnv := os.Getenv("ARROWHEAD_TLS")
	if ArrowheadTLSEnv == "true" {
		ArrowheadTLS = true
	} else if ArrowheadTLSEnv == "false" {
		ArrowheadTLS = false
	}

	var err error
	ArrowheadAuthorizationHost = os.Getenv("ARROWHEAD_AUTHORIZATION_HOST")
	ArrowheadAuthorizationPortStr := os.Getenv("ARROWHEAD_AUTHORIZATION_PORT")
	ArrowheadAuthorizationPort, err = strconv.Atoi(ArrowheadAuthorizationPortStr)
	if err != nil {
		log.WithFields(log.Fields{"Error": err, "BuildVersion": build.BuildVersion, "BuildTime": build.BuildTime}).Error("Failed to parse ARROWHEAD_AUTHORIZATION_PORT")
	}

	ArrowheadServiceRegistryHost = os.Getenv("ARROWHEAD_SERVICEREGISTRY_HOST")

	ArrowheadServiceRegistryPortStr := os.Getenv("ARROWHEAD_SERVICEREGISTRY_PORT")
	ArrowheadServiceRegistryPort, err = strconv.Atoi(ArrowheadServiceRegistryPortStr)
	if err != nil {
		log.WithFields(log.Fields{"Error": err, "BuildVersion": build.BuildVersion, "BuildTime": build.BuildTime}).Error("Failed to parse ARROWHEAD_SERVICEREGISTRY_PORT")
	}

	ArrowheadOrchestratorHost = os.Getenv("ARROWHEAD_ORCHESTRATOR_HOST")

	ArrowheadOrchestratorPortStr := os.Getenv("ARROWHEAD_ORCHESTRATOR_PORT")
	ArrowheadOrchestratorPort, err = strconv.Atoi(ArrowheadOrchestratorPortStr)
	if err != nil {
		log.WithFields(log.Fields{"Error": err, "BuildVersion": build.BuildVersion, "BuildTime": build.BuildTime}).Error("Failed to parse ARROWHEAD_ORCHESTRATOR_PORT")
	}

	checkEnv()
}

func generateRPCSysOpsConfig() *rpc.Config {
	return &rpc.Config{
		TLS:                 ArrowheadTLS,
		AuthorizationHost:   ArrowheadAuthorizationHost,
		AuthorizationPort:   ArrowheadAuthorizationPort,
		ServiceRegistryHost: ArrowheadServiceRegistryHost,
		ServiceRegistryPort: ArrowheadServiceRegistryPort,
		OrchestratorHost:    ArrowheadOrchestratorHost,
		OrchestratorPort:    ArrowheadOrchestratorPort,
		KeystorePath:        SysOpsKeystorePath,
		Password:            KeystorePassword,
		TruststorePath:      Truststore,
	}
}
