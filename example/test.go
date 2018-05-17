package main

import (
	"fmt"
)

//import (
//	"../config"
//	"log"
//	"rgoq"
//	"rgoq/example/conf"
//)

//func b() {
//	defer fmt.Print(88)
//
//	for i :=0; i <4; i++ {
//		defer  fmt.Print(i)
//	}
//}



func testq() {
	fmt.Println(2)
}

func main() {
//	mtest := rgoq.CreateQueue(config.REDIS_OPTIONS, conf.TEST_KEY)
//	log.Println(mtest.Length())
	//go testq()
	//go testq()

	go func () {
		fmt.Println(33)
	}()

	//time.Sleep(1 * time.Second)


}