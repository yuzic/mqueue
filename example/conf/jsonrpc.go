package conf

import (
	"log"
	"rgoq"
	"../../config"
	"fmt"
	"os"
)

//Holds arguments to be passed to service Arith in RPC call
type Arguments struct {
	Message, Created_at string
}

//Representss service Arith with method Multiply
type PackageJsonRpc string



type ResponceResultString string

//This procedure is invoked by rpc and calls rpcexample.Multcdiply which stores product of args.A and args.B in result pointer
func (t *PackageJsonRpc) PushResult(args Arguments, responce *ResponceResultString) error {
	return Push(args, responce)
}


//This procedure is invoked by rpc and calls rpcexample.Multcdiply which stores product of args.A and args.B in result pointer
func (t *PackageJsonRpc) PullResult(args Arguments, responce *ResponceResultString) error {
	return Pull(args, responce)
}


func Push(args Arguments, responce *ResponceResultString) error {
	log.Printf("Multiplying %d with %d\n", args.Message, args.Created_at)
	// пишим в очередь
	mTest := rgoq.CreateQueue(config.REDIS_OPTIONS, config.SERVER_KEY)
	err := mTest.PushString(string( args.Message))
	checkError(err)
	*responce = ResponceResultString("test message responce")

	return nil
}


func Pull(args Arguments, responce *ResponceResultString) error {
	//log.Printf("Run Pull %d with %d\n", args.A, args.B)
	// пишим в очередь
	mTest := rgoq.CreateQueue(config.REDIS_OPTIONS, config.SERVER_KEY)
	*responce = ResponceResultString(mTest.PullString())

	return nil
}


func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}
