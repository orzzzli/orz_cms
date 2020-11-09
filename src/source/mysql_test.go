package source

import (
	"errors"
	"testing"
)

var testScheme = "root:root@tcp(127.0.0.1:3306)/cms?charset=utf8"
var testSelect = "select * from `data_sources`"
var testInsert = "insert into `data_sources`(`title`,`desc`,`type`,`scheme`,`struct`) value (\"abc\",\"ddd\",1,\"qqq\",\"aaa\")"

func TestMysql_Connect(t *testing.T) {
	m := NewMysql("test_cms", 30)
	err := m.Connect(testScheme)
	if err != nil {
		t.Error(err)
	}
}

func TestMysql_Ping(t *testing.T) {
	m := NewMysql("test_cms", 30)
	//no connect
	err := m.Ping()
	if err == nil {
		t.Error("no connection, still can work")
	}
	err = m.Connect(testScheme)
	if err != nil {
		t.Error(err)
	}
	err = m.Ping()
	if err != nil {
		t.Error(err)
	}
}

func TestMysql_Get(t *testing.T) {
	m := NewMysql("test_cms", 30)
	err := m.Connect(testScheme)
	if err != nil {
		t.Error(err)
	}
	res, err := m.Get(testSelect)
	if err != nil {
		t.Error(err)
	}
	_, ok := res.([]map[string]string)
	if !ok {
		t.Error(errors.New("get result cant change to map list"))
	}
}

func TestMysql_Set(t *testing.T) {
	m := NewMysql("test_cms", 30)
	err := m.Connect(testScheme)
	if err != nil {
		t.Error(err)
	}
	res, err := m.Set(testInsert)
	if err != nil {
		t.Error(err)
	}
	res1, ok := res.([]int)
	if !ok {
		t.Error(errors.New("set result cant change to list"))
	}
	if res1[1] != 1 {
		t.Error(errors.New("insert fail"))
	}
}
