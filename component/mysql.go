//go:build mysql || full || mini
// +build mysql full mini

package build

import (
	_ "github.com/ipfans/trojan-go/statistic/mysql"
)
