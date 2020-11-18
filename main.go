package main

import (
	"flag"
	"net/http"
	"os"
	"syscall"

	"github.com/orzzzli/orz_cms/src/core"

	"github.com/orzzzli/orz_cms/src/controller"

	"github.com/orzzzli/orz_cms/src/source"

	"github.com/orzzzli/orzconfiger"

	"github.com/orzzzli/orz_cms/src/logger"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "path", "./config.ini", "config file full path")
	flag.Parse()

	logger.Info("config path is " + configPath)

	if configPath == "" {
		logger.Fatal("config path is empty")
	}
	orzconfiger.InitConfiger(configPath)

	initBaseDB()
	//initLogger()

	http.HandleFunc("/", controller.IndexHandler)
	http.HandleFunc("/install", controller.InstallHandler)

	initServer()
}

//初始化base mysql
func initBaseDB() {
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

	core.GlobalDB = source.NewMysql(title, timeGap)
	err := core.GlobalDB.Connect(username + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?charset=" + charset)
	if err != nil {
		logger.Fatal("base db init error " + err.Error())
	}
}

//初始化logger
func initLogger() {
	logPath, ok := orzconfiger.GetString("log", "path")
	if !ok {
		logger.Fatal("cant found path in log.path")
	}
	logger.Info("log path is:" + logPath)
	info, err := os.Stat(logPath)
	if err == nil {
		//非文件夹
		if !info.IsDir() {
			logger.Fatal("log path is not a folder")
		}
	}
	//是否存在
	if os.IsNotExist(err) {
		logger.Fatal("log path is not exist")
	}
	//是否有读写权限
	err = syscall.Access(logPath, syscall.O_RDWR)
	if err != nil {
		logger.Fatal("log path cant read and write.")
	}
}

//初始化httpServer
func initServer() {
	listenPort, ok := orzconfiger.GetString("server", "port")
	if !ok {
		logger.Fatal("cant found port in server.port")
	}
	logger.Info("server listen:" + listenPort)

	err := http.ListenAndServe("0.0.0.0:"+listenPort, nil)
	if err != nil {
		logger.Fatal("server start error:" + err.Error())
	}
}
