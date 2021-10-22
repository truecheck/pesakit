//

package airtel

import (
	"context"
	"fmt"
	"github.com/techcraftlabs/base"
	"time"
)

type Authenticator interface {
	Token(ctx context.Context) (TokenResponse, error)
}

func (c *Client) checkToken(ctx context.Context) (string, error) {
	var token string
	if *c.token == "" {
		str, err := c.Token(ctx)
		if err != nil {
			return "", err
		}
		token = fmt.Sprintf("%s", str.AccessToken)
	}
	//Add Auth Header
	if *c.token != "" {
		if !c.tokenExpiresAt.IsZero() && time.Until(c.tokenExpiresAt) < (60*time.Second) {
			if _, err := c.Token(ctx); err != nil {
				return "", err
			}
		}
		token = *c.token
	}

	return token, nil
}

func (c *Client) Token(ctx context.Context) (TokenResponse, error) {
	body := TokenRequest{
		ClientID:     c.Conf.ClientID,
		ClientSecret: c.Conf.Secret,
		GrantType:    defaultGrantType,
	}

	var opts []base.RequestOption
	hs := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "*/*",
	}

	opts = append(opts, base.WithRequestHeaders(hs))
	req := c.makeInternalRequest(Authorization, body, opts...)

	res := new(TokenResponse)
	reqName := Authorization.String()
	_, err := c.base.Do(ctx, reqName, req, res)
	if err != nil {
		return TokenResponse{}, err
	}
	duration := time.Duration(res.ExpiresIn)
	now := time.Now()
	later := now.Add(time.Second * duration)
	c.tokenExpiresAt = later
	*c.token = res.AccessToken
	return *res, nil
}

//func PinEncryption(pin string, pubKey string) (string, error) {
//
//	decodedBase64, err := base64.StdEncoding.DecodeString(pubKey)
//	if err != nil {
//		return "", fmt.Errorf("could not decode pub key to Base64 string: %w", err)
//	}
//
//	publicKeyInterface, err := x509.ParsePKIXPublicKey(decodedBase64)
//	if err != nil {
//		return "", fmt.Errorf("could not parse encoded public key (encryption key) : %w", err)
//	}
//
//	//check if the public key is RSA public key
//	publicKey, isRSAPublicKey := publicKeyInterface.(*rsa.PublicKey)
//	if !isRSAPublicKey {
//		return "", fmt.Errorf("public key parsed is not an RSA public key : %w", err)
//	}
//
//	msg := []byte(pin)
//	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, msg)
//	if err != nil {
//		return "", fmt.Errorf("could not encrypt api key using generated public key: %w", err)
//	}
//
//	return base64.StdEncoding.EncodeToString(encrypted), nil
//
//}
