package rgoq

import (
	"testing"
	"github.com/user/rgoq"
	"github.com/user/rgoq/config"
	"github.com/user/rgoq/example/conf"
	"time"
)



func TestCreateQueu(t *testing.T) {
	q := rgoq.CreateQueue(config.REDIS_OPTIONS, conf.TEST_KEY)
	if q.Name != conf.TEST_KEY {
		t.Errorf("q is not equvalent %s", conf.TEST_KEY)
	}
}


func TestPushString(t *testing.T) {
	q := rgoq.CreateQueue(config.REDIS_OPTIONS, conf.TEST_KEY)

	stest := "Test"
	msg := conf.Message{stest, time.Now()}
	err := q.Push(msg)

	if err != nil {
		t.Errorf(" expected nil, actual %s", err.Error())
	}

	var m conf.Message
	q.Pull(m)

	if m.Msg != stest {
		t.Errorf(" expected %s, actual %s", stest, m.Msg)
	}


}
