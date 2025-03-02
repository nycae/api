package main

import (
	"flag"

	"github.com/sekuradev/api/gateway/pkg/gateway"
)

func main() {
	flag.Parse()
	gateway.NewServer(gateway.UIHandler()).Serve()
}
