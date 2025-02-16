package arrowhead

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"software.sslmate.com/src/go-pkcs12"

	"github.com/gin-contrib/cors"
	"github.com/johankristianss/arrowhead/pkg/rpc"
	log "github.com/sirupsen/logrus"
)

type Framework struct {
	ginHandler        *gin.Engine
	rpcConfig         *rpc.Config
	arrowhead         *rpc.RPC
	systemName        string
	address           string
	port              int
	httpServer        *http.Server
	tlsCertPath       string
	tlsPrivateKeyPath string
	trustStorePath    string
	pemCertPath       string
	pemPrivateKeyPath string
	keyStorePath      string
	keyStorePassword  string
}

func CreateFramework() (*Framework, error) {
	verbose := os.Getenv("ARROWHEAD_VERBOSE")
	if verbose == "false" || verbose == "FALSE" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
	}

	if verbose == "true" || verbose == "TRUE" {
		log.SetLevel(log.DebugLevel)
	}

	framework := &Framework{}
	arrowheadTLSStr := os.Getenv("ARROWHEAD_TLS")
	arrowheadTLS := false
	if arrowheadTLSStr == "true" || arrowheadTLSStr == "TRUE" {
		arrowheadTLS = true
	}

	arrowheadAuthorizationHost := os.Getenv("ARROWHEAD_AUTHORIZATION_HOST")
	arrowheadAuthorizationPortStr := os.Getenv("ARROWHEAD_AUTHORIZATION_PORT")
	arrowheadAuthorizationPort, err := strconv.Atoi(arrowheadAuthorizationPortStr)
	if err != nil {
		return nil, err
	}
	arrowheadServiceRegistryHost := os.Getenv("ARROWHEAD_SERVICEREGISTRY_HOST")
	arrowheadServiceRegistryPortStr := os.Getenv("ARROWHEAD_SERVICEREGISTRY_PORT")
	arrowheadServiceRegistryPort, err := strconv.Atoi(arrowheadServiceRegistryPortStr)
	if err != nil {
		return nil, err
	}
	arrowheadOrchestratorHost := os.Getenv("ARROWHEAD_ORCHESTRATOR_HOST")
	arrowheadOrchestratorPortStr := os.Getenv("ARROWHEAD_ORCHESTRATOR_PORT")
	arrowheadOrchestratorPort, err := strconv.Atoi(arrowheadOrchestratorPortStr)
	if err != nil {
		return nil, err
	}
	keystorePath := os.Getenv("ARROWHEAD_KEYSTORE_PATH")
	framework.keyStorePath = keystorePath
	keystorePassword := os.Getenv("ARROWHEAD_KEYSTORE_PASSWORD")
	framework.keyStorePassword = keystorePassword
	truststore := os.Getenv("ARROWHEAD_TRUSTSTORE")
	framework.trustStorePath = truststore
	framework.systemName = os.Getenv("ARROWHEAD_SYSTEM_NAME")
	framework.address = os.Getenv("ARROWHEAD_SYSTEM_ADDRESS")
	portStr := os.Getenv("ARROWHEAD_SYSTEM_PORT")
	framework.port, err = strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	log.WithFields(log.Fields{"arrowheadTLS": arrowheadTLSStr, "arrowheadAuthorizationHost": arrowheadAuthorizationHost, "arrowheadAuthorizationPort": arrowheadAuthorizationPort, "arrowheadServiceRegistryHost": arrowheadServiceRegistryHost, "arrowheadServiceRegistryPort": arrowheadServiceRegistryPort, "arrowheadOrchestratorHost": arrowheadOrchestratorHost, "arrowheadOrchestratorPort": arrowheadOrchestratorPort, "keystorePath": keystorePath, "truststore": truststore, "systemName": framework.systemName, "address": framework.address, "port": framework.port, "tlsCertPath": framework.tlsCertPath, "tlsPrivateKeyPath": framework.tlsPrivateKeyPath, "trustStorePath": framework.trustStorePath}).Debug("Arrowhead configuration")

	framework.rpcConfig = &rpc.Config{
		TLS:                 arrowheadTLS,
		AuthorizationHost:   arrowheadAuthorizationHost,
		AuthorizationPort:   arrowheadAuthorizationPort,
		ServiceRegistryHost: arrowheadServiceRegistryHost,
		ServiceRegistryPort: arrowheadServiceRegistryPort,
		OrchestratorHost:    arrowheadOrchestratorHost,
		OrchestratorPort:    arrowheadOrchestratorPort,
		KeystorePath:        keystorePath,
		Password:            keystorePassword,
		TruststorePath:      truststore,
	}
	framework.arrowhead, err = rpc.CreateArrowhead(framework.rpcConfig)
	if err != nil {
		return nil, err
	}

	// Setup TLS
	p12Data, err := os.ReadFile(framework.keyStorePath)
	if err != nil {
		return nil, err
	}

	privateKey, certificate, caCerts, err := pkcs12.DecodeChain(p12Data, framework.keyStorePassword)
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

	caCertPEM, err := os.ReadFile(framework.trustStorePath)
	if err != nil {
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

	framework.ginHandler = gin.Default()
	framework.ginHandler.Use(cors.Default())

	framework.httpServer = &http.Server{
		Addr:      ":" + strconv.Itoa(framework.port),
		Handler:   framework.ginHandler,
		TLSConfig: tlsConfig,
	}

	return framework, nil
}

