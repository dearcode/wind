package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/dearcode/crab/handler"
	"github.com/dearcode/wind/server"
	_ "github.com/dearcode/wind/server/view"
)

func main() {
	addr := flag.String("h", ":9000", "api listen address")
	flag.Parse()

	if err := server.Init(); err != nil {
		panic(err.Error())
	}

	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err.Error())
	}

	if err = http.Serve(ln, handler.Server); err != nil {
		panic(err.Error())
	}

}
