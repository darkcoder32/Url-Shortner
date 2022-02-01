package myconfig

import (
	"github.com/go-redis/redis"
)

var MyMap map[string]string
var RedisClient *redis.Client

func Init() {
	MyMap = make(map[string]string)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
