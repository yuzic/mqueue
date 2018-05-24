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
	//args := &conf.Message{
	//	Msg: "{subsribe:199, validate:true}",
	//	CreatedAt: time.Now(),
	//}

	args := &conf.Arguments{
		Message: "{subscribe:1, validate:true, Russian:'Русский текст'}",
		Created_at: "12323213",
	}
	//this will store returned result
	var result conf.ResponceResultString
	//call remote procedure with args
	err = client.Call("PackageJsonRpc.PushResult", args, &result)
	if err != nil {
		log.Fatalf("error in Arith", err)
	}
	//we got our result in result
	log.Printf("%d\n", result)
}


