package transport

import (
	"net"

	"github.com/ipfans/trojan-go/tunnel"
)

type Conn struct {
	net.Conn
}

func (c *Conn) Metadata() *tunnel.Metadata {
	return nil
}
