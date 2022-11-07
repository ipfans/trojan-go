package main

import (
	"flag"

	_ "github.com/ipfans/trojan-go/component"
	"github.com/ipfans/trojan-go/log"
	"github.com/ipfans/trojan-go/option"
)

func main() {
	flag.Parse()
	for {
		h, err := option.PopOptionHandler()
		if err != nil {
			log.Fatal("invalid options")
		}
		err = h.Handle()
		if err == nil {
			break
		}
	}
}
