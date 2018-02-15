package config

import "gopkg.in/redis.v5"
import "time"


// REDIS WW
var REDIS_OPTIONS  = &redis.Options{
	Addr: "127.0.0.1:6379",
	PoolSize: 10000,
	PoolTimeout: 5 * time.Second,
	Password: "",
	DB: 9,
}

