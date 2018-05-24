package main

import (
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
	"./conf"
)


//
//func Push(args Arguments, responce *ResponceResultString) error {
//	log.Printf("Multiplying %d with %d\n", args.A, args.B)
//	// пишим в очередь
//	mTest := rgoq.CreateQueue(config.REDIS_OPTIONS, config.SERVER_KEY)
//	err := mTest.PushString(string( args.A))
//	checkError(err)
//	*responce = ResponceResultString("test message responce")
//
//	return nil
//}



func main() {
	//arith instance
	arith := new(conf.PackageJsonRpc)
	//make listen in 127.0.0.1:4444
	l, e := net.Listen("tcp", ":1234")
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