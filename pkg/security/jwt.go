package security

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
)

func LoadRSAPrivateKey(filename string) (*rsa.PrivateKey, error) {
	keyBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// DecryptJWE decrypts a JWE token using a private key
func DecryptJWE(encryptedJWT string, privateKey *rsa.PrivateKey) (string, error) {
	decrypted, err := jwe.Decrypt([]byte(encryptedJWT), jwe.WithKey(jwa.RSA_OAEP, privateKey))
	if err != nil {
		return "", fmt.Errorf("error decrypting JWE: %v", err)
	}
	return string(decrypted), nil
}

// func LoadRSAPrivateKey(filename string) (*rsa.PrivateKey, error) {
// 	keyBytes, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return privateKey, nil
// }
//
// // DecryptJWE decrypts a JWE token using a private key
// func DecryptJWE(encryptedJWT string, privateKey *rsa.PrivateKey) (string, error) {
// 	if encryptedJWT == "" {
// 		return "", errors.New("empty encrypted JWT")
// 	}
//
// 	decrypted, err := jwe.Decrypt([]byte(encryptedJWT), jwe.WithKey(jwa.RSA_OAEP, privateKey))
// 	if err != nil {
// 		return "", fmt.Errorf("error decrypting JWE: %v", err)
// 	}
// 	return string(decrypted), nil
// }
//
//
// func LoadRSAPrivateKey(path string) (*rsa.PrivateKey, error) {
// 	keyBytes, err := os.ReadFile(path)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read private key file: %v", err)
// 	}
//
// 	block, _ := pem.Decode(keyBytes)
// 	if block == nil {
// 		return nil, fmt.Errorf("failed to decode PEM block containing private key")
// 	}
//
// 	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse private key: %v", err)
// 	}
//
// 	rsaKey, ok := privateKey.(*rsa.PrivateKey)
// 	if !ok {
// 		return nil, fmt.Errorf("not an RSA private key")
// 	}
//
// 	return rsaKey, nil
// }
//
// // Decrypt JWE using the RSA private key
// func DecryptJWE(jweToken string, privateKey *rsa.PrivateKey) (string, error) {
// 	decrypted, err := jwe.Decrypt([]byte(jweToken), jwe.WithKey(jwa.RSA_OAEP_256, privateKey))
// 	if err != nil {
// 		return "", fmt.Errorf("failed to decrypt JWE: %v", err)
// 	}
// 	return string(decrypted), nil
// }
//
//
// // Parse RSA Public Key from PEM string
// func ParseRSAPublicKey(pubKeyPEM string) (*rsa.PublicKey, error) {
// 	block, _ := pem.Decode([]byte(pubKeyPEM))
// 	if block == nil || block.Type != "PUBLIC KEY" {
// 		return nil, errors.New("failed to decode PEM block containing public key")
// 	}
//
// 	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse public key: %v", err)
// 	}
//
// 	rsaPub, ok := pub.(*rsa.PublicKey)
// 	if !ok {
// 		return nil, errors.New("not an RSA public key")
// 	}
//
// 	return rsaPub, nil
// }
//
// // Verify JWT Signature
// func VerifyJWT(tokenString string, publicKey *rsa.PublicKey) (*jwt.Token, error) {
// 	// Ensure JWT is properly formatted
// 	parts := strings.Split(tokenString, ".")
// 	if len(parts) != 3 {
// 		return nil, fmt.Errorf("invalid JWT format: expected 3 parts, got %d", len(parts))
// 	}
//
// 	// Parse and verify the JWT
// 	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Ensure the token is signed using RS256
// 		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return publicKey, nil
// 	})
// }
//
//
// func ParseRSAPublicKey(pubKeyPEM string) (*rsa.PublicKey, error) {
// 	block, _ := pem.Decode([]byte(pubKeyPEM))
// 	if block == nil || block.Type != "PUBLIC KEY" {
// 		return nil, fmt.Errorf("failed to decode PEM block containing public key")
// 	}
//
// 	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse public key: %v", err)
// 	}
//
// 	rsaKey, ok := pubKey.(*rsa.PublicKey)
// 	if !ok {
// 		return nil, fmt.Errorf("not an RSA public key")
// 	}
//
// 	return rsaKey, nil
// }
//
// func VerifyJWT(tokenString string, publicKey *rsa.PublicKey) (*jwt.Token, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return publicKey, nil
// 	})
//
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	if !token.Valid {
// 		return nil, fmt.Errorf("invalid token")
// 	}
//
// 	return token, nil
// }
