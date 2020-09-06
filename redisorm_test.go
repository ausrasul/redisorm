package redisorm

import (
	"testing"
)



func TestConfigure(t *testing.T){
	err := Configure(
		map[string]interface{}{
			"poolMaxIdle": 10,
			"poolMaxActive": 20,
			"port": "6379",
			"ip": "127.0.0.1",
		},
	)
	if err != nil{
		t.Error("Expected err nil, got err ", err)
	}
	if std.conf.maxIdle != 10 {
		t.Error("Expected maxIdle = 10, got ", std.conf.maxIdle )
	}
	if std.conf.maxActive != 20 {
		t.Error("Expected maxActive = 20, got ", std.conf.maxActive )
	}
	if std.conf.port != "6379" {
		t.Error("Expected port = \"6379\", got ", std.conf.port )
	}
	if std.conf.ip != "127.0.0.1" {
		t.Error("Expected ip = \"127.0.0.1\", got ", std.conf.ip)
	}
	if std.conf.db != 0 {
		t.Error("Expected db = 0, got ", std.conf.db)
	}
	err = Configure(
		map[string]interface{}{
			"poolMaxIdle": 10,
			"poolMaxActive": 20,
			"port": "6379",
			"ip": "127.0.0.1",
			"db": 1,
		},
	)
	if err != nil{
		t.Error("Expected err nil, got err ", err)
	}
	if std.conf.db != 1 {
		t.Error("Expected db = 1, got ", std.conf.db)
	}
}


func TestSetGet(t *testing.T){
	var objt string = "sss"
	err := Set("test1___", objt)
	if err != nil {
		t.Error("Expected set status nil, got ", err)
	}
	var obj string
	ok, err := Get("test1___", &obj)
	if err != nil {
		t.Error("Expected get status nil, got ", err)
	}
	if !ok {
		t.Error("Expected get ok true, got ", ok)
	}
	if obj != "sss" {
		t.Error("Expected obj = \"sss\", got ", obj)
	}

}

func TestPool(t *testing.T){
	for i:=0; i<1000; i++{
		var objt string = "sss"
		err := Set("test2___", objt)
		if err != nil {
			t.Error("Expected set status nil, got ", err)
		}
		c := pool.ActiveCount()
		if c != 1 {
			t.Error("Expected single active connection, got ", c)
		}
	}
}

func TestDbNumber(t *testing.T){
	db1 := 1
	db2 := 2
	objt := "sss"
	var obj string

	Configure(
		map[string]interface{}{
			"poolMaxIdle": 10,
			"poolMaxActive": 20,
			"port": "6379",
			"ip": "127.0.0.1",
			"db": db1,
		},
	)
	Set("test3___", objt)
	ok, err := Get("test3___", &obj)
	if err != nil || !ok{
		t.Error("Expected get status nil and ok true, got ", err, ok)
	}
	if obj != "sss" {
		t.Error("Expected obj = \"sss\", got ", obj)
	}

	Configure(
		map[string]interface{}{
			"poolMaxIdle": 10,
			"poolMaxActive": 20,
			"port": "6379",
			"ip": "127.0.0.1",
			"db": db2,
		},
	)
	ok, err = Get("test3___", &obj)
	if err != nil || ok {
		t.Error("Expected get status nil and ok false, got ", err, ok)
	}
	Configure(
		map[string]interface{}{
			"poolMaxIdle": 10,
			"poolMaxActive": 20,
			"port": "6379",
			"ip": "127.0.0.1",
			"db": db1,
		},
	)
}

func TestGetNotExist(t *testing.T){
	var obj string
	ok, err := Get("test4___", &obj)
	if err != nil || ok {
		t.Error("Expected get status nil, and ok false got ", err, ok)
	}
	if obj != "" {
		t.Error("Expected obj = \"\", got ", obj)
	}
}

func TestGetWrongType(t *testing.T){
	obj := struct {
		Name string
		Age int
	}{} // does not exist in db.
	Set("test5___", "sss")
	ok, err := Get("test5___", &obj)
	if err == nil || !ok {
		t.Error("Expected get status not nil and ok true, got ", err, ok)
	}
	if obj.Name != "" {
		t.Error("Expected obj.Name = \"\", got ", obj)
	}
}

func TestDel(t *testing.T){
	var objt string = "sss"
	Set("test6___", objt)
	i, err := Del("test6___")
	if err != nil {
		t.Error("Expected del status nil, got ", err)
	}
	if i != 1 {
		t.Error("Expected del to return 1, got ", i)
	}
	var obj string
	ok, err := Get("test6___", &obj)
	if err != nil || ok {
		t.Error("Expected get status nil and ok false, got ", err)
	}
	if obj != "" {
		t.Error("Expected obj = \"\", got ", obj)
	}
}

func TestDelNotExist(t *testing.T){
	i, err := Del("test7___")
	if err != nil {
		t.Error("Expected del status nil, got ", err)
	}
	if i != 0 {
		t.Error("Expected del to return nil, got ", i)
	}
}

func TestDelMulti(t *testing.T){
	_ = Set("test8___", "sss")
	_ = Set("test9___", "sss")
	i, err := Del([]string{"test8___", "test9___"})
	if err != nil {
		t.Error("Expected del status nil, got ", err)
	}
	if i != 2 {
		t.Error("Expected del to return 2, got ", i)
	}
	var obj string
	ok, _ := Get("test8___", &obj)
	if obj != "" || ok{
		t.Error("Expected obj = \"\" and ok false, got ", obj, ok)
	}
	ok, _ = Get("test9___", &obj)
	if obj != "" {
		t.Error("Expected obj = \"\" and ok false, got ", obj, ok)
	}
}
