package data

import (
	"context"
	"github.com/fabian4/kavicat/event"
	"github.com/go-redis/redis/v8"
	"log"
	"regexp"
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

	err := rdc.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		event.Emit("connection_fail", "Connection Fail", err.Error(), "redis")
		return
	}

	event.Emit("connection_success", "Connected", redisConn.Host+" : "+redisConn.Port)
	rdc.Del(ctx, "key")

	redisConn.Client = rdc
	AddRedisConn(redisConn)

}

func ReconnectRedis(redisConn *RedisConn, index int) {
	rdc = redis.NewClient(&redis.Options{
		Addr:     redisConn.Host + ":" + redisConn.Port,
		Password: redisConn.Auth,
		DB:       index,
	})

	err := rdc.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		event.Emit("connection_fail", "Connection Fail", err.Error(), "redis")
		return
	}

	event.Emit("connection_success", "Connected", redisConn.Host+" : "+redisConn.Port)
	rdc.Del(ctx, "key")

	redisConn.Client = rdc
}

func getRedisKeys() []string {
	keys := rdc.Keys(ctx, "*").Val()
	return keys
}

func get(key string) string {
	value := rdc.Get(ctx, key).Val()
	log.Println("get " + key + ": " + value)
	return value
}

func del(key string) {
	log.Println("del " + key)
	err := rdc.Del(ctx, key).Err()
	if err != nil {
		event.Emit("operation_fail", "Delete Fail", err.Error())
	}
	event.Emit("operation_success", "Delete success", "del "+key)
}

func save(key string, value string) {
	log.Println("save " + key + ": " + value)
	err := rdc.Set(ctx, key, value, 0).Err()
	if err != nil {
		event.Emit("operation_fail", "Save Fail", err.Error())
	}
	event.Emit("operation_success", "Save success", "save [ key: "+key+", value: "+value+" ]")
}

func getInfo() {
	clientsInfo := rdc.Info(ctx, "Clients").Val()
	re1 := regexp.MustCompile("connected_clients:[0-9]*")
	clients := re1.FindAllString(clientsInfo, -1)[0]
	_ = Client.Set("clients: " + clients[18:])

	memoryInfo := rdc.Info(ctx, "Memory").Val()
	re2 := regexp.MustCompile("used_memory_human:[0-9]*\\.?[0-9]*K")
	memory := re2.FindAllString(memoryInfo, -1)[0]
	_ = Memory.Set("memory: " + memory[18:])
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
