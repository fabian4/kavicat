package data

import (
	"context"
	"fmt"
	"github.com/fabian4/kavicat/event"
	"github.com/go-redis/redis/v8"
)

var (
	rdc        *redis.Client
	ctx        = context.Background()
	redisConns = make(map[string]*RedisConn)
)

type RedisConn struct {
	Host   string
	Port   string
	Auth   string
	Name   string
	Client *redis.Client
}

func NewRedisConn(redisConn *RedisConn) {
	if redisConn.Name == "" {
		redisConn.Name = redisConn.Host + ":" + redisConn.Port
	}

	if HasRedisConn(redisConn.Name) {
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
	AddRedisConn(redisConn)

}

func ReconnectRedis(redisConn *RedisConn) {
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
}

func getRedisKeys(client *redis.Client) []string {
	keys := client.Keys(ctx, "*").Val()
	return keys
}

func get(key string) string {
	value := rdc.Get(ctx, key).Val()
	fmt.Println(value)
	return value
}

func AddRedisConn(redisConn *RedisConn) {
	redisConns[redisConn.Name] = redisConn
	appendRedisConnkeys(redisConn.Name)
}

func GetRedisConn(name string) *RedisConn {
	return redisConns[name]
}

func HasRedisConn(name string) bool {
	_, found := redisConns[name]
	return found
}

func InitRedisConns() {
	redisConns["127.0.0.1:6379"] = &RedisConn{
		Host:   "127.0.0.1",
		Port:   "6379",
		Auth:   "",
		Name:   "",
		Client: nil,
	}
	redisConns["192.168.34.67:2579"] = &RedisConn{
		Host:   "192.168.34.67",
		Port:   "2579",
		Auth:   "",
		Name:   "",
		Client: nil,
	}
	redisConns["10.103.24.3:6738"] = &RedisConn{
		Host:   "10.103.24.3",
		Port:   "6738",
		Auth:   "",
		Name:   "",
		Client: nil,
	}
}
