package data

import "github.com/go-redis/redis/v8"

var redisConns = make(map[string]RedisConn)

type RedisConn struct {
	Host   string
	Port   string
	Auth   string
	Name   string
	Client *redis.Client
}

func AddRedisConn(redisConn RedisConn) {
	//todo: 已存在链接
	redisConns[redisConn.Name] = redisConn
	appendRedisConnkeys(redisConn.Name)
}
