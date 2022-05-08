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
	log.Println("open " + uri)
	LevelDBConnName = strings.ReplaceAll(uri, "\\\\", "\\")
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		_ = LevelDBKeys.Append(string(key))
	}
	iter.Release()
	if LevelDBKeys.Length() == 0 {
		event.Emit("empty", "LevelDB")
		return
	}
	event.Emit("switchUI", "LevelDB")
}

func SetLevelDBValuesByKeyId(id int) {
	key, _ := LevelDBKeys.GetValue(id)
	value, _ := db.Get([]byte(key), nil)
	log.Println("get " + key + ": " + string(value))
	_ = LevelDBKey.Set(key)
	_ = LevelDBValue.Set(string(value))

	RefreshLevelDBKeyLists()
}

func SetLevelDBValuesByKey(key string) {
	value, _ := db.Get([]byte(key), nil)
	log.Println("get " + key + ": " + string(value))
	_ = LevelDBValue.Set(string(value))

	RefreshLevelDBKeyLists()
}

func DeleteLevelDBValuesByKey(key string) {
	_ = db.Delete([]byte(key), nil)
	log.Println("del " + key)
	_ = LevelDBKey.Set("")
	_ = LevelDBValue.Set("")

	RefreshLevelDBKeyLists()
}

func SaveLevelDBValuesByKeyAndValue(key string, value string) {
	log.Println("save " + key + ": " + value)
	_ = db.Put([]byte(key), []byte(value), nil)
	_ = LevelDBKey.Set(key)
	_ = LevelDBValue.Set(value)

	RefreshLevelDBKeyLists()
}

func RefreshLevelDBKeyLists() {
	_ = LevelDBKeys.Set(nil)
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		_ = LevelDBKeys.Append(string(key))
	}
	iter.Release()
}
