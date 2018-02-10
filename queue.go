package mqueue

import (
	"gopkg.in/redis.v5"
)

const MATER_KEY = "redis:queue"

type Queue struct {
	redisClient *redis.Client
	Name string
}




func CreateQueue(opt *redis.Options, name string) *Queue {
	q :=&Queue{Name: name}
	q.redisClient = redis.NewClient(opt)
	q.redisClient.SAdd(MATER_KEY, name)

	return q
}


func (queue *Queue) LPush(pack string) error{
	qlpush := queue.redisClient.LPush(queue.Name, pack)

	return qlpush.Err();

}