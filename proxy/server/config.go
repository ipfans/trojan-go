package server

import (
	"github.com/ipfans/trojan-go/config"
	"github.com/ipfans/trojan-go/proxy/client"
)

func init() {
	config.RegisterConfigCreator(Name, func() interface{} {
		return new(client.Config)
	})
}
