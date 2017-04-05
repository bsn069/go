package third_test

import (
	"testing"

	"github.com/ssdb/gossdb/ssdb"
)

func Test_ssdb(t *testing.T) {
	ip := "127.0.0.1"
	port := 8888
	db, err := ssdb.Connect(ip, port)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	ssdb_get_set_del(t, db)
	ssdb_zget_zset_zdel(t, db)
}

func ssdb_zget_zset_zdel(t *testing.T, db *ssdb.Client) {
	strName := "ssdb_zset_name1"
	strKey := "ssdb_zset_key1"

	resp, err := db.Do("zdel", strName, strKey)
	if err != nil || len(resp) != 2 || resp[0] != "ok" || (resp[1] != "0" && resp[1] != "1") {
		t.Error(resp, err)
	}

	resp, err = db.Do("zdel", strName, strKey)
	if err != nil || len(resp) != 2 || resp[0] != "ok" || resp[1] != "0" {
		t.Error(resp, err)
	}

	resp, err = db.Do("zget", strName, strKey)
	if err != nil || len(resp) != 1 || resp[0] != "not_found" {
		t.Error(resp, err)
	}

	resp, err = db.Do("zset", strName, strKey, 111)
	if err != nil || len(resp) != 2 || resp[0] != "ok" || resp[1] != "1" {
		t.Error(resp, err)
	}

	resp, err = db.Do("zget", strName, strKey)
	if err != nil || len(resp) != 2 || resp[0] != "ok" || resp[1] != "111" {
		t.Error(resp, err)
	}

	resp, err = db.Do("zset", strName, strKey, 222)
	if err != nil || len(resp) != 2 || resp[0] != "ok" || resp[1] != "0" {
		t.Error(resp, err)
	}

	resp, err = db.Do("zget", strName, strKey)
	if err != nil || len(resp) != 2 || resp[0] != "ok" || resp[1] != "222" {
		t.Error(resp, err)
	}

	resp, err = db.Do("zdel", strName, strKey+"1")
	if err != nil || len(resp) != 2 || resp[0] != "ok" || (resp[1] != "0" && resp[1] != "1") {
		t.Error(resp, err)
	}

	resp, err = db.Do("zset", strName, strKey+"1", 333)
	if err != nil || len(resp) != 2 || resp[0] != "ok" || resp[1] != "1" {
		t.Error(resp, err)
	}

	resp, err = db.Do("zset", strName, strKey+"1", 333)
	if err != nil || len(resp) != 2 || resp[0] != "ok" || resp[1] != "0" {
		t.Error(resp, err)
	}
}

func ssdb_get_set_del(t *testing.T, db *ssdb.Client) {
	val, err := db.Get("no this key")
	if err != nil {
		t.Fatal(val, err)
	}
	if val != nil {
		t.Fatal(val, err)
	}

	strKey := "ssdb_set_key"
	strValue := "ssdb_set_value"

	val, err = db.Del(strKey)
	if err != nil {
		t.Fatal(val, err)
	}
	if val.(bool) != true {
		t.Fatal(val, err)
	}

	val, err = db.Set(strKey, strValue)
	if err != nil {
		t.Fatal(val, err)
	}
	if val.(bool) != true {
		t.Fatal(val, err)
	}

	val, err = db.Set(strKey, strValue)
	if err != nil {
		t.Fatal(val, err)
	}
	if val.(bool) != true {
		t.Fatal(val, err)
	}

	val, err = db.Get(strKey)
	if err != nil {
		t.Fatal(val, err)
	}
	if val != strValue {
		t.Fatal(val, err)
	}

	val, err = db.Del(strKey)
	if err != nil {
		t.Fatal(val, err)
	}
	if val.(bool) != true {
		t.Fatal(val, err)
	}

	val, err = db.Del(strKey)
	if err != nil {
		t.Fatal(val, err)
	}
	if val.(bool) != true {
		t.Fatal(val, err)
	}

	val, err = db.Get(strKey)
	if err != nil {
		t.Fatal(val, err)
	}
	if val != nil {
		t.Fatal(val, err)
	}
}
