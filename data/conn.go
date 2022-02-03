package data

import "fyne.io/fyne/v2/data/binding"

var (
	redisConnKeys []string
	conns         binding.StringList
)

func GetRedisConnKeys() binding.StringList {
	redisConnKeys = append(redisConnKeys, "localhost")
	redisConnKeys = append(redisConnKeys, "127.127.127.127:6379")
	redisConnKeys = append(redisConnKeys, "localhost")
	conns = binding.BindStringList(&redisConnKeys)
	return conns
}

func appendRedisConnkeys(name string) {
	_ = conns.Append(name)
}
