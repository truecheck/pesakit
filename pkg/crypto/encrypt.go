package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

// EncryptKey ....
//1.	Copy and save the API Key.
//2.	Copy the Public Key from the below section.
//3.	Generate a decoded Base64 string from the Public Key
//4.	Generate an instance of an RSA cipher and use the Base 64 string as the input
//5.	Encode the API Key with the RSA cipher and digest as Base64 string format
//6.	The result is your encrypted API Key.
func EncryptKey(apiKey, pubKey string) (string, error) {
	decodedBase64, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return "", fmt.Errorf("could not decode pub key to Base64 string: %w", err)
	}
	publicKeyInterface, err := x509.ParsePKIXPublicKey(decodedBase64)
	if err != nil {
		return "", fmt.Errorf("could not parse encoded public key: %w", err)
	}

	//check if the public key is RSA public key
	publicKey, isRSAPublicKey := publicKeyInterface.(*rsa.PublicKey)
	if !isRSAPublicKey {
		return "", fmt.Errorf("public key parsed is not an RSA public key : %w", err)
	}

	msg := []byte(apiKey)

	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, msg)

	if err != nil {
		return "", fmt.Errorf("could not encrypt key using generated public key: %w", err)
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}
