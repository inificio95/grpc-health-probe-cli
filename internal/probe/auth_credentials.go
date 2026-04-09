package probe

import (
	"context"
	"encoding/base64"
	"fmt"
)

// bearerCredentials implements credentials.PerRPCCredentials for Bearer tokens.
type bearerCredentials struct {
	token string
}

func (b *bearerCredentials) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", b.token),
	}, nil
}

func (b *bearerCredentials) RequireTransportSecurity() bool {
	return false
}

// basicCredentials implements credentials.PerRPCCredentials for Basic auth.
type basicCredentials struct {
	username string
	password string
}

func (c *basicCredentials) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	encoded := base64.StdEncoding.EncodeToString(
		[]byte(fmt.Sprintf("%s:%s", c.username, c.password)),
	)
	return map[string]string{
		"authorization": fmt.Sprintf("Basic %s", encoded),
	}, nil
}

func (c *basicCredentials) RequireTransportSecurity() bool {
	return false
}