func loadTrustStore(trustStorePath string) (*x509.CertPool, error) {
	caCert, err := ioutil.ReadFile(trustStorePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read truststore: %v", err)
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to append CA certificates")
	}

	return caCertPool, nil
}

func (f *Framework) ServeForever() error {
	// TODO: Examine why the two arguments are empty strings?
	// If we set pem cert and key paths, it won't work
	if err := f.httpServer.ListenAndServeTLS("", ""); err != nil && errors.Is(err, http.ErrServerClosed) {
		fmt.Println(err)
		return err
	}

	return nil
}

func (f *Framework) SendRequest(serviceDef string, params *Params) ([]byte, error) {
	orchestrationRequest := rpc.BuildOrchestrationRequest(f.systemName, f.address, f.port, serviceDef)
	orchestrationResponse, err := f.arrowhead.Client.Orchestrate(orchestrationRequest)
	if err != nil {
		return nil, err
	}

	if len(orchestrationResponse.Response) == 0 {
		return nil, fmt.Errorf("Failed to send request to service: %s", serviceDef)
	}

	return f.arrowhead.SendRequest(&orchestrationResponse.Response[0], params.QueryParams, params.Payload)
}

func (f *Framework) buildLambda(service Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		jsonBytes, err := io.ReadAll(c.Request.Body)
		params := &Params{
			QueryParams: make(map[string]string),
			Payload:     jsonBytes,
		}
		for key, value := range c.Request.URL.Query() {
			if key == "token" {
				// TODO: Verify token
				// token := value[0]
				// fmt.Println(token)
			} else {
				params.QueryParams[key] = value[0] /// TODO: we only support one value per key
			}
		}

		jsonResponse, err := service.HandleRequest(params)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.String(http.StatusOK, string(jsonResponse))
	}
}

func (f *Framework) HandleService(service Service, httpMethod int, serviceDefinition, serviceURI string) {
	log.WithFields(log.Fields{"serviceDefinition": serviceDefinition, "serviceURI": serviceURI}).Debug("Registering service")
	switch httpMethod {
	case rpc.GET:
		f.ginHandler.GET(serviceURI, f.buildLambda(service))
	case rpc.POST:
		f.ginHandler.POST(serviceURI, f.buildLambda(service))
	case rpc.PUT:
		f.ginHandler.PUT(serviceURI, f.buildLambda(service))
	case rpc.DELETE:
		f.ginHandler.DELETE(serviceURI, f.buildLambda(service))
	}
}
