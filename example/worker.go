package main

import (
	"../config"
	"log"
	"mqueue"
	"strconv"
	"time"
	"./conf"
)


// нужно сделать разные соединители с очередью чтобы

// 1. Нужно сделать общую очередь, чтобы можно было ложить произвольные  структуры в очередь
// 2. Нужна возможнсть делать кастомные consumer для дальнего считывания ключей и навешивания обработчиков


func testKeyOne() {
	log.Println("start worker one key: " + time.Now().String())
	mTest := mqueue.CreateQueue(config.REDIS_OPTIONS, conf.TEST_KEY)
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


func testKeyTwo() {
	log.Println("start worker two key: " + time.Now().String())
	mTest := mqueue.CreateQueue(config.REDIS_OPTIONS, conf.TEST_KEY2)
	for i:=1; i< 10; i++ {
		stest := ("test message " +  strconv.Itoa(i))
		msg := conf.Message{stest, time.Now()}
		err := mTest.Push(msg)
		if err != nil {
			panic(err)
		}
	}

	log.Println("finish worker two key: " + time.Now().String())
}



func main() {
	testKeyOne()
	testKeyTwo()
}