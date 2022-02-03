package data

import "fyne.io/fyne/v2/data/binding"

var (
	redisConnKeys []string
	conns         binding.StringList
)

func GetRedisConnkeysById(id int) string {
	return redisConnKeys[id]
}

func GetRedisConnKeys() binding.StringList {
	redisConnKeys = append(redisConnKeys, "aaaaaaaaaaa")
	redisConnKeys = append(redisConnKeys, "bbbbbbbbbbb")
	redisConnKeys = append(redisConnKeys, "ccccccccccc")
	conns = binding.BindStringList(&redisConnKeys)
	return conns
}

func appendRedisConnkeys(name string) {
	_ = conns.Append(name)
}
