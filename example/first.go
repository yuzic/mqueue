package main

import (
	"../config"
	"log"
	"gopkg.in/redis.v5"
	"fmt"
	"encoding/json"
	"mqueue"
)

type Message struct {
	Msg string
	Created_at string
}



func readQueue(redisKey string, client *redis.Client) string {
	//var client = redis.NewClient(config.REDIS_WW_OPTIONS)
	var ququeList = client.RPopLPush(redisKey, "t")
	log.Println("test application")

	return ququeList.Val()
	//log.Panic("Hash is empty in redis")5
	//panic("Hash is empty in redis")

}

// 1. Нужно сделать общую очередь, чтобы можно было ложить произвольные  структуры в очередь
// 2.
func main() {
	var redisKey = "ddwrt"
	mTest := mqueue.CreateQueue(config.REDIS_WW_OPTIONS, redisKey)



	msg := Message{"test message", "2018-02-09"}
	js, _ := json.Marshal(msg)
	mTest.LPush(string(js))


	var client = redis.NewClient(config.REDIS_WW_OPTIONS)

	//s := string(js)
	//var m Message
	//fmt.Println(json.Unmarshal([]byte(s), &m))

	//client.HSet(redisKey, "subkey", string(js))
	//client.RPush(redisKey, string(js))
	//
	r := readQueue(redisKey, client);
	//var m Message

	var m Message
	//r:= client.HGet(redisKey, "subkey").Val()

	//
	err := json.Unmarshal([]byte(r), &m)

	if err != nil {
		log.Printf("error [%s]", err.Error())
	}

	fmt.Println(m.Msg)
	fmt.Println(m.Created_at)
	//fmt.Println(r)
	//fmt.Printf("key[%s]\n", r.msg)

}