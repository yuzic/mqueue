package rgoq

import (
	"gopkg.in/redis.v5"
)

const MASTER_KEY = "redis:queue"

type Queue struct {
	redisClient *redis.Client
	Name string
}


func CreateQueue(opt *redis.Options, name string) *Queue {
	q :=&Queue{Name: name}
	q.redisClient = redis.NewClient(opt)
	q.redisClient.SAdd(MASTER_KEY, name)
	return q
}


func (queue *Queue) Push(v interface{}) error {
	sjon, err := jsonEncode(v)
	if err != nil {
		return err
	}
	qpush := queue.redisClient.LPush(queue.Name, sjon)
	return qpush.Err();
}

func (queue *Queue) PushString(sjon string) error {
	qpush := queue.redisClient.LPush(queue.Name, sjon)
	return qpush.Err();
}

func (queue Queue) PullString() (string) {
	qpull := queue.redisClient.RPopLPush(queue.Name, "t")
	return qpull.Val()
}


func (queue *Queue) Length() int64 {
	return queue.redisClient.LLen(queue.Name).Val()
}

func (queue Queue) Pull(v interface{}) (error) {
	qpull := queue.redisClient.RPopLPush(queue.Name, "t")
	return jsonDecode(qpull.Val(), v)
}
