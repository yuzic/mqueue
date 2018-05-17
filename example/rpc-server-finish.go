package main

import (
	"errors"
	"net/rpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}


func main() {
	//arith instance
	arith := new(Arith)

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