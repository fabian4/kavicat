package data

import (
	"fyne.io/fyne/v2/data/binding"
	"github.com/dgraph-io/badger/v3"
	"github.com/fabian4/kavicat/event"
	"log"
	"strings"
)

var (
	badgerDB       *badger.DB
	BadgerConnName string
	BadgerKey      = binding.NewString()
	BadgerValue    = binding.NewString()
	BadgerKeys     = binding.NewStringList()
)

func NewBadgerConn(uri string) {
	badgerDB, _ = badger.Open(badger.DefaultOptions(uri))
	log.Println("open " + uri)
	BadgerConnName = strings.ReplaceAll(uri, "\\\\", "\\")
	//
	//_ = badgerDB.Update(func(txn *badger.Txn) error {
	//	_ = txn.Set([]byte("aaa"), []byte("111asdv1"))
	//	_ = txn.Set([]byte("bbb"), []byte("11asdv11"))
	//	_ = txn.Set([]byte("ccc"), []byte("111acs1"))
	//	_ = txn.Set([]byte("eee"), []byte("111avsd1"))
	//	_ = txn.Set([]byte("dsfvs"), []byte("1av1"))
	//	return nil
	//})
	_ = badgerDB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			k := it.Item().Key()
			_ = BadgerKeys.Append(string(k))
		}
		return nil
	})
	if BadgerKeys.Length() == 0 {
		event.Emit("empty", "Badger")
		return
	}
	event.Emit("switchUI", "Badger")
}

func SetBadgerValuesByKeyId(id int) {
	key, _ := BadgerKeys.GetValue(id)
	_ = BadgerKey.Set(key)
	_ = badgerDB.View(func(txn *badger.Txn) error {
		item, _ := txn.Get([]byte(key))
		_ = item.Value(func(val []byte) error {
			_ = BadgerValue.Set(string(val))
			log.Println("get " + key + ": " + string(val))
			return nil
		})
		return nil
	})

	RefreshBadgerKeyLists()
}

func SetBadgerValuesByKey(key string) {
	_ = BadgerKey.Set(key)
	_ = badgerDB.View(func(txn *badger.Txn) error {
		item, _ := txn.Get([]byte(key))
		_ = item.Value(func(val []byte) error {
			_ = BadgerValue.Set(string(val))
			log.Println("get " + key + ": " + string(val))
			return nil
		})
		return nil
	})

	RefreshBadgerKeyLists()
}

func DeleteBadgerValuesByKey(key string) {
	log.Println("del " + key)
	_ = badgerDB.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		return err
	})
	_ = BadgerKey.Set("")
	_ = BadgerValue.Set("")

	RefreshBadgerKeyLists()
}

func SaveBadgerValuesByKeyAndValue(key string, value string) {
	log.Println("save " + key + ": " + value)
	_ = BadgerKey.Set(key)
	_ = BadgerValue.Set(value)
	_ = badgerDB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})

	RefreshBadgerKeyLists()
}

func RefreshBadgerKeyLists() {
	_ = BadgerKeys.Set(nil)
	_ = badgerDB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			k := it.Item().Key()
			_ = BadgerKeys.Append(string(k))
		}
		return nil
	})
}
