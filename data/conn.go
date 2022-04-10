package data

import (
	"fyne.io/fyne/v2/data/binding"
	"github.com/fabian4/kavicat/conn"
)

var (
	redisConnKeys []string
	conns         binding.StringList
)

func SetDataInfoById(id int) string {
	connection := GetRedisConn(redisConnKeys[id])
	if connection.Client == nil {
		conn.ReconnectRedis(connection)
	}

	return "ok"
}

func GetRedisConnKeys() binding.StringList {
	redisConnKeys = append(redisConnKeys, "127.0.0.1:6379")
	redisConnKeys = append(redisConnKeys, "192.168.34.67:2579")
	redisConnKeys = append(redisConnKeys, "10.103.24.3:6738")
	conns = binding.BindStringList(&redisConnKeys)
	return conns
}

func appendRedisConnkeys(name string) {
	_ = conns.Append(name)
}
