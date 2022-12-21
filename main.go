package main

import (
	"flag"
	"fmt"
	"go_micro_test/client"
	"go_micro_test/server"
)

var (
	mode = flag.String("mode", "none", "go flag test")
)

func main() {
	flag.Parse()
	switch *mode {
	case `s`:
		server.GrpcServer()
	case `c`:
		client.GrpcStart()
	default:
		fmt.Println(`no case`)
	}
}
