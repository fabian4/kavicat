package conn

import (
	"context"
	"github.com/fabian4/kavicat/data"
	"github.com/fabian4/kavicat/event"
	"github.com/go-redis/redis/v8"
)

var (
	rdc *redis.Client
	ctx = context.Background()
)

func NewRedisConn(redisConn *data.RedisConn) {
	if redisConn.Name == "" {
		redisConn.Name = redisConn.Host + ":" + redisConn.Port
	}

	if data.HasRedisConn(redisConn.Name) {
		event.Emit("connection_exist", "The connection exists, No need to recreate")
		return
	}

	rdc = redis.NewClient(&redis.Options{
		Addr:     redisConn.Host + ":" + redisConn.Port,
		Password: redisConn.Auth,
		DB:       0,
	})

	//todo: pick a test api
	err := rdc.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		event.Emit("connection_fail", "Connection Fail", err.Error(), "redis")
		return
	}

	event.Emit("connection_success", "Connected", redisConn.Host+" : "+redisConn.Port)

	redisConn.Client = rdc
	data.AddRedisConn(*redisConn)

}
