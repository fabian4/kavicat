package data

import "github.com/go-redis/redis/v8"

var redisConns = make(map[string]*RedisConn)

type RedisConn struct {
	Host   string
	Port   string
	Auth   string
	Name   string
	Client *redis.Client
}

func AddRedisConn(redisConn *RedisConn) {
	redisConns[redisConn.Name] = redisConn
	appendRedisConnkeys(redisConn.Name)
}

func GetRedisConn(name string) *RedisConn {
	return redisConns[name]
}

func HasRedisConn(name string) bool {
	_, found := redisConns[name]
	return found
}
