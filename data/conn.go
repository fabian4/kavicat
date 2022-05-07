package data

import (
	"fyne.io/fyne/v2/data/binding"
	"strconv"
)

var (
	redisConnKeys []string
	conn          *RedisConn
	Keys          = binding.NewStringList()
	Conns         = binding.NewStringList()
	Key           = binding.NewString()
	Value         = binding.NewString()
	Count         = binding.NewString()
	Client        = binding.NewString()
	Memory        = binding.NewString()
)

func SetConnInfoById(id int) {
	connection := GetRedisConn(redisConnKeys[id])
	conn = connection
	if connection.Client == nil {
		ReconnectRedis(connection, 0)
	}
	refreshKeyLists()
}

func SwitchDB(index int) {
	if conn != nil {
		ReconnectRedis(conn, index)
		refreshKeyLists()
	}
}

func SetValuesByKeyId(id int) {
	key, _ := Keys.GetValue(id)
	_ = Key.Set(key)
	_ = Value.Set(get(key))

	refreshKeyLists()
}

func SetValuesByKey(key string) {
	_ = Value.Set(get(key))

	refreshKeyLists()
}

func DeleteValuesByKey(key string) {
	_ = Key.Set(" ")
	_ = Value.Set(" ")
	del(key)

	refreshKeyLists()
}

func SaveValuesByKeyAndValue(key string, value string) {
	_ = Key.Set(key)
	_ = Value.Set(value)
	save(key, value)

	refreshKeyLists()
}

func refreshKeyLists() {
	redisKeys := getRedisKeys()
	_ = Keys.Set(redisKeys)
	_ = Count.Set("keys: " + strconv.Itoa(len(redisKeys)))
	getInfo()
}

func GetRedisConnKeys() binding.StringList {
	redisConnKeys = append(redisConnKeys, "127.0.0.1:6379")
	redisConnKeys = append(redisConnKeys, "192.168.34.67:2579")
	redisConnKeys = append(redisConnKeys, "10.103.24.3:6738")
	_ = Conns.Set(redisConnKeys)
	InitRedisConns()
	return Conns
}

func appendRedisConnkeys(name string) {
	_ = Conns.Append(name)
}
