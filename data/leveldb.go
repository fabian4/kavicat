package data

import (
	"fyne.io/fyne/v2/data/binding"
	"github.com/fabian4/kavicat/event"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"strings"
)

var (
	db              *leveldb.DB
	LevelDBConnName string
	LevelDBKey      = binding.NewString()
	LevelDBValue    = binding.NewString()
	LevelDBKeys     = binding.NewStringList()
)

func NewLevelDBConn(uri string) {
	db, _ = leveldb.OpenFile(uri, nil)
	LevelDBConnName = strings.ReplaceAll(uri, "\\\\", "\\")

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		_ = LevelDBKeys.Append(string(key))
	}
	iter.Release()

	event.Emit("switchUI", "LevelDB")
}

func SetLevelDBValuesByKeyId(id int) {
	key, _ := LevelDBKeys.GetValue(id)
	value := rdc.Get(ctx, key).Val()
	log.Println("get " + key + ": " + value)
	_ = RedisKey.Set(key)
	_ = RedisValue.Set(value)

	RefreshRedisKeyLists()
}
