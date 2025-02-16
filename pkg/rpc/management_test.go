package rpc

import (
	"os"
	"testing"

	"github.com/johankristianss/arrowhead/pkg/core"
	"github.com/johankristianss/arrowhead/pkg/security/openssl"
	"github.com/stretchr/testify/assert"
)

func TestRPCMgmtGetSystems(t *testing.T) {
	arrowhead, err := CreateArrowhead(testConfig)
	assert.Nil(t, err)

	systems, err := arrowhead.Management.GetSystems()
	assert.Nil(t, err)
	assert.True(t, len(systems) > 0)
}

func TestRPCMgmtGetSystemByID(t *testing.T) {
	arrowhead, err := CreateArrowhead(testConfig)
	assert.Nil(t, err)

	systems, err := arrowhead.Management.GetSystems()
	assert.Nil(t, err)
	assert.True(t, len(systems) > 0)

	system, err := arrowhead.Management.GetSystemByID(systems[0].ID)
	assert.Nil(t, err)
	assert.True(t, len(system.SystemName) > 0)
}

func TestRPCMgmtGetServices(t *testing.T) {
	arrowhead, err := CreateArrowhead(testConfig)
	assert.Nil(t, err)

	services, err := arrowhead.Management.GetServices()
	assert.Nil(t, err)
	assert.True(t, len(services) > 0)
}

func TestRPCMgmtGetServiceByID(t *testing.T) {
	arrowhead, err := CreateArrowhead(testConfig)
	assert.Nil(t, err)

	services, err := arrowhead.Management.GetServices()
	assert.Nil(t, err)
	assert.True(t, len(services) > 0)

	service, err := arrowhead.Management.GetServiceByID(services[0].ID)
	assert.Nil(t, err)
	assert.True(t, len(service.Provider.SystemName) > 0)
}

func TestRPCMgmtGetAuthorizations(t *testing.T) {
	arrowhead, err := CreateArrowhead(testConfig)
	assert.Nil(t, err)

	//auths, err := arrowhead.Management.GetAuthorizations()
	_, err = arrowhead.Management.GetAuthorizations()
	assert.Nil(t, err)
	//assert.True(t, len(auths) > 0)
}

func TestRPCMgmtAddRemoveSystem(t *testing.T) {
	arrowhead, err := CreateArrowhead(testConfig)
	assert.Nil(t, err)

	san := "DNS:testconsumer,DNS:testconsumer-ip,DNS:localhost,IP:127.0.0.1"
	if extra := os.Getenv("SUBJECT_ALTERNATIVE_NAME"); extra != "" {
		san += "," + extra
	}

	certManager := openssl.CreateOpenSSLCertManager()
	err = certManager.CreateSystemKeystore(
		"../security/testcerts/master.p12", "arrowhead.eu",
		"../security/testcerts/c1.p12", "c1.ltu.arrowhead.eu",
		"./testconsumer.p12", "testconsumer.c1.ltu.arrowhead.eu",
		"testconsumer", san, "123456",
	)

	// Make a deep copy of the testConfig
	systemConfig := *testConfig
	systemConfig.KeystorePath = "./testconsumer.p12"

	defer os.Remove("./testconsumer.p12")
	defer os.Remove("./testconsumer.pub")

	authInfo, err := certManager.GetPublicKey("./testconsumer.p12", "123456")
	assert.Nil(t, err)
	systemReg := core.SystemRegistration{
		Address:            "localhost",
		AuthenticationInfo: authInfo,
		Metadata:           map[string]string{},
		Port:               8091,
		SystemName:         "testconsumer",
	}

	system, err := arrowhead.Management.RegisterSystem(systemReg)
	assert.Nil(t, err)
	assert.True(t, len(system.SystemName) > 0)

	systems, err := arrowhead.Management.GetSystems()
	assert.Nil(t, err)
	found := false
	var foundSystem *core.System
	for _, system := range systems {
		if system.SystemName == "testconsumer" {
			found = true
			foundSystem = &system
			break
		}
	}
	assert.True(t, found)

	err = arrowhead.Management.UnregisterSystem(foundSystem)
	assert.Nil(t, err)
}

