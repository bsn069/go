package third_test

import (
	"testing"

	"fmt"

	"github.com/bsn069/go/path"
	"github.com/bsn069/go/utils"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func Test_leveldb(t *testing.T) {
	strGoPath := utils.Env("BSN_GITHUB_PATH")
	strDBPath := path.Join(strGoPath, "nogit/go_test/leveldb")
	db, err := leveldb.OpenFile(strDBPath, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	k1 := []byte("k11")
	v1 := []byte("v11")
	err = db.Put(k1, v1, nil)
	if err != nil {
		t.Fatal(err)
	}

	data, err := db.Get(k1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != string(v1) {
		t.Fatal()
	}

	err = db.Delete(k1, nil)
	if err != nil {
		t.Fatal(err)
	}

	batch := new(leveldb.Batch)
	batch.Put([]byte("foo"), []byte("value"))
	batch.Put([]byte("a1"), []byte("b1"))
	batch.Put([]byte("a2"), []byte("b2"))
	batch.Put([]byte("a3"), []byte("b3"))
	batch.Put([]byte("a4"), []byte("b4"))
	batch.Put([]byte("bar"), []byte("another value"))
	batch.Delete([]byte("baz"))
	err = db.Write(batch, nil)
	if err != nil {
		t.Fatal(err)
	}

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Println(string(key), string(value))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(`util.BytesPrefix([]byte("a"))`)
	iter = db.NewIterator(util.BytesPrefix([]byte("a")), nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Println(string(key), string(value))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(`&util.Range{[]byte("a2"), []byte("a4")}`)
	iter = db.NewIterator(&util.Range{[]byte("a2"), []byte("a4")}, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Println(string(key), string(value))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		t.Fatal(err)
	}
}
