package data

import "fyne.io/fyne/v2/data/binding"

var (
	redisConnKeys []string
	conns         binding.StringList
)

func GetRedisConnKeys() binding.StringList {
	conns = binding.BindStringList(&redisConnKeys)
	return conns
}

func appendRedisConnkeys(name string) {
	_ = conns.Append(name)
}
