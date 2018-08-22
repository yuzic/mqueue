package main

import (
	"github.com/user/rgoq/config"
	"log"
	"strconv"
	"time"
	"github.com/user/rgoq/example/conf"
	"github.com/user/rgoq"
)


// нужно сделать разные соединители с очередью чтобы

// 1. Нужно сделать общую очередь, чтобы можно было ложить произвольные  структуры в очередь
// 2. Нужна возможнсть делать кастомные consumer для дальнего считывания ключей и навешивания обработчиков


func testKeyOne() {
	log.Println("start worker one key: " + time.Now().String())
	mTest := rgoq.CreateQueue(config.REDIS_OPTIONS, conf.TEST_KEY)
	for i:=1; i< 100000; i++ {
		stest := ("test message " +  strconv.Itoa(i))
		msg := conf.Message{stest, time.Now()}
		err := mTest.Push(msg)
		if err != nil {
			panic(err)
		}
	}

	log.Println("finish worker one_key: " + time.Now().String())
}



func main() {
	testKeyOne()
	//testKeyTwo()
}