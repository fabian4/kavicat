package conn

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var ctx = context.Background()

func NewRedisConn(host, port, auth, name string) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: auth,
		DB:       0,
	})
	log.Println(rdb.Keys(ctx, "*"))
}
