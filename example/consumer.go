package main

import (
	"../config"
	"log"
	"mqueue"
	"./conf"

	"time"
)



func main() {

	log.Println("start worker: " + time.Now().String())
	mTest := mqueue.CreateQueue(config.REDIS_OPTIONS, conf.TEST_KEY)
	var m conf.Message

	for i :=1; i< 100000; i++ {
		errm := mTest.Pull(&m)
		if errm != nil {
			panic(errm)
		}
		log.Println(m.Msg)
		log.Println(m.CreatedAt)
	}

	log.Println("finish worker: " + time.Now().String())
}