func TestRPCMgmtAddRemoveService(t *testing.T) {
	arrowhead, err := CreateArrowhead(testConfig)
	assert.Nil(t, err)

	san := "DNS:testconsumer,DNS:testconsumer-ip,DNS:localhost,IP:127.0.0.1"
	if extra := os.Getenv("SUBJECT_ALTERNATIVE_NAME"); extra != "" {
		san += "," + extra
	}

	certManager := openssl.CreateOpenSSLCertManager()
	err = certManager.CreateSystemKeystore(
		"../security/testcerts/master.p12", "arrowhead.eu",
		"../security/testcerts/c1.p12", "c1.ltu.arrowhead.eu",
		"./testconsumer.p12", "testconsumer.c1.ltu.arrowhead.eu",
		"testconsumer", san, "123456",
	)

	defer os.Remove("./testconsumer.p12")
	defer os.Remove("./testconsumer.pub")

	authInfo, err := certManager.GetPublicKey("./testconsumer.p12", "123456")
	assert.Nil(t, err)
	systemReg := core.SystemRegistration{
		Address:            "localhost",
		AuthenticationInfo: authInfo,
		Metadata:           map[string]string{},
		Port:               8091,
		SystemName:         "testconsumer",
	}

	system, err := arrowhead.Management.RegisterSystem(systemReg)
	assert.Nil(t, err)
	assert.True(t, len(system.SystemName) > 0)

	systems, err := arrowhead.Management.GetSystems()
	assert.Nil(t, err)
	found := false
	var foundSystem *core.System
	for _, system := range systems {
		if system.SystemName == "testconsumer" {
			found = true
			foundSystem = &system
			break
		}
	}
	assert.True(t, found)

	// Make a deep copy of the testConfig
	systemConfig := *testConfig
	systemConfig.KeystorePath = "./testconsumer.p12"

	service, err := arrowhead.Management.RegisterService(foundSystem, POST, "testservice", "/dummy")
	assert.Nil(t, err)

	err = arrowhead.Management.UnregisterService(service.ID)
	assert.Nil(t, err)

	err = arrowhead.Management.UnregisterSystem(foundSystem)
	assert.Nil(t, err)
}

func TestRPCMgmtAddRemoveAuthorization(t *testing.T) {
	arrowhead, err := CreateArrowhead(testConfig)
	assert.Nil(t, err)

	// Add a provder system
	san := "DNS:testprovider,DNS:testconsumer-ip,DNS:localhost,IP:127.0.0.1"
	if extra := os.Getenv("SUBJECT_ALTERNATIVE_NAME"); extra != "" {
		san += "," + extra
	}

	certManager := openssl.CreateOpenSSLCertManager()
	err = certManager.CreateSystemKeystore(
		"../security/testcerts/master.p12", "arrowhead.eu",
		"../security/testcerts/c1.p12", "c1.ltu.arrowhead.eu",
		"./testprovider.p12", "testprovider.c1.ltu.arrowhead.eu",
		"testprovider", san, "123456",
	)

	defer os.Remove("./testprovider.p12")
	defer os.Remove("./testprovider.pub")

	authInfo, err := certManager.GetPublicKey("./testprovider.p12", "123456")
	assert.Nil(t, err)
	systemReg := core.SystemRegistration{
		Address:            "localhost",
		AuthenticationInfo: authInfo,
		Metadata:           map[string]string{},
		Port:               8091,
		SystemName:         "testprovider",
	}

	provider, err := arrowhead.Management.RegisterSystem(systemReg)
	assert.Nil(t, err)
	defer arrowhead.Management.UnregisterSystem(&provider)

	// Add a service
	service, err := arrowhead.Management.RegisterService(&provider, POST, "testservice", "/dummy")
	assert.Nil(t, err)
	defer arrowhead.Management.UnregisterService(service.ID)

	// Add a consumer system
	san = "DNS:testconsumer,DNS:testconsumer-ip,DNS:localhost,IP:127.0.0.1"
	if extra := os.Getenv("SUBJECT_ALTERNATIVE_NAME"); extra != "" {
		san += "," + extra
	}

	err = certManager.CreateSystemKeystore(
		"../security/testcerts/master.p12", "arrowhead.eu",
		"../security/testcerts/c1.p12", "c1.ltu.arrowhead.eu",
		"./testconsumer.p12", "testconsumer.c1.ltu.arrowhead.eu",
		"testconsumer", san, "123456",
	)

	defer os.Remove("./testconsumer.p12")
	defer os.Remove("./testconsumer.pub")

	authInfo, err = certManager.GetPublicKey("./testconsumer.p12", "123456")
	assert.Nil(t, err)
	systemReg = core.SystemRegistration{
		Address:            "localhost",
		AuthenticationInfo: authInfo,
		Metadata:           map[string]string{},
		Port:               8091,
		SystemName:         "testconsumer",
	}

	consumer, err := arrowhead.Management.RegisterSystem(systemReg)
	assert.Nil(t, err)
	defer arrowhead.Management.UnregisterSystem(&consumer)

	// Add an authorization
	auth, err := arrowhead.Management.AddAuthorization(consumer.SystemName, provider.SystemName, "testservice")
	assert.Nil(t, err)
	defer arrowhead.Management.RemoveAuthorization(auth.ID)
}
