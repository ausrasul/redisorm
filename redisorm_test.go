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
	err := Set("test", objt)
	if err != nil {
		t.Error("Expected set status nil, got ", err)
	}
	var obj string
	err = Get("test", &obj)
	if err != nil {
		t.Error("Expected get status nil, got ", err)
	}
	if obj != "sss" {
		t.Error("Expected obj = \"sss\", got ", obj)
	}

}

func TestPool(t *testing.T){
	for i:=0; i<1000; i++{
		var objt string = "sss"
		err := Set("test", objt)
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
	Set("test", objt)
	err := Get("test", &obj)
	if err != nil {
		t.Error("Expected get status nil, got ", err)
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
	err = Get("test", &obj)
	if err == nil {
		t.Error("Expected get status not nil, got ", err)
	}
}

func TestGetNotExist(t *testing.T){
	var obj string
	err := Get("test1", &obj)
	if err == nil {
		t.Error("Expected get status not nil, got ", err)
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
	err := Get("test1", &obj)
	if err == nil {
		t.Error("Expected get status not nil, got ", err)
	}
	if obj.Name != "" {
		t.Error("Expected obj.Name = \"\", got ", obj)
	}
}
