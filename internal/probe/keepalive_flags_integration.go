package probe

import (
	"fmt"
	"time"

	"google.golang.org/grpc/keepalive"
)

// KeepaliveDialOption returns the gRPC keepalive client parameters dial option
// based on the KeepaliveConfig. Returns nil if keepalive is disabled.
func (c *KeepaliveConfig) DialOption() (keepalive.ClientParameters, bool) {
	if c == nil || !c.Enabled {
		return keepalive.ClientParameters{}, false
	}
	return keepalive.ClientParameters{
		Time:                c.Time,
		Timeout:             c.Timeout,
		PermitWithoutStream: c.PermitWithoutStream,
	}, true
}

// String returns a human-readable representation of the KeepaliveConfig.
func (c *KeepaliveConfig) String() string {
	if c == nil || !c.Enabled {
		return "keepalive: disabled"
	}
	return fmt.Sprintf("keepalive: time=%s timeout=%s permit_without_stream=%v",
		c.Time.Round(time.Millisecond),
		c.Timeout.Round(time.Millisecond),
		c.PermitWithoutStream,
	)
}
