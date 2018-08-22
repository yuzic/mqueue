package main

import (
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
	"github.com/user/rgoq/config"
	"github.com/user/rgoq/packages"
)

func main() {
	arith := new(jsonpackage.PackageJsonRpc)
	l, e := net.Listen("tcp", config.JSON_RPC_PORT)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	defer l.Close()

	//instance rpc server
	rpcserver := rpc.NewServer()
	rpcserver.Register(arith)
	//updated for multiples requests
	for {
		//block until acceptation of client
		c, e := l.Accept()
		if e != nil {
			log.Fatal("accept error:", e)
		}
		//instance codec
		jsoncodec := jsonrpc.NewServerCodec(c)
		rpcserver.ServeCodec(jsoncodec)
	}
}