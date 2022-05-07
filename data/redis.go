package data

import (
	"context"
	"fyne.io/fyne/v2/data/binding"
	"github.com/fabian4/kavicat/event"
	"github.com/go-redis/redis/v8"
	"log"
	"regexp"
	"strconv"
)

var (
	RedisConnAuth string
	RedisConnName string
	rdc           *redis.Client
	ctx           = context.Background()
	RedisKey      = binding.NewString()
	RedisValue    = binding.NewString()
	RedisCount    = binding.NewString()
	RedisClient   = binding.NewString()
	RedisMemory   = binding.NewString()
	RedisKeys     = binding.NewStringList()
)

func NewRedisConn(host string, port string, auth string) {
	RedisConnAuth = auth
	RedisConnName = host + ":" + port

	rdc = redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: auth,
		DB:       0,
	})

	err := rdc.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		event.Emit("connection_fail", "Connection Fail", err.Error(), "redis")
		return
	}

	event.Emit("connection_success", "Connected", host+" : "+port)
	rdc.Del(ctx, "key")

	event.Emit("switchUI", "Redis")

}

func SwitchDB(index int) {
	rdc = redis.NewClient(&redis.Options{
		Addr:     RedisConnName,
		Password: RedisConnAuth,
		DB:       index,
	})
	RefreshRedisKeyLists()
}

func SetRedisValuesByKeyId(id int) {
	key, _ := RedisKeys.GetValue(id)
	value := rdc.Get(ctx, key).Val()
	log.Println("get " + key + ": " + value)
	_ = RedisKey.Set(key)
	_ = RedisValue.Set(value)

	RefreshRedisKeyLists()
}

func SetRedisValuesByKey(key string) {
	value := rdc.Get(ctx, key).Val()
	log.Println("get " + key + ": " + value)
	_ = RedisValue.Set(value)

	RefreshRedisKeyLists()
}

func DeleteRedisValuesByKey(key string) {
	_ = RedisKey.Set(" ")
	_ = RedisValue.Set(" ")
	del(key)

	RefreshRedisKeyLists()
}

func SaveRedisValuesByKeyAndValue(key string, value string) {
	_ = RedisKey.Set(key)
	_ = RedisValue.Set(value)
	save(key, value)

	RefreshRedisKeyLists()
}

func RefreshRedisKeyLists() {
	redisKeys := getRedisKeys()
	_ = RedisKeys.Set(redisKeys)
	_ = RedisKey.Set("")
	_ = RedisValue.Set("")
	_ = RedisCount.Set("keys: " + strconv.Itoa(len(redisKeys)))
	getInfo()
}

func getRedisKeys() []string {
	keys := rdc.Keys(ctx, "*").Val()
	return keys
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
	event.Emit("switchUI", "Redis")
}

func getInfo() {
	clientsInfo := rdc.Info(ctx, "Clients").Val()
	re1 := regexp.MustCompile("connected_clients:[0-9]*")
	clients := re1.FindAllString(clientsInfo, -1)[0]
	_ = RedisClient.Set("clients: " + clients[18:])

	memoryInfo := rdc.Info(ctx, "Memory").Val()
	re2 := regexp.MustCompile("used_memory_human:[0-9]*\\.?[0-9]*K")
	memory := re2.FindAllString(memoryInfo, -1)[0]
	_ = RedisMemory.Set("memory: " + memory[18:])
}
