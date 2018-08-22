package conf

import (
	"time"
)

const TEST_KEY = "rgoq-test-key"
const TEST_KEY2 = "dd-key-test"

type Message struct {
	Msg string
	CreatedAt time.Time
}
