package rpc

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/johankristianss/arrowhead/pkg/core"
	log "github.com/sirupsen/logrus"

	pkcs12 "software.sslmate.com/src/go-pkcs12"
)

type Config struct {
	TLS                 bool
	AuthorizationHost   string
	AuthorizationPort   int
	ServiceRegistryHost string
	ServiceRegistryPort int
	OrchestratorHost    string
	OrchestratorPort    int
	KeystorePath        string
	TruststorePath      string
	Password            string
}

type RPC struct {
	http       *http.Client
	Config     *Config
	Management *Management
	Client     *Client
}

type Management struct {
	rpc *RPC
}

type Client struct {
	rpc *RPC
}

func CreateArrowhead(config *Config) (*RPC, error) {
	rpc := &RPC{Config: config}
	rpc.Management = &Management{rpc: rpc}
	rpc.Client = &Client{rpc: rpc}

	var err error
	rpc.http, err = rpc.createHTTPClient(config.KeystorePath, config.Password, config.TruststorePath)
	if err != nil {
		return nil, err
	}

	return rpc, nil
}

func (rpc *RPC) createHTTPClient(certPath string, password string, truststorePath string) (*http.Client, error) {
	p12Data, err := os.ReadFile(certPath)
	if err != nil {
		log.WithFields(log.Fields{"Error": err, "certPath": certPath}).Error("failed to read certificate file")
		return nil, err
	}

	privateKey, certificate, caCerts, err := pkcs12.DecodeChain(p12Data, password)
	if err != nil {
		log.WithFields(log.Fields{"Error": err}).Error("failed to decode PKSCS12 chain")
		return nil, err
	}

	var certChain [][]byte
	certChain = append(certChain, certificate.Raw)
	for _, cert := range caCerts {
		certChain = append(certChain, cert.Raw)
	}

	tlsCert := tls.Certificate{
		Certificate: certChain,
		PrivateKey:  privateKey,
	}

	caCertPEM, err := os.ReadFile(truststorePath)
	if err != nil {
		log.WithFields(log.Fields{"Error": err, "truststorePath": truststorePath}).Error("failed to read truststore file")
		return nil, err
	}

	caPool := x509.NewCertPool()
	if ok := caPool.AppendCertsFromPEM(caCertPEM); !ok {
		log.WithFields(log.Fields{"Error": err}).Error("failed to append CA certificate(s)")
		return nil, errors.New("failed to append CA certificate(s)")
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		RootCAs:      caPool,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	return &http.Client{
		Transport: transport,
	}, nil
}

func (rpc *RPC) SendRequest(matchedService *core.MatchedService, params map[string]string, payload []byte) ([]byte, error) {
	address := matchedService.Provider.Address
	port := strconv.Itoa(matchedService.Provider.Port)

	token := matchedService.AuthorizationTokens["HTTP-SECURE-JSON"]
	if token == "" {
		return nil, errors.New("No token found")
	}
	url := "https://" + address + ":" + port + matchedService.ServiceURI + "?token=" + token

	// TODO: append query parameters

	method := matchedService.Metadata["http-method"]
	if method == "" {
		return nil, errors.New("No HTTP method found")
	}

	log.WithFields(log.Fields{"method": method, "URL": url}).Debug("Sending request")

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return makeRequest(rpc.http, req, 200, "Failed to send request")
}
