package cli

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/johankristianss/arrowhead/pkg/build"
	"github.com/johankristianss/arrowhead/pkg/core"
	"github.com/johankristianss/arrowhead/pkg/security"
	log "github.com/sirupsen/logrus"
)

func truncateString(s string) string {
	if len(s) > MAX_VALUE_LENGTH {
		return s[:MAX_VALUE_LENGTH] + "..."
	}
	return s
}

func filterSystemByName(systems []core.System, name string) []core.System {
	var result []core.System
	for _, system := range systems {
		if strings.Contains(system.SystemName, name) {
			result = append(result, system)
		}
	}
	return result
}

func filterServiceByName(services []core.Service, name string) []core.Service {
	var result []core.Service
	for _, service := range services {
		if strings.Contains(service.Provider.SystemName, name) {
			result = append(result, service)
		}
	}
	return result
}

func checkIfDirExists(dirPath string) error {
	fileInfo, err := os.Stat(dirPath)
	if err == nil {
		if fileInfo.IsDir() {
			return errors.New(dirPath + " already exists")
		}
	}
	return nil
}

func checkIfDirIsEmpty(dirPath string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return nil
	}

	return errors.New("Current directory is not empty, try create a new direcory and retry")
}

func CheckError(err error) {
	if err != nil {
		log.WithFields(log.Fields{"Error": err, "BuildVersion": build.BuildVersion, "BuildTime": build.BuildTime}).Error(err.Error())
		os.Exit(-1)
	}
}

func formatTimestamp(timestamp string) string {
	return strings.Replace(timestamp, "T", " ", 1)
}

func generateCert(systemName string) (string, error) {
	if !isValidSystemName(systemName) {
		return "", errors.New("System name is invalid")
	}

	certManager, err := security.LoadCertManager()
	if err != nil {
		return "", err
	}

	err = certManager.CreateSystemKeystore(
		RootKeystorePath, RootKeystoreAlias,
		CloudKeystorePath, CloudKeystoreAlias,
		systemName+".p12", systemName+"."+CloudKeystoreAlias,
		systemName, security.GenerateSubjectAlternativeName(systemName), KeystorePassword,
	)
	if err != nil {
		return "", err
	}

	authInfo, err := certManager.GetPublicKey(systemName+".p12", KeystorePassword)
	if err != nil {
		return "", err
	}

	// err = certManager.ConvertP12ToPEM(systemName+".p12", KeystorePassword, systemName+".pem", systemName+".key")
	// if err != nil {
	// 	return "", err
	// }

	return authInfo, nil
}

// func createSystemEnvFile(systemName, systemAddress string, systemPort int, keystorePath, keystorePassword, pemCertPath, pemKeyPath string) error {
func createSystemEnvFile(systemName, systemAddress string, systemPort int, keystorePath, keystorePassword string) error {
	filename := systemName + ".env"
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	arrowheadTLSStr := strconv.FormatBool(ArrowheadTLS)
	_, err = file.WriteString("export ARROWHEAD_TLS=" + arrowheadTLSStr + "\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("export ARROWHEAD_VERBOSE=false\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("export ARROWHEAD_AUTHORIZATION_HOST=" + ArrowheadAuthorizationHost + "\n")
	if err != nil {
		return err
	}

	arrowheadAuthorizationPortStr := strconv.Itoa(ArrowheadAuthorizationPort)
	_, err = file.WriteString("export ARROWHEAD_AUTHORIZATION_PORT=" + arrowheadAuthorizationPortStr + "\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("export ARROWHEAD_SERVICEREGISTRY_HOST=" + ArrowheadServiceRegistryHost + "\n")
	if err != nil {
		return err
	}

	arrowheadServiceRegistryPortStr := strconv.Itoa(ArrowheadServiceRegistryPort)
	_, err = file.WriteString("export ARROWHEAD_SERVICEREGISTRY_PORT=" + arrowheadServiceRegistryPortStr + "\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("export ARROWHEAD_ORCHESTRATOR_HOST=" + ArrowheadOrchestratorHost + "\n")
	if err != nil {
		return err
	}

	arrowheadOrchestratorPortStr := strconv.Itoa(ArrowheadOrchestratorPort)
	_, err = file.WriteString("export ARROWHEAD_ORCHESTRATOR_PORT=" + arrowheadOrchestratorPortStr + "\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("export ARROWHEAD_KEYSTORE_PATH=" + keystorePath + "\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("export ARROWHEAD_KEYSTORE_PASSWORD=" + keystorePassword + "\n")
	if err != nil {
		return err
	}

	// _, err = file.WriteString("export ARROWHEAD_PEM_CERT_PATH=" + pemCertPath + "\n")
	// if err != nil {
	// 	return err
	// }
	//
	// _, err = file.WriteString("export ARROWHEAD_PEM_KEY_PATH=" + pemKeyPath + "\n")
	// if err != nil {
	// 	return err
	// }

	_, err = file.WriteString("export ARROWHEAD_TRUSTSTORE=" + Truststore + "\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("export ARROWHEAD_SYSTEM_NAME=" + systemName + "\n")
	if err != nil {
		return err
	}

	_, err = file.WriteString("export ARROWHEAD_SYSTEM_ADDRESS=" + systemAddress + "\n")
	if err != nil {
		return err
	}

	systemPortStr := strconv.Itoa(systemPort)
	_, err = file.WriteString("export ARROWHEAD_SYSTEM_PORT=" + systemPortStr + "\n")
	if err != nil {
		return err
	}

	// _, err = file.WriteString("export GIN_MODE=release\n")
	// if err != nil {
	// 	return err
	// }

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}
