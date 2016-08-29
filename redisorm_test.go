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
		t.Error("Expected set status nil, got ", err)
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
