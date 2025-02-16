package openssl

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

type OpenSSLCertManager struct {
}

func CreateOpenSSLCertManager() *OpenSSLCertManager {
	return &OpenSSLCertManager{}
}

func (certManager *OpenSSLCertManager) runOpenSSLCommand(args ...string) error {
	cmd := exec.Command("openssl", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (certManager *OpenSSLCertManager) CreateSystemKeystore(rootKeystore, rootAlias, cloudKeystore, cloudAlias, systemKeystore, systemDname, systemAlias, san, password string) error {
	rootCertFile := strings.TrimSuffix(rootKeystore, ".p12") + ".crt"
	cloudCertFile := strings.TrimSuffix(cloudKeystore, ".p12") + ".crt"
	systemPubFile := strings.TrimSuffix(systemKeystore, ".p12") + ".pub"

	csrFile := "csrfile.csr"
	signedCertFile := "signed_cert.crt"
	systemKeyFile := strings.TrimSuffix(systemKeystore, ".p12") + ".key"
	cloudKeyFile := strings.TrimSuffix(cloudKeystore, ".p12") + ".key"

	systemKeystore = filepath.Base(systemKeystore)
	systemPubFile = filepath.Base(systemPubFile)

	if utils.FileExists(systemKeystore) {
		log.WithFields(log.Fields{"systemKeystore": systemKeystore}).Error("System keystore already exists")
		return errors.New("System keystore already exists")
	} else {
		// 1. Generate the system RSA private key.
		log.Debug("Generating system private key...")
		if err := certManager.runOpenSSLCommand("genrsa", "-out", systemKeyFile, "2048"); err != nil {
			log.WithError(err).Error("Failed to generate system private key")
			return err
		}

		// 2. Generate a certificate signing request (CSR) with the desired subject and SAN.
		// The "-addext" option (available in OpenSSL 1.1.1+) is used to add the subjectAltName.
		subj := "/CN=" + systemDname
		log.Debug("Generating CSR for system certificate...")
		if err := certManager.runOpenSSLCommand("req", "-new", "-key", systemKeyFile,
			"-subj", subj,
			"-addext", "subjectAltName="+san,
			"-out", csrFile); err != nil {
			log.WithError(err).Error("Failed to generate CSR")
			return err
		}

		// 3. Extract the cloud CA’s key and certificate if not already available.
		pass := "pass:" + password
		log.Debug("Extracting cloud key from cloud PKCS#12 file...")
		if err := certManager.runOpenSSLCommand("pkcs12", "-in", cloudKeystore, "-nocerts", "-nodes",
			"-passin", pass, "-out", cloudKeyFile); err != nil {
			log.WithError(err).Error("Failed to extract cloud key")
			return errors.New("Failed to extract cloud key")
		}
		if utils.FileExists(cloudCertFile) {
			log.Debug("Extracting cloud certificate from cloud PKCS#12 file...")
			if err := certManager.runOpenSSLCommand("pkcs12", "-in", cloudKeystore, "-clcerts", "-nokeys",
				"-passin", pass, "-out", cloudCertFile); err != nil {
				log.WithError(err).Error("Failed to extract cloud certificate")
				return errors.New("Failed to extract cloud certificate")
			}
		} else {
			log.WithFields(log.Fields{"cloudCertFile": cloudCertFile}).Error("Cloud certificate file not found")
			return errors.New("Cloud certificate file not found")
		}

		// 4. Create a temporary extension file to supply the subjectAltName when signing.
		extFile := "v3ext.cnf"
		extContent := fmt.Sprintf("[v3_ext]\nsubjectAltName = %s\n", san)
		if err := os.WriteFile(extFile, []byte(extContent), 0644); err != nil {
			log.WithError(err).Error("Failed to write extension file")
			return err
		}

		// 5. Sign the CSR with the cloud CA’s key and certificate.
		log.Debug("Signing CSR with cloud CA...")
		if err := certManager.runOpenSSLCommand("x509", "-req",
			"-in", csrFile,
			"-CA", cloudCertFile,
			"-CAkey", cloudKeyFile,
			"-CAcreateserial", // creates a file (e.g. cloudCertFile.srl) with the serial number
			"-out", signedCertFile,
			"-days", "3650",
			"-extfile", extFile,
			"-extensions", "v3_ext"); err != nil {
			log.WithError(err).Error("Failed to sign CSR with cloud CA")
			return err
		}

		// 6. Build the certificate chain file.
		// The chain file contains first the cloud certificate, then the root certificate.
		chainFile := "chain.pem"
		chainContent := ""
		cloudCertBytes, err := os.ReadFile(cloudCertFile)
		if err != nil {
			log.WithError(err).Error("Failed to read cloud certificate")
			return err
		}
		chainContent += string(cloudCertBytes) + "\n"
		rootCertBytes, err := os.ReadFile(rootCertFile)
		if err != nil {
			log.WithError(err).Error("Failed to read root certificate")
			return err
		}
		chainContent += string(rootCertBytes) + "\n"
		if err := os.WriteFile(chainFile, []byte(chainContent), 0644); err != nil {
			log.WithError(err).Error("Failed to write chain file")
			return err
		}

		// 7. Create the system PKCS#12 keystore.
		// It bundles the system’s private key, the signed certificate, and the CA chain.
		log.Debug("Creating system PKCS#12 keystore...")
		if err := certManager.runOpenSSLCommand("pkcs12", "-export",
			"-inkey", systemKeyFile,
			"-in", signedCertFile,
			"-certfile", chainFile,
			"-out", systemKeystore,
			"-passout", pass); err != nil {
			log.WithError(err).Error("Failed to create system PKCS#12 keystore")
			return err
		}

		// 8. Extract the system public key from the signed certificate.
		log.Debug("Extracting system public key...")
		pubKeyCmd := exec.Command("openssl", "x509", "-in", signedCertFile, "-pubkey", "-noout")
		pubKeyOut, err := pubKeyCmd.CombinedOutput()
		if err != nil {
			log.WithFields(log.Fields{"Error": err, "Output": string(pubKeyOut)}).
				Error("Failed to extract system public key")
			return err
		}
		if err := os.WriteFile(systemPubFile, pubKeyOut, 0644); err != nil {
			log.WithError(err).Error("Failed to write system public key file")
			return err
		}

		// 9. Clean up temporary files.
		os.Remove(csrFile)
		os.Remove(signedCertFile)
		os.Remove(extFile)
		os.Remove(chainFile)
		if utils.FileExists(cloudCertFile + ".srl") {
			os.Remove(cloudCertFile + ".srl")
		}
		if utils.FileExists(cloudCertFile + ".key") {
			os.Remove(cloudCertFile + ".key")
		}
		if utils.FileExists(systemKeyFile) {
			os.Remove(systemKeyFile)
		}
	}

	// In case the system public key file is still missing, extract it from the system keystore.
	if !utils.FileExists(systemPubFile) {
		tempCert := "temp_cert.crt"
		if err := certManager.runOpenSSLCommand("pkcs12", "-in", systemKeystore, "-nokeys", "-clcerts",
			"-passin", "pass:"+os.Getenv("PASSWORD"), "-out", tempCert); err != nil {
			log.WithError(err).Error("Failed to extract certificate from system keystore")
			return err
		}
		pubKeyCmd := exec.Command("openssl", "x509", "-in", tempCert, "-pubkey", "-noout")
		pubKeyOut, err := pubKeyCmd.CombinedOutput()
		if err != nil {
			log.WithError(err).Error("Failed to extract public key from certificate")
			return err
		}
		if err := os.WriteFile(systemPubFile, pubKeyOut, 0644); err != nil {
			log.WithError(err).Error("Failed to write system public key file")
			return err
		}
		os.Remove(tempCert)
	}

	return nil
}

func (certManager *OpenSSLCertManager) GetPublicKey(keystorePath, password string) (string, error) {
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
		return "", fmt.Errorf("Failed to extract certificate: %v", err)
	}

	x509Cmd := exec.Command("openssl", "x509", "-pubkey", "-noout")
	x509Cmd.Stdin = strings.NewReader(string(certOut))
	pubKeyOut, err := x509Cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Failed to extract public key: %v", err)
	}

	// Clean output: Remove headers and newlines
	publicKey := strings.TrimSpace(string(pubKeyOut))
	publicKey = strings.ReplaceAll(publicKey, "-----BEGIN PUBLIC KEY-----", "")
	publicKey = strings.ReplaceAll(publicKey, "-----END PUBLIC KEY-----", "")
	publicKey = strings.ReplaceAll(publicKey, "\n", "")

	return publicKey, nil
}

func (certManager *OpenSSLCertManager) ConvertP12ToPEM(p12File, password, outputCert, outputKey string) error {
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
