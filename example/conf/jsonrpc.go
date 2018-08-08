package conf

import (
	"log"
	"github.com/user/rgoq/config"
	"fmt"
	"os"
	"github.com/user/rgoq"
)

//Holds arguments to be passed to service Arith in RPC call
type Arguments struct {
	Message, Created_at string
}

//Representss service Arith with method Multiply
type PackageJsonRpc string



type ResponceResultString string


func (t *PackageJsonRpc)  Push(args Arguments, responce *ResponceResultString) error {
	log.Printf("Debug %d with %d\n", args.Message, args.Created_at)
	// пишим в очередь
	mTest := rgoq.CreateQueue(config.REDIS_OPTIONS, config.SERVER_KEY)
	err := mTest.PushString(string( args.Message))
	checkError(err)
	*responce = ResponceResultString("ok")

	return nil
}


func (t *PackageJsonRpc) Pull(args Arguments, responce *ResponceResultString) error {
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
