package data

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
)

func NewLevelDBConn(uri string) {
	db, err := leveldb.OpenFile(uri, nil)
	if err != nil {
		fmt.Println(err)
	}
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Println(string(key) + ": " + string(value))
	}
	iter.Release()
}
