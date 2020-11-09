package main

import (
	"flag"

	"github.com/orzzzli/orz_cms/src/source"

	"github.com/orzzzli/orzconfiger"

	"github.com/orzzzli/orz_cms/src/logger"
)

var GlobalDB *source.Mysql

func main() {
	var configPath string

	flag.StringVar(&configPath, "path", "./config.ini", "config file full path")
	flag.Parse()

	logger.Info("config path is " + configPath)

	if configPath == "" {
		logger.Fatal("config path is empty")
	}
	orzconfiger.InitConfiger(configPath)

	title, ok := orzconfiger.GetString("mysql", "title")
	if !ok {
		logger.Fatal("cant found host in mysql.title")
	}
	host, ok := orzconfiger.GetString("mysql", "host")
	if !ok {
		logger.Fatal("cant found host in mysql.host")
	}
	port, ok := orzconfiger.GetString("mysql", "port")
	if !ok {
		logger.Fatal("cant found host in mysql.port")
	}
	db, ok := orzconfiger.GetString("mysql", "database")
	if !ok {
		logger.Fatal("cant found host in mysql.database")
	}
	charset, ok := orzconfiger.GetString("mysql", "charset")
	if !ok {
		logger.Fatal("cant found host in mysql.charset")
	}
	timeGap, ok := orzconfiger.GetInt("mysql", "timeGap")
	if !ok {
		logger.Fatal("cant found host in mysql.timeGap")
	}
	username, ok := orzconfiger.GetString("mysql", "username")
	if !ok {
		logger.Fatal("cant found host in mysql.username")
	}
	password, ok := orzconfiger.GetString("mysql", "password")
	if !ok {
		logger.Fatal("cant found host in mysql.password")
	}

	GlobalDB = source.NewMysql(title, timeGap)
	err := GlobalDB.Connect(username + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?charset=" + charset)
	if err != nil {
		logger.Fatal("base db init error " + err.Error())
	}

	logger.Info("server started.")
}
