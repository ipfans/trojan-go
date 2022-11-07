//go:build api || full
// +build api full

package build

import (
	_ "github.com/ipfans/trojan-go/api/control"
	_ "github.com/ipfans/trojan-go/api/service"
)
