package mpesa

import (
	"context"
	"fmt"
	"github.com/pesakit/pesakit/pkg/crypto"
	"time"
)

func (c *Client) getEncryptionKey() (string, error) {
	isAvailable := c.encryptedApiKey != nil && *c.encryptedApiKey != ""
	//notExpired := client.sessionExpiration.Sub(time.Now()).Minutes() > 1
	if isAvailable {
		return *c.encryptedApiKey, nil
	}

	return crypto.EncryptKey(c.Conf.APIKey, c.Conf.PublicKey)
}

//checkSessionID examine if there is a session id saved as Client.sessionID
//if it is available it checks if it has already expired or have more than
//1 minute till expiration date and returns it
//if the above conditions are not fulfilled it calls Client.SessionID
//then save it and increment the expiration date
func (c *Client) checkSessionID() (string, error) {
	sessAvailable := c.sessionID != nil && *c.sessionID != ""
	sessExpiresAt := c.sessionExpiration
	sessExpired := !sessExpiresAt.IsZero() && time.Until(sessExpiresAt) < (60*time.Second)

	if sessAvailable && !sessExpired {
		return *c.sessionID, nil
	}

	resp, err := c.SessionID(context.Background())
	if err != nil {
		return "", fmt.Errorf("could not fetch session id: %w", err)
	}

	if resp.OutputErr != "" {
		return "", fmt.Errorf("could not fetch session id: %s", resp.OutputErr)
	}

	return resp.ID, err

}
