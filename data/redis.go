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
	connAuth string
	connName string
	rdc      *redis.Client
	Key      = binding.NewString()
	Value    = binding.NewString()
	Count    = binding.NewString()
	Client   = binding.NewString()
	Memory   = binding.NewString()
	ctx      = context.Background()
	Keys     = binding.NewStringList()
)

func NewRedisConn(host string, port string, auth string) {
	connAuth = auth
	connName = host + ":" + port

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
		Addr:     connName,
		Password: connAuth,
		DB:       index,
	})
	RefreshKeyLists()
}

func GetConnName() string {
	return connName
}

func SetValuesByKeyId(id int) {
	key, _ := Keys.GetValue(id)
	value := rdc.Get(ctx, key).Val()
	log.Println("get " + key + ": " + value)
	_ = Key.Set(key)
	_ = Value.Set(value)

	RefreshKeyLists()
}

func SetValuesByKey(key string) {
	value := rdc.Get(ctx, key).Val()
	log.Println("get " + key + ": " + value)
	_ = Value.Set(value)

	RefreshKeyLists()
}

func DeleteValuesByKey(key string) {
	_ = Key.Set(" ")
	_ = Value.Set(" ")
	del(key)

	RefreshKeyLists()
}

func SaveValuesByKeyAndValue(key string, value string) {
	_ = Key.Set(key)
	_ = Value.Set(value)
	save(key, value)

	RefreshKeyLists()
}

func RefreshKeyLists() {
	redisKeys := getRedisKeys()
	_ = Keys.Set(redisKeys)
	_ = Key.Set("")
	_ = Value.Set("")
	_ = Count.Set("keys: " + strconv.Itoa(len(redisKeys)))
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
	_ = Client.Set("clients: " + clients[18:])

	memoryInfo := rdc.Info(ctx, "Memory").Val()
	re2 := regexp.MustCompile("used_memory_human:[0-9]*\\.?[0-9]*K")
	memory := re2.FindAllString(memoryInfo, -1)[0]
	_ = Memory.Set("memory: " + memory[18:])
}
