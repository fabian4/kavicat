package data

import (
	"fyne.io/fyne/v2/data/binding"
)

var (
	redisConnKeys []string
	Keys          = binding.NewStringList()
	Conns         = binding.NewStringList()
	Key           = binding.NewString()
	Value         = binding.NewString()
)

func SetConnInfoById(id int) {
	connection := GetRedisConn(redisConnKeys[id])
	if connection.Client == nil {
		ReconnectRedis(connection)
	}
	refreshKeyLists()
}

func SetValuesByKeyId(id int) {
	key, _ := Keys.GetValue(id)
	_ = Key.Set(key)
	_ = Value.Set(get(key))

	refreshKeyLists()
}

func SetValuesByKey(key string) {
	_ = Value.Set(get(key))
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
}

func refreshKeyLists() {
	redisKeys := getRedisKeys()
	_ = Keys.Set(redisKeys)
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
