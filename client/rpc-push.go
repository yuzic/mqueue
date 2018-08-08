package main

import (
 	"github.com/user/rgoq/example/conf"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	//make connection to rpc server
	client, err := jsonrpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}
	args := &conf.Arguments{
		Message: "{subscribe:1, validate:true, Russian:'rerere'}",
		Created_at: "12323213",
	}
	//this will store returned result
	var result conf.ResponceResultString
	//call remote procedure with args
	err = client.Call("PackageJsonRpc.Push", args, &result)
	if err != nil {
		log.Fatalf("error in Arith", err)
	}
	//we got our result in result
	log.Printf("%d\n", result)
}


