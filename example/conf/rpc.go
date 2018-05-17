package conf


//Holds arguments to be passed to service Arith in RPC call
type Args struct {
	A, B int
}

//Representss service Arith with method Multiply
type Arith int

//Result of RPC call is of this type
type Result int




//This procedure is invoked by rpc and calls rpcexample.Multiply which stores product of args.A and args.B in result pointer
//func (t *Arith) Multiply(args Args, result *Result) error {
//	return Multiply(args, result)
//}
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

//func (t *Arith) Divide(args *Args, quo *Quotient) error {
//	if args.B == 0 {
//		return errors.New("divide by zero")
//	}
//	quo.Quo = args.A / args.B
//	quo.Rem = args.A % args.B
//	return nil
//}


//This procedure is invoked by rpc and calls rpcexample.Multcdiply which stores product of args.A and args.B in result pointer
//func (t *Arith) MultiplyDiff(args Args, responce *ResponceResult) error {
//	return Diff(args, responce)
//}

//stores product of args.A and args.B in result pointer
//func Multiply(args Args, result *Result) error {
//	log.Printf("Multiplying %d with %d\n", args.A, args.B)
//	*result = Result(args.A * args.B)
//	return nil
//}

//func Diff(args Args, responce *ResponceResult) error {
//	log.Printf("Multiplying %d with %d\n", args.A, args.B)
//	*responce = ResponceResult("test message responce")
//	return nil
//}
