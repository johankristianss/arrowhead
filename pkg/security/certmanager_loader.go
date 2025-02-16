package security

import (
	"errors"
	"io"
	"os/exec"

	"github.com/johankristianss/arrowhead/pkg/security/keytool"
	"github.com/johankristianss/arrowhead/pkg/security/openssl"
	log "github.com/sirupsen/logrus"
)

func GenerateSubjectAlternativeName(name string) string {
	return "DNS:" + name + ",DNS:" + name + "-ip,DNS:localhost,IP:127.0.0.1"
}

func LoadCertManager() (CertManager, error) {
	opensslFound := true
	cmd := exec.Command("openssl", "version")
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	err := cmd.Run()
	if err != nil {
		opensslFound = false
	} else {
		log.Debug("openssl command found")
	}

	keytoolFound := true
	cmd = exec.Command("keytool")
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard
	err = cmd.Run()
	if err != nil {
		keytoolFound = false
	} else {
		log.Debug("keytool command found")
	}

	if opensslFound {
		log.Info("Using openssl as certificate manager")
		return openssl.CreateOpenSSLCertManager(), nil
	}

	if keytoolFound {
		log.Info("Using keytool as certificate manager")
		return keytool.CreateKeytoolCertManager(), nil
	}

	log.Error("No certificate manager found. Please install either openssl or keytool.")
	return nil, errors.New("No certificate manager found. Please install either openssl or keytool.")
}
