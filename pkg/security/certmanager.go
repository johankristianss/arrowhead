package security

type CertManager interface {
	CreateSystemKeystore(rootKeystore, rootAlias, cloudKeystore, cloudAlias, systemKeystore, systemDname, systemAlias, san, password string) error
	GetPublicKey(keystorePath, password string) (string, error)
	ConvertP12ToPEM(p12File, password, outputCert, outputKey string) error
}
