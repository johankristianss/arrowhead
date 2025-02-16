package openssl

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecurityOpenSSLCreateSystemKeystore(t *testing.T) {
	san := "DNS:cardemoconsumer,DNS:cardemoconsumer-ip,DNS:localhost,IP:127.0.0.1"
	if extra := os.Getenv("SUBJECT_ALTERNATIVE_NAME"); extra != "" {
		san += "," + extra
	}

	certManager := CreateOpenSSLCertManager()
	err := certManager.CreateSystemKeystore(
		"../testcerts/master.p12", "arrowhead.eu",
		"../testcerts/c1.p12", "c1.ltu.arrowhead.eu",
		"./cardemoconsumer.p12", "cardemoconsumer.c1.ltu.arrowhead.eu",
		"cardemoconsumer", san, "123456",
	)

	defer os.Remove("./cardemoconsumer.p12")
	defer os.Remove("./cardemoconsumer.pub")

	assert.Nil(t, err)
}

func TestSecurityOpenSSLGetPublicKey(t *testing.T) {
	san := "DNS:cardemoconsumer,DNS:cardemoconsumer-ip,DNS:localhost,IP:127.0.0.1"
	if extra := os.Getenv("SUBJECT_ALTERNATIVE_NAME"); extra != "" {
		san += "," + extra
	}

	certManager := CreateOpenSSLCertManager()
	err := certManager.CreateSystemKeystore(
		"../testcerts/master.p12", "arrowhead.eu",
		"../testcerts/c1.p12", "c1.ltu.arrowhead.eu",
		"./cardemoconsumer.p12", "cardemoconsumer.c1.ltu.arrowhead.eu",
		"cardemoconsumer", san, "123456",
	)

	pubKey, err := certManager.GetPublicKey("./cardemoconsumer.p12", "123456")
	assert.Nil(t, err)
	assert.True(t, len(pubKey) > 0)

	defer os.Remove("./cardemoconsumer.p12")
	defer os.Remove("./cardemoconsumer.pub")

	assert.Nil(t, err)
}

//func ConvertP12ToPEM(p12File, password, outputCert, outputKey string) error {

func TestSecurityOpenSSLConvertP12ToPEM(t *testing.T) {
	san := "DNS:cardemoconsumer,DNS:cardemoconsumer-ip,DNS:localhost,IP:127.0.0.1"
	if extra := os.Getenv("SUBJECT_ALTERNATIVE_NAME"); extra != "" {
		san += "," + extra
	}

	certManager := CreateOpenSSLCertManager()
	err := certManager.CreateSystemKeystore(
		"../testcerts/master.p12", "arrowhead.eu",
		"../testcerts/c1.p12", "c1.ltu.arrowhead.eu",
		"./cardemoconsumer.p12", "cardemoconsumer.c1.ltu.arrowhead.eu",
		"cardemoconsumer", san, "123456",
	)

	defer os.Remove("./cardemoconsumer.p12")
	defer os.Remove("./cardemoconsumer.pub")

	err = certManager.ConvertP12ToPEM("./cardemoconsumer.p12", "123456", "./cardemoconsumer.pem", "./cardemoconsumer.key")
	assert.Nil(t, err)

	defer os.Remove("./cardemoconsumer.pem")
	defer os.Remove("./cardemoconsumer.key")

	assert.Nil(t, err)
}
