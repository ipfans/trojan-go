//go:build custom || full
// +build custom full

package build

import (
	_ "github.com/ipfans/trojan-go/proxy/custom"
	_ "github.com/ipfans/trojan-go/tunnel/adapter"
	_ "github.com/ipfans/trojan-go/tunnel/dokodemo"
	_ "github.com/ipfans/trojan-go/tunnel/freedom"
	_ "github.com/ipfans/trojan-go/tunnel/http"
	_ "github.com/ipfans/trojan-go/tunnel/mux"
	_ "github.com/ipfans/trojan-go/tunnel/router"
	_ "github.com/ipfans/trojan-go/tunnel/shadowsocks"
	_ "github.com/ipfans/trojan-go/tunnel/simplesocks"
	_ "github.com/ipfans/trojan-go/tunnel/socks"
	_ "github.com/ipfans/trojan-go/tunnel/tls"
	_ "github.com/ipfans/trojan-go/tunnel/tproxy"
	_ "github.com/ipfans/trojan-go/tunnel/transport"
	_ "github.com/ipfans/trojan-go/tunnel/trojan"
	_ "github.com/ipfans/trojan-go/tunnel/websocket"
)
