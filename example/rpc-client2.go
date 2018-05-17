package main

import (
 "./conf"
"log"
"net/rpc/jsonrpc"
)

func main() {
	//make connection to rpc server
	client, err := jsonrpc.Dial("tcp", ":1234")
	
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	//make arguments object
	args := &conf.Args{
		A: 25,
		B: 3,
	}
	//this will store returned result
	var result conf.Result
	//call remote procedure with args
	err = client.Call("Arith.Multiply", args, &result)
	if err != nil {
		log.Fatalf("error in Arith", err)
	}
	//we got our result in result
	log.Printf("%d*%d=%d\n", args.A, args.B, result)
	log.Printf("%d", result)
}


