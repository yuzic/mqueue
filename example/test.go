package main

import (
	"../config"
	"log"
	"mqueue"
	"mqueue/example/conf"
)

//func b() {
//	defer fmt.Print(88)
//
//	for i :=0; i <4; i++ {
//		defer  fmt.Print(i)
//	}
//}

func main() {
	mtest := mqueue.CreateQueue(config.REDIS_OPTIONS, conf.TEST_KEY)
	log.Println(mtest.Length())

}