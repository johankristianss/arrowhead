package keytool

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/johankristianss/arrowhead/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type KeytoolCertManager struct {
}

func CreateKeytoolCertManager() *KeytoolCertManager {
	return &KeytoolCertManager{}
}

// TODO: It might be a good idea to get rid of keytool and use openssl for everything
func (certManager *KeytoolCertManager) runOpenSSLCommand(args ...string) error {
	cmd := exec.Command("openssl", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (certManager *KeytoolCertManager) runKeytoolCommand(args ...string) error {
	cmd := exec.Command("keytool", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (certManager *KeytoolCertManager) CreateSystemKeystore(rootKeystore, rootAlias, cloudKeystore, cloudAlias, systemKeystore, systemDname, systemAlias, san, password string) error {
	rootCertFile := strings.TrimSuffix(rootKeystore, ".p12") + ".crt"
	cloudCertFile := strings.TrimSuffix(cloudKeystore, ".p12") + ".crt"
	systemPubFile := strings.TrimSuffix(systemKeystore, ".p12") + ".pub"
	csrFile := "csrfile.csr"
	signedCertFile := "signed_cert.crt"

	systemKeystore = filepath.Base(systemKeystore)
	systemPubFile = filepath.Base(systemPubFile)

	// Generate the system keystore if it doesn't exist
	if utils.FileExists(systemKeystore) {
		log.WithFields(log.Fields{"systemKeystore": systemKeystore}).Error("System keystore already exists")
		return errors.New("System keystore already exists")
	} else {
		log.WithFields(log.Fields{"systemKeystore": systemKeystore}).Debug("Creating system keystore")
		err := certManager.runKeytoolCommand(
			"-genkeypair", "-v",
			"-keystore", systemKeystore,
			"-storepass", password,
			"-keyalg", "RSA",
			"-keysize", "2048",
			"-validity", "3650",
			"-alias", systemAlias,
			"-keypass", password,
			"-dname", "CN="+systemDname,
			"-ext", "SubjectAlternativeName="+san,
			"-noprompt",
		)
		if err != nil {
			log.WithFields(log.Fields{"Error": err, "systemKeystore": systemKeystore}).Error("Failed to create keystore")
			return errors.New("Failed to create keystore")
		}

		// Import root certificate
		log.WithFields(log.Fields{"rootCertFile": rootCertFile}).Debug("Importing root certificate")
		err = certManager.runKeytoolCommand(
			"-importcert", "-v",
			"-keystore", systemKeystore,
			"-storepass", password,
			"-alias", rootAlias,
			"-file", rootCertFile,
			"-trustcacerts",
			"-noprompt",
		)
		if err != nil {
			log.WithFields(log.Fields{"Error": err, "rootAlias": rootAlias}).Error("Failed to import root certificate")
			return errors.New("Failed to import root certificate")
		}

		// Import cloud certificate
		log.WithFields(log.Fields{"cloudCertFile": cloudCertFile}).Debug("Importing cloud certificate")
		err = certManager.runKeytoolCommand(
			"-importcert", "-v",
			"-keystore", systemKeystore,
			"-storepass", password,
			"-alias", cloudAlias,
			"-file", cloudCertFile,
			"-trustcacerts",
			"-noprompt",
		)
		if err != nil {
			log.WithFields(log.Fields{"Error": err, "cloudAlias": cloudAlias}).Error("Failed to import cloud certificate")
			return errors.New("Failed to import cloud certificate")
		}

		// Generate CSR and save to a file
		log.WithFields(log.Fields{"csrFile": csrFile}).Debug("Generating CSR")
		csrCmd := exec.Command("keytool", "-certreq", "-v",
			"-keystore", systemKeystore,
			"-storepass", password,
			"-alias", systemAlias,
			"-keypass", password,
			"-file", csrFile,
			"-noprompt",
		)
		err = csrCmd.Run()
		if err != nil {
			log.WithFields(log.Fields{"Error": err, "systemAlias": systemAlias}).Error("Failed to generate CSR")
			return errors.New("Failed to generate CSR")
		}

		// Sign the CSR with the cloud keystore
		log.WithFields(log.Fields{"signedCertFile": signedCertFile}).Debug("Signing CSR")
		signedCertCmd := exec.Command("keytool", "-gencert", "-v",
			"-keystore", cloudKeystore,
			"-storepass", password,
			"-validity", "3650",
			"-alias", cloudAlias,
			"-keypass", password,
			"-ext", "SubjectAlternativeName="+san,
			"-rfc",
			"-infile", csrFile,
			"-outfile", signedCertFile,
			"-noprompt",
		)
		err = signedCertCmd.Run()
		if err != nil {
			log.WithFields(log.Fields{"Error": err, "cloudAlias": cloudAlias}).Error("Failed to sign CSR")
			return errors.New("Failed to sign CSR")
		}

		// Import the signed certificate back into the system keystore
		log.WithFields(log.Fields{"signedCertFile": signedCertFile}).Debug("Importing signed certificate")
		importCertCmd := exec.Command("keytool", "-importcert",
			"-keystore", systemKeystore,
			"-storepass", password,
			"-alias", systemAlias,
			"-keypass", password,
			"-trustcacerts",
			"-noprompt",
			"-file", signedCertFile,
			"-noprompt",
		)
		err = importCertCmd.Run()
		if err != nil {
			log.WithFields(log.Fields{"Error": err, "systemAlias": systemAlias}).Error("Failed to import signed certificate")
			return errors.New("Failed to import signed certificate")
		}
	}

	// Create system public key file if missing
	if utils.FileExists(systemPubFile) {
		log.WithFields(log.Fields{"systemPubFile": systemPubFile}).Error("Public key file already exists")
		return errors.New("Public key file already exists")
	} else {
		log.WithFields(log.Fields{"systemPubFile": systemPubFile}).Debug("Creating public key file")
		extractPublicKey(systemKeystore, systemAlias, systemPubFile, password)
	}

	// Clean up temporary files
	os.Remove(csrFile)
	os.Remove(signedCertFile)

	return nil
}

// Extracts the public key from a keystore and writes it to a .pub file.
func extractPublicKey(systemKeystore, systemAlias, systemPubFile, password string) error {
	pubKeyCmd := exec.Command("keytool", "-list",
		"-keystore", systemKeystore,
		"-storepass", password,
		"-alias", systemAlias,
		"-rfc",
		"-noprompt",
	)
	pubKeyOutput, err := pubKeyCmd.Output()
	if err != nil {
		log.WithFields(log.Fields{"Error": err, "systemAlias": systemAlias}).Error("Failed to extract public key")
		return err
	}

	// Convert to PEM format
	opensslCmd := exec.Command("openssl", "x509",
		"-inform", "pem",
		"-pubkey",
		"-noout",
	)
	opensslCmd.Stdin = strings.NewReader(string(pubKeyOutput))
	pubKey, err := opensslCmd.Output()
	if err != nil {
		log.WithFields(log.Fields{"Error": err, "systemAlias": systemAlias}).Error("Failed to convert public key")
		return err
	}

	// Write public key to file
	err = os.WriteFile(systemPubFile, pubKey, 0644)
	if err != nil {
		log.WithFields(log.Fields{"Error": err, "systemPubFile": systemPubFile}).Error("Failed to write public key file")
		return err
	}

	return nil
}

// TODO: Refactor this code to use keytool and not openssl?
func (certManager *KeytoolCertManager) GetPublicKey(keystorePath, password string) (string, error) {
	sslEnabled := true

	if !sslEnabled {
		return "", nil // Return empty string if SSL is disabled
	}

	if _, err := os.Stat(keystorePath); os.IsNotExist(err) {
		return "", errors.New("Error: SSL is enabled, but cardemoprovider.p12 not found")
	}

	// Extract public key from certificate
	opensslCmd := exec.Command("openssl", "pkcs12", "-in", keystorePath, "-clcerts", "-nokeys", "-passin", "pass:"+password)
	certOut, err := opensslCmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to extract certificate: %v", err)
	}

	x509Cmd := exec.Command("openssl", "x509", "-pubkey", "-noout")
	x509Cmd.Stdin = strings.NewReader(string(certOut))
	pubKeyOut, err := x509Cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to extract public key: %v", err)
	}

	// Clean output: Remove headers and newlines
	publicKey := strings.TrimSpace(string(pubKeyOut))
	publicKey = strings.ReplaceAll(publicKey, "-----BEGIN PUBLIC KEY-----", "")
	publicKey = strings.ReplaceAll(publicKey, "-----END PUBLIC KEY-----", "")
	publicKey = strings.ReplaceAll(publicKey, "\n", "")

	return publicKey, nil
}

func (certManager *KeytoolCertManager) ConvertP12ToPEM(p12File, password, outputCert, outputKey string) error {
	err := certManager.runOpenSSLCommand("pkcs12", "-in", p12File, "-clcerts", "-nokeys", "-passin", "pass:"+password, "-out", outputCert)
	if err != nil {
		return fmt.Errorf("Failed to extract certificate: %w", err)
	}

	err = certManager.runOpenSSLCommand("pkcs12", "-in", p12File, "-nocerts", "-nodes", "-passin", "pass:"+password, "-out", outputKey)
	if err != nil {
		return fmt.Errorf("Failed to extract private key: %w", err)
	}

	return nil
}
