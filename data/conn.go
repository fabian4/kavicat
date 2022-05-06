package data

import (
	"fmt"
	"fyne.io/fyne/v2/data/binding"
)

var (
	redisConnKeys []string
	Keys          = binding.NewStringList()
	Conns         = binding.NewStringList()
	Key           = binding.NewString()
	Value         = binding.NewString()
)

func SetConnInfoById(id int) string {
	connection := GetRedisConn(redisConnKeys[id])
	if connection.Client == nil {
		ReconnectRedis(connection)
	}
	redisKeys := getRedisKeys(connection.Client)
	fmt.Println(redisKeys)
	_ = Keys.Set(redisKeys)
	return "ok"
}

func SetValuesByKeyId(id int) {
	key, _ := Keys.GetValue(id)
	_ = Key.Set(key)
	_ = Value.Set(get(key))
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
