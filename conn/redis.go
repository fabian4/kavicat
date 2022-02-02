package conn

import (
	"context"
	"github.com/fabian4/kavicat/client"
	"github.com/fabian4/kavicat/data"
	"github.com/go-redis/redis/v8"
)

var (
	rdc *redis.Client
	ctx = context.Background()
)

func NewRedisConn(redisConn *data.RedisConn) {
	rdc = redis.NewClient(&redis.Options{
		Addr:     redisConn.Host + ":" + redisConn.Port,
		Password: redisConn.Auth,
		DB:       0,
	})

	err := rdc.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		client.Fail("connect fail")
	} else {
		client.Success("connect success")
	}

	if redisConn.Name == "" {
		redisConn.Name = redisConn.Host + ":" + redisConn.Name
	}

	data.AddRedisConn(*redisConn)
}
