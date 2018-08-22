package main

import (
	"github.com/user/rgoq/config"
	"log"
	"github.com/user/rgoq/example/conf"
	"time"
	"github.com/user/rgoq"
)


func readQueu(mTest *rgoq.Queue, chuckSize int) {
	var m conf.Message
	for i :=1; i< chuckSize; i++ {
		errm := mTest.Pull(&m)
		if errm != nil {
			panic(errm)
		}
		log.Println(m.Msg)
		log.Println(m.CreatedAt)
	}
}


func main() {

	log.Println("start worker: " + time.Now().String())
	mTest := rgoq.CreateQueue(config.REDIS_OPTIONS, conf.TEST_KEY)

	chankSize := 100;
	for i := 0; i< 1000; i++ {
		go readQueu(mTest, chankSize)
	}


	log.Println("finish worker: " + time.Now().String())
	time.Sleep(1 * time.Second)
}