package main

import (
	"flag"
	"net/http"

	"github.com/orzzzli/orz_cms/src/controller"

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
		logger.Fatal("cant found title in mysql.title")
	}
	host, ok := orzconfiger.GetString("mysql", "host")
	if !ok {
		logger.Fatal("cant found host in mysql.host")
	}
	port, ok := orzconfiger.GetString("mysql", "port")
	if !ok {
		logger.Fatal("cant found port in mysql.port")
	}
	db, ok := orzconfiger.GetString("mysql", "database")
	if !ok {
		logger.Fatal("cant found database in mysql.database")
	}
	charset, ok := orzconfiger.GetString("mysql", "charset")
	if !ok {
		logger.Fatal("cant found charset in mysql.charset")
	}
	timeGap, ok := orzconfiger.GetInt("mysql", "timeGap")
	if !ok {
		logger.Fatal("cant found timeGap in mysql.timeGap")
	}
	username, ok := orzconfiger.GetString("mysql", "username")
	if !ok {
		logger.Fatal("cant found username in mysql.username")
	}
	password, ok := orzconfiger.GetString("mysql", "password")
	if !ok {
		logger.Fatal("cant found password in mysql.password")
	}

	GlobalDB = source.NewMysql(title, timeGap)
	err := GlobalDB.Connect(username + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?charset=" + charset)
	if err != nil {
		logger.Fatal("base db init error " + err.Error())
	}

	listenPort, ok := orzconfiger.GetString("server", "port")
	if !ok {
		logger.Fatal("cant found port in server.port")
	}

	http.HandleFunc("/", controller.IndexHandler)

	logger.Info("server started. listen:" + listenPort)

	err = http.ListenAndServe("0.0.0.0:"+listenPort, nil)
	if err != nil {
		logger.Fatal("server start error:" + err.Error())
	}
}